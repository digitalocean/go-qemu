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

// LibvirtGoEventsInterval constant used for getPollInterval
// method
const LibvirtGoEventsInterval = "LIBVIRTGO_EVENTS_INTERVAL"

type libvirtGoMonitorLinux struct {
	LibvirtGoMonitor
	domainName string
	domain     *libvirt.VirDomain
	uri        string
	virConn    *libvirt.VirConnection
	once       sync.Once
	doneChan   chan bool
	eventsChan chan Event
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
		uri:        uri,
		domainName: domain,
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

	// Already connected?
	if mon.virConn != nil {
		return nil
	}

	virConn, err := newVirConnectionInternal(mon.uri)
	if err != nil {
		return err
	}

	domain, err := lookupDomainByName(&virConn, mon.domainName)
	if err != nil {
		closeConnectionInternal(&virConn)
		return err
	}

	mon.domain = &domain
	mon.virConn = &virConn
	mon.eventsChan = make(chan Event)
	mon.doneChan = make(chan bool)

	return err
}

var newVirConnectionInternal = func(
	uri string) (libvirt.VirConnection, error) {
	return libvirt.NewVirConnection(uri)
}

// Disconnect tears down open QMP socket connections.
func (mon *libvirtGoMonitorLinux) Disconnect() error {
	var err error
	if mon.domain != nil {
		mon.domain.Free()
		mon.domain = nil
	}

	if mon.virConn != nil {
		_, err = closeConnectionInternal(mon.virConn)
		mon.virConn = nil
		close(mon.doneChan)   // stop listenning to events
		close(mon.eventsChan) // stop sending events to clients
	}

	return err
}

var closeConnectionInternal = func(
	virConn *libvirt.VirConnection) (int, error) {
	return virConn.CloseConnection()
}

// Run executes the given QAPI command against a domain's QEMU instance.
// For a list of available QAPI commands, see:
//	http://git.qemu.org/?p=qemu.git;a=blob;f=qapi-schema.json;hb=HEAD
func (mon *libvirtGoMonitorLinux) Run(cmd []byte) ([]byte, error) {
	if mon.virConn == nil {
		return nil, errors.New("Run() need a established connection to proceed.")
	}

	result, err := qemuMonitorCommandInternal(mon.domain, string(cmd))
	if err != nil {
		return nil, err
	}

	return []byte(result), nil
}

var qemuMonitorCommandInternal = func(
	domain *libvirt.VirDomain, cmd string) (string, error) {
	return domain.QemuMonitorCommand(
		libvirt.VIR_DOMAIN_QEMU_MONITOR_COMMAND_DEFAULT,
		cmd)
}

// Events streams QEMU QMP Events.
// If a problem is encountered setting up the event monitor connection
// an error will be returned. Errors encountered during streaming will
// cause the returned event channel to be closed.
func (mon *libvirtGoMonitorLinux) Events() (<-chan Event, error) {

	if mon.virConn == nil {
		return nil,
			errors.New("Events() need a established connection to proceed.")
	}

	callbackID, err := domainEventRegister(mon)
	if err != nil {
		return nil, err
	}

	go pollEventsLoop(mon.doneChan)

	go domainEventDeregister(mon, callbackID)

	return mon.eventsChan, nil
}

var lookupDomainByName = func(
	virConn *libvirt.VirConnection,
	domainName string) (libvirt.VirDomain, error) {
	return virConn.LookupDomainByName(domainName)
}

// eventRegisterDefaultImpl registers a default event implementation.
func eventRegisterDefaultImpl() {
	eventRegisterID := eventRegisterDefaultImplInternal()
	if eventRegisterID == -1 {
		log.Printf(
			"EventRegisterDefaultImpl returned an unexpected value %d\n",
			eventRegisterID)
		//TODO: panic or what???
	}
}

var eventRegisterDefaultImplInternal = func() int {
	return libvirt.EventRegisterDefaultImpl()
}

// domainEventRegister register with libvirt to receive events for the
// specified domain.
func domainEventRegister(mon *libvirtGoMonitorLinux) (int, error) {
	callback :=
		libvirt.DomainEventCallback(newEventCallback(mon))

	//TODO: add ability to register for more event types
	callbackID := domainEventRegisterInternal(
		mon,
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

var domainEventRegisterInternal = func(
	mon *libvirtGoMonitorLinux,
	callback *libvirt.DomainEventCallback,
	fn func()) int {
	return mon.virConn.DomainEventRegister(
		*mon.domain,
		libvirt.VIR_DOMAIN_EVENT_ID_LIFECYCLE,
		callback,
		fn,
	)
}

// domainEventDeregister stops the registration with libvirt from receiving
// events for the specified domain.
func domainEventDeregister(
	mon *libvirtGoMonitorLinux,
	callbackID int) {
	<-mon.doneChan
	ret := domainEventDeregisterInternal(mon.virConn, callbackID)
	if ret != 0 {
		log.Printf("DomainEventDeregister returned an unexpected value: %d\n",
			ret)
	}
}

var domainEventDeregisterInternal = func(virConn *libvirt.VirConnection,
	callbackID int) int {
	return virConn.DomainEventDeregister(callbackID)
}

// newEventCallback convenient function to provide access
// to the eventChan to the returned closure/callback.
func newEventCallback(mon *libvirtGoMonitorLinux) libvirt.DomainEventCallback {
	return func(c *libvirt.VirConnection,
		d *libvirt.VirDomain,
		eventDetails interface{}, f func()) int {

		// if monitor is not connected,
		// we should not continue processing messages
		// as the mon.eventsChan will be closed.
		if mon.virConn == nil {
			return 0
		}

		// We only support lifecycleEvents for now
		if lifecycleEvent, ok :=
			eventDetails.(libvirt.DomainLifecycleEvent); ok {
			mon.eventsChan <- constructEvent(lifecycleEvent)
			f()
		}

		// according to the libvirt-do
		// the return code is ignored
		return 0
	}
}

// pollEventsLoop keeps polling libvirt for new domain events
func pollEventsLoop(doneChan chan bool) {
	checkInterval := time.Tick(getPollInterval())
	for {
		select {
		case <-doneChan: // stop polling for events
			return
		case <-checkInterval:
			ret := eventRunDefaultImplInternal()
			if ret == -1 {
				doneChan <- true
				return
			}
		}
	}
}

var eventRunDefaultImplInternal = func() int {
	return libvirt.EventRunDefaultImpl()
}

// getPollInterval retreives the events poll interval from the
// LIBVIRTGO_EVENTS_INTERVAL environment variable. Defaults to 1s.
func getPollInterval() time.Duration {
	intervalStr := os.Getenv(LibvirtGoEventsInterval)
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
	// This timestamp represents the moment in time when
	// the event was received on this end and not when it
	// actually occurred.
	now := time.Now()
	type TimeStamp struct {
		Seconds      int64 `json:"seconds"`
		Microseconds int64 `json:"microseconds"`
	}
	timestamp := TimeStamp{
		int64(now.Second()),
		int64(now.Second() / int(time.Microsecond)),
	}
	data := make(map[string]interface{})
	data["details"] = eventDetails
	return Event{
		Event:     eventDetails.String(),
		Data:      data,
		Timestamp: timestamp,
	}
}
