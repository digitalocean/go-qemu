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
	"testing"

	libvirt "github.com/rgbkrk/libvirt-go"
)

func TestLibvirtGoConnetOK(t *testing.T) {
	libvirtGoMonitor := NewLibvirtGoMonitor("testURI", "testDomain")
	libvirtGoMonitor.libvirtGoMonitorInternal = &fakeLibvirtGoMonitorInternal{}
	err := libvirtGoMonitor.Connect()
	if err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
}

func TestLibvirtGoConnectionFailed(t *testing.T) {
	libvirtGoMonitor := NewLibvirtGoMonitor("testURI", "testDomain")
	libvirtGoMonitor.libvirtGoMonitorInternal = &fakeLibvirtGoMonitorInternal{
		connErr: errors.New("[TEST] server down"),
	}
	err := libvirtGoMonitor.Connect()
	if err == nil {
		t.Errorf("connection error expected")
	}
}

func TestLibvirtGoDomainLookupFailed(t *testing.T) {
	libvirtGoMonitor := NewLibvirtGoMonitor("testURI", "testDomain")
	libvirtGoMonitor.libvirtGoMonitorInternal = &fakeLibvirtGoMonitorInternal{
		domainErr: errors.New("[TEST] domain not found"),
	}
	err := libvirtGoMonitor.Connect()
	if err == nil {
		t.Errorf("domain error expected")
	}
}

func TestLibvirtGoRunOK(t *testing.T) {
	libvirtGoMonitor := NewLibvirtGoMonitor("testURI", "testDomain")
	expectedResult := `{"current":true,"CPU":0,"qom_path":"/machine/unattached/device[0]","pc":33504,"halted":false,"thread_id":25213}],"id":"libvirt-32"}`
	libvirtGoMonitor.libvirtGoMonitorInternal = &fakeLibvirtGoMonitorInternal{
		expectedResult: expectedResult,
	}
	err := libvirtGoMonitor.Connect()
	if err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	result, err := libvirtGoMonitor.Run([]byte("{\"execute\" : \"query-cpus\"}"))
	if err != nil {
		t.Fatalf("unexpected error: %v\n", err)
	}

	if expectedResult != string(result) {
		t.Fatalf("Unexpected value. Expected %s and got %s\n", expectedResult, string(result))
	}
}

func TestLibvirtGoRunNoConnection(t *testing.T) {
	libvirtGoMonitor := NewLibvirtGoMonitor("testURI", "testDomain")
	_, err := libvirtGoMonitor.Run([]byte("{\"execute\" : \"query-cpus\"}"))
	if err == nil {
		t.Fatalf("no connection error expected")
	}
}

// func TestLibvirtGoEventsDomainEventRegisterOK(t *testing.T) {
// 	libvirtGoMonitor := NewLibvirtGoMonitor("testURI", "testDomain")
// 	libvirtGoMonitor.eventDefaultImplProvider = &fakeEventDefaultImplProviderOK{}
// 	libvirtGoMonitor.connectionProvider = &fakeConnectionProvider{fail: false}
// 	libvirtGoMonitor.domainFinder = &fakeDomainFinder{fail: false}
// 	err := libvirtGoMonitor.Connect()
// 	if err != nil {
// 		t.Errorf("unexpected error: %v\n", err)
// 	}

// 	var wg sync.WaitGroup
// 	expectedEvent := libvirt.DomainLifecycleEvent{
// 		Event:  libvirt.VIR_DOMAIN_EVENT_STARTED,
// 		Detail: 0,
// 	}

// 	libvirtGoMonitor.domainRegister = &fakeDomainRegisterOK{
// 		wg:            &wg,
// 		expectedEvent: &expectedEvent,
// 	}

// 	events, err := libvirtGoMonitor.Events()
// 	if err != nil {
// 		t.Errorf("unexpected error: %v\n", err)
// 	}
// 	var resultEvent Event
// 	go func(wg *sync.WaitGroup) {
// 		wg.Add(1)
// 		fmt.Println("Waiting for event....")
// 		for event := range events {
// 			resultEvent = event
// 			wg.Done()
// 			break
// 		}
// 		fmt.Println("done waiting....")
// 	}(&wg)

// 	wg.Wait()

// 	detailsEvent, found := resultEvent.Data["details"]
// 	if !found {
// 		t.Errorf("Expected at least one Event")
// 	}
// 	if expectedEvent != detailsEvent {
// 		t.Errorf("Unexpected event. Expected %#v and got %#v\n", expectedEvent, detailsEvent)
// 	}
// }

func TestLibvirtGoEventsNoConnection(t *testing.T) {
	libvirtGoMonitor := NewLibvirtGoMonitor("testURI", "testDomain")
	libvirtGoMonitor.libvirtGoMonitorInternal = &fakeLibvirtGoMonitorInternal{}
	_, err := libvirtGoMonitor.Events()
	if err == nil {
		t.Fatalf("no connection error expected")
	}
}

func TestLibvirtGoEventsDomainEventRegisterFailed(t *testing.T) {
	libvirtGoMonitor := NewLibvirtGoMonitor("testURI", "testDomain")
	libvirtGoMonitor.libvirtGoMonitorInternal = &fakeLibvirtGoMonitorInternal{
		domainEventRetCode: -1,
	}
	err := libvirtGoMonitor.Connect()
	if err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	_, err = libvirtGoMonitor.Events()
	if err == nil {
		t.Fatalf("domain register error expected")
	}
}

type fakeLibvirtGoMonitorInternal struct {
	virnConn           libvirt.VirConnection
	domain             libvirt.VirDomain
	connErr            error
	domainErr          error
	qemuErr            error
	closeErr           error
	retCode            int
	domainEventRetCode int
	expectedResult     string
}

func (fake *fakeLibvirtGoMonitorInternal) newVirConnectionInternal(uri string) (libvirt.VirConnection, error) {
	return fake.virnConn, fake.connErr
}

func (fake *fakeLibvirtGoMonitorInternal) lookupDomainByNameInternal(virConn *libvirt.VirConnection, domainName string) (libvirt.VirDomain, error) {
	return fake.domain, fake.domainErr
}

func (fake *fakeLibvirtGoMonitorInternal) qemuMonitorCommandInternal(domain *libvirt.VirDomain, cmd string) (string, error) {
	return fake.expectedResult, fake.qemuErr
}

func (fake *fakeLibvirtGoMonitorInternal) domainEventRegisterInternal(virConn *libvirt.VirConnection, domain *libvirt.VirDomain, callback *libvirt.DomainEventCallback, fn func()) int {
	return fake.domainEventRetCode
}

func (fake *fakeLibvirtGoMonitorInternal) eventRegisterDefaultImplInternal() int {
	return fake.retCode
}

func (fake *fakeLibvirtGoMonitorInternal) closeConnectionInternal(virConn *libvirt.VirConnection) (int, error) {
	return fake.retCode, fake.closeErr
}

func (fake *fakeLibvirtGoMonitorInternal) domainEventDeregisterInternal(virConn *libvirt.VirConnection, callbackID int) int {
	return fake.retCode
}

func (fake *fakeLibvirtGoMonitorInternal) eventRunDefaultImplInternal() int {
	return fake.retCode
}

// type fakeDomainRegisterOK struct {
// 	wg            *sync.WaitGroup
// 	expectedEvent *libvirt.DomainLifecycleEvent
// }

// func (fake *fakeDomainRegisterOK) domainEventRegisterInternal(
// 	callback *libvirt.DomainEventCallback, fn func()) int {
// 	fake.wg.Add(1)
// 	go fake.simulateSendingEvents(*callback)
// 	return 0
// }

// func (fake *fakeDomainRegisterOK) simulateSendingEvents(callback libvirt.DomainEventCallback) {
// 	fmt.Println("Sending event...")
// 	callback(nil, nil, fake.expectedEvent, func() {})
// 	fake.wg.Done()
// }
