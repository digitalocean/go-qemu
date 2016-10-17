// Copyright 2016 The go-qemu Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package qmp

import (
	"errors"
	"log"
	"os"
	"sync"
	"time"

	libvirt "github.com/rgbkrk/libvirt-go"
)

type libvirtGoMonitorLinux struct {
	LibvirtGoMonitor
	domain  string
	uri     string
	virConn *libvirt.VirConnection
	once    sync.Once
}

var eventsTable = map[int]string{
	libvirt.VIR_DOMAIN_EVENT_STARTED:   "VIR_DOMAIN_EVENT_STARTED",
	libvirt.VIR_DOMAIN_EVENT_SUSPENDED: "VIR_DOMAIN_EVENT_SUSPENDED",
	libvirt.VIR_DOMAIN_EVENT_RESUMED:   "VIR_DOMAIN_EVENT_RESUMED",
	libvirt.VIR_DOMAIN_EVENT_STOPPED:   "VIR_DOMAIN_EVENT_STOPPED",
	libvirt.VIR_DOMAIN_EVENT_SHUTDOWN:  "VIR_DOMAIN_EVENT_SHUTDOWN",
}

// NewLibvirtGoMonitor configures a connection to the provided hypervisor
// and domain.
// An error is returned if the provided libvirt connection URI is invalid.
//
// Hypervisor URIs may be local or remote, e.g.,
//	qemu:///system
//	qemu+ssh://libvirt@example.com/system
func NewLibvirtGoMonitor(uri, domain string) Monitor {
	return &libvirtGoMonitorLinux{
		uri:    uri,
		domain: domain,
	}
}

// Connect  sets up QEMU QMP connection via libvirt using
// the libvirt-go package.
// An error is returned if the libvirtd daemon is unreachable.
func (mon *libvirtGoMonitorLinux) Connect() error {

	// As per the libvirt core library:
	// For proper event handling, it is important
	// that the event implementation is registered
	// before a connection to the Hypervisor is opened.
	// We only do this once the first a Connect is called.
	mon.once.Do(eventRegisterDefaultImpl)

	virConn, err := libvirt.NewVirConnection(mon.uri)
	if err == nil {
		mon.virConn = &virConn
	}
	return err
}

// Disconnect tears down open QMP socket connections.
func (mon *libvirtGoMonitorLinux) Disconnect() error {
	var err error
	if mon.virConn != nil {
		_, err = mon.virConn.CloseConnection()
		mon.virConn = nil
	}
	return err
}

// Run executes the given QAPI command against a domain's QEMU instance.
// For a list of available QAPI commands, see:
//	http://git.qemu.org/?p=qemu.git;a=blob;f=qapi-schema.json;hb=HEAD
func (mon *libvirtGoMonitorLinux) Run(cmd []byte) ([]byte, error) {
	domain, err := mon.virConn.LookupDomainByName(mon.domain)
	if err != nil {
		return nil, err
	}

	result, err := domain.QemuMonitorCommand(
		libvirt.VIR_DOMAIN_QEMU_MONITOR_COMMAND_DEFAULT,
		string(cmd))
	if err != nil {
		return nil, err
	}

	return []byte(result), nil
}

// Events streams QEMU QMP Events.
// If a problem is encountered setting up the event monitor connection
// an error will be returned. Errors encountered during streaming will
// cause the returned event channel to be closed.
func (mon *libvirtGoMonitorLinux) Events() (<-chan Event, error) {

	if mon.virConn == nil {
		return nil, errors.New("Events() need a established connection to proceed.")
	}

	eventsChan := make(chan Event, 1)

	domain, err := mon.virConn.LookupDomainByName(mon.domain)
	if err != nil {
		return nil, err
	}

	callbackID, err := domainEventRegister(mon.virConn, &domain, eventsChan)
	if err != nil {
		return nil, err
	}

	//TODO: need a way to trigger doneChan.
	//      Maybe we need to return another type of Object
	//      instead of the Event channel so the caller has the
	//      ability to stop processing events.
	doneChan := make(chan bool)

	go pollEventsLoop(doneChan)

	go domainEventDeregister(mon.virConn, callbackID, doneChan)

	return eventsChan, nil
}

// eventRegisterDefaultImpl registers a default event implementation.
func eventRegisterDefaultImpl() {
	eventRegisterID := libvirt.EventRegisterDefaultImpl()
	if eventRegisterID == -1 {
		log.Printf(
			"EventRegisterDefaultImpl returned an unexpected value %d\n",
			eventRegisterID)
		//TODO: panic or what???
	}
}

// domainEventRegister register with libvirt to receive events for the
// specified domain.
func domainEventRegister(virConn *libvirt.VirConnection, domain *libvirt.VirDomain,
	eventsChan chan Event) (int, error) {
	callback :=
		libvirt.DomainEventCallback(newEventCallback(eventsChan))

	//TODO: add ability to register for more event types
	callbackID := virConn.DomainEventRegister(
		*domain,
		libvirt.VIR_DOMAIN_EVENT_ID_LIFECYCLE,
		&callback,
		func() {
			// empty on purpose until it's decided
			// what to do here.
		},
	)
	if callbackID == -1 {
		return -1, errors.New("Domain event registration failed!")
	}

	return callbackID, nil
}

// domainEventDeregister stops the registration with libvirt from receiving
// events for the specified domain.
func domainEventDeregister(virConn *libvirt.VirConnection, callbackID int, doneChan chan bool) {
	<-doneChan
	ret := virConn.DomainEventDeregister(callbackID)
	if ret != 0 {
		log.Printf("DomainEventDeregister returned an unexpected value: %d\n", ret)
	}
}

// newEventCallback convenient function to provide access
// to the eventChan to the returned closure/callback.
func newEventCallback(eventChan chan Event) libvirt.DomainEventCallback {
	return func(c *libvirt.VirConnection,
		d *libvirt.VirDomain,
		eventDetails interface{}, f func()) int {

		// We only support lifecycleEvents for now
		if lifecycleEvent, ok :=
			eventDetails.(libvirt.DomainLifecycleEvent); ok {
			eventChan <- constructEvent(lifecycleEvent)
			f()
		}

		// according to the libvirt-do
		// the return code is ignored
		return 0
	}
}

// pollEventsLoop keeps polling libvirt for new domain events
func pollEventsLoop(doneChan chan bool) {
	//TODO: Maybe interval should come
	//      from an ENV variable
	checkInterval := time.Tick(getPollInterval())
	for {
		select {
		case <-doneChan: // stop polling for events
			return
		case <-checkInterval:
			go func() {
				ret := libvirt.EventRunDefaultImpl()
				if ret == -1 {
					doneChan <- true
					return
				}
			}()
		}
	}
}

// getPollInterval retreives the events poll interval from the
// LIBVIRTGO_EVENTS_INTERVAL environment variable. Defaults to 1s.
func getPollInterval() time.Duration {
	intervalStr := os.Getenv("LIBVIRTGO_EVENTS_INTERVAL")
	interval, _ := time.ParseDuration("1s")
	if intervalStr != "" {
		desiredInterval, err := time.ParseDuration(intervalStr)
		if err != nil {
			log.Printf(
				"Unable value specified in 'LIBVIRTGO_EVENTS_INTERVAL': %v\n",
				err)
			return interval
		}
		interval = desiredInterval
	}
	return interval
}

// constructEvent helper function to map DomainLifecycleEvent
// into Event.
func constructEvent(eventDetails libvirt.DomainLifecycleEvent) Event {
	// Technically, the timestamp is not accurate
	// as events might occur before
	// the next libvirt.EventRunDefaultImpl() execution
	// at which point the notification is received.
	now := time.Now()
	type TimeStamp struct {
		Seconds      int64 `json:"seconds"`
		Microseconds int64 `json:"microseconds"`
	}
	timestamp := TimeStamp{
		int64(now.Second()),
		int64(now.Nanosecond()),
	}
	eventDescription := eventsTable[eventDetails.Event]
	data := make(map[string]interface{})
	data[eventDescription] = eventDetails
	return Event{
		Event:     eventDescription,
		Data:      data,
		Timestamp: timestamp,
	}
}
