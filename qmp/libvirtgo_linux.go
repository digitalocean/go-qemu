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
	"fmt"
	"sync"
	"time"

	libvirt "github.com/rgbkrk/libvirt-go"
)

type libvirtGoMonitorLinux struct {
	LibvirtGoMonitor
	domain  string
	uri     string
	virConn *libvirt.VirConnection
}

var once sync.Once

var eventsTable map[int]string

// Connect  sets up QEMU QMP connection via libvirt using
// the libvirt-go package.
// An error is returned if the libvirtd daemon is unreachable.
func (mon *libvirtGoMonitorLinux) Connect() error {
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
		return nil, fmt.Errorf("Events() need a established connection to proceed.")
	}

	eventsChan := make(chan Event)

	initialEventSetup()

	domain, err := mon.virConn.LookupDomainByName(mon.domain)
	if err != nil {
		return nil, err
	}

	// doneChan will allow us to stop receiving events from libvirt
	doneChan, err := domainEventRegistration(mon.virConn, &domain, eventsChan)
	if err != nil {
		return nil, err
	}

	go eventRunDefaultImplLoop(doneChan)

	go domainEventDeregister(mon.virConn, doneChan)

	return eventsChan, nil
}

// Initial setup to receive events.
// It's only done once.
func initialEventSetup() {
	once.Do(onceEventRegisterDefaultImpl)
	once.Do(populateEventTable)
}

func onceEventRegisterDefaultImpl() {
	eventRegisterId := libvirt.EventRegisterDefaultImpl()
	if eventRegisterId == -1 {
		fmt.Printf(
			"got %d on libvirt.EventRegisterDefaultImpl instead of 0\n",
			ret)
		//TODO: panic or what???
	}
}

func populateEventTable() {
	eventsTable = map[int]string{
		libvirt.VIR_DOMAIN_EVENT_STARTED:   "VIR_DOMAIN_EVENT_STARTED",
		libvirt.VIR_DOMAIN_EVENT_SUSPENDED: "VIR_DOMAIN_EVENT_SUSPENDED",
		libvirt.VIR_DOMAIN_EVENT_RESUMED:   "VIR_DOMAIN_EVENT_RESUMED",
		libvirt.VIR_DOMAIN_EVENT_STOPPED:   "VIR_DOMAIN_EVENT_STOPPED",
		libvirt.VIR_DOMAIN_EVENT_SHUTDOWN:  "VIR_DOMAIN_EVENT_SHUTDOWN",
	}
}

func domainEventRegistration(virConn *libvirt.VirConnection, domain *libvirt.VirDomain,
	eventsChan chan Event) (chan bool, error) {
	// doneChan will allow us to stop receiving events from libvirt
	doneChan := make(chan bool)
	callback :=
		libvirt.DomainEventCallback(newEventCallback(eventsChan, doneChan))

	//TODO: add ability to register for more event types
	callbackId := virConn.DomainEventRegister(
		*domain,
		libvirt.VIR_DOMAIN_EVENT_ID_LIFECYCLE,
		&callback,
		func() {
			//fmt.Printf("VIR_DOMAIN_EVENT_ID_LIFECYCLE called\n")
		},
	)
	if callback == -1 {
		close(doneChan)
		return nil, fmt.Errorf("Domain event registration failed!")
	}

	return doneChan, nil
}

func domainEventDeregister(virConn, doneChan chan bool) {
	<-doneChan
	ret := virConn.DomainEventDeregister(callbackId)
	if ret != 0 {
		fmt.Printf("got %d on DomainEventDeregister instead of 0\n", ret)
	}
}

func newEventCallback(eventChan <-chan Event,
	doneChan chan bool) libvirt.DomainEventCallback {
	return func(c *libvirt.VirConnection,
		d *libvirt.VirDomain,
		eventDetails interface{}, f func()) int {

		//TODO: Probably need to check 
		//      if eventChan is closed so we stop processing

		// We only support lifecycleEvents for now
		if lifecycleEvent, ok :=
			eventDetails.(libvirt.DomainLifecycleEvent); ok {
			eventChan <- constructEvent(lifecycleEvent)
			f()
		} else {
			// Closing the channel indicates error
			close(eventChan)
			doneChan <- true
		}
	}
}

// eventRunDefaultImplLoop keeps polling libvirt for new events
func eventRunDefaultImplLoop(doneChan chan bool) {
	//TODO: Maybe interval should come
	//      from an ENV variable
	tick := time.Tick(1 * time.Second)
	for {
		select {
		case <-doneChan: // stop listening for events
			return
		case <-tick:
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

func constructEvent(eventDetails libvirt.DomainLifecycleEvent) *Event {
	// Technically, the timestamp is not accurate
	// as events might occur before 
	// the next libvirt.EventRunDefaultImpl() execution
	// at which point the notification is received.
	now := time.Now()
	var timestamp struct {
		Seconds      int64
		Microseconds int64
	} = {
		now.Second(),
		now.Nanosecond(),
	}
	return &Event{
		Event: eventsTable[eventDetails.Event],
		Data:  eventDetails,
		Timestamp: timestamp,
	}
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
