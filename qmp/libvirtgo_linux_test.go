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
	"os"
	"sync"
	"testing"
	"time"

	libvirt "github.com/rgbkrk/libvirt-go"
)

func TestLibvirtGoConnetOK(t *testing.T) {

	eventRegisterDefaultImplInternalOrg := eventRegisterDefaultImplInternal
	eventRegisterDefaultImplInternal = func() int {
		return 0
	}
	defer func() {
		eventRegisterDefaultImplInternal = eventRegisterDefaultImplInternalOrg
	}()

	newVirConnectionInternalOrg := newVirConnectionInternal
	newVirConnectionInternal = func(uri string) (libvirt.VirConnection, error) {
		return libvirt.VirConnection{}, nil
	}
	defer func() {
		newVirConnectionInternal = newVirConnectionInternalOrg
	}()

	lookupDomainByNameOrg := lookupDomainByName
	lookupDomainByName = func(virConn *libvirt.VirConnection,
		domainName string) (libvirt.VirDomain, error) {
		return libvirt.VirDomain{}, nil
	}
	defer func() {
		lookupDomainByName = lookupDomainByNameOrg
	}()

	libvirtGoMonitor := NewLibvirtGoMonitor("testURI", "testDomain")
	err := libvirtGoMonitor.Connect()
	if err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

}

func TestLibvirtGoConnectionFailed(t *testing.T) {

	eventRegisterDefaultImplInternalOrg := eventRegisterDefaultImplInternal
	eventRegisterDefaultImplInternal = func() int {
		return 0
	}
	defer func() {
		eventRegisterDefaultImplInternal = eventRegisterDefaultImplInternalOrg
	}()

	newVirConnectionInternalOrg := newVirConnectionInternal
	newVirConnectionInternal = func(uri string) (libvirt.VirConnection, error) {
		return libvirt.VirConnection{}, errors.New("server down")
	}
	defer func() {
		newVirConnectionInternal = newVirConnectionInternalOrg
	}()

	libvirtGoMonitor := NewLibvirtGoMonitor("testURI", "testDomain")
	err := libvirtGoMonitor.Connect()
	if err == nil {
		t.Fatalf("connection error expected")
	}

}

func TestLibvirtGoDomainLookupFailed(t *testing.T) {

	eventRegisterDefaultImplInternalOrg := eventRegisterDefaultImplInternal
	eventRegisterDefaultImplInternal = func() int {
		return 0
	}
	defer func() {
		eventRegisterDefaultImplInternal = eventRegisterDefaultImplInternalOrg
	}()

	newVirConnectionInternalOrg := newVirConnectionInternal
	newVirConnectionInternal = func(uri string) (libvirt.VirConnection, error) {
		return libvirt.VirConnection{}, nil
	}
	defer func() {
		newVirConnectionInternal = newVirConnectionInternalOrg
	}()

	lookupDomainByNameOrg := lookupDomainByName
	lookupDomainByName = func(virConn *libvirt.VirConnection,
		domainName string) (libvirt.VirDomain, error) {
		return libvirt.VirDomain{}, errors.New("Unable to find domain")
	}
	defer func() {
		lookupDomainByName = lookupDomainByNameOrg
	}()

	libvirtGoMonitor := NewLibvirtGoMonitor("testURI", "testDomain")
	err := libvirtGoMonitor.Connect()
	if err == nil {
		t.Errorf("unexpected error: %v\n", err)
	}

}

func TestLibvirtGoRunOK(t *testing.T) {

	eventRegisterDefaultImplInternalOrg := eventRegisterDefaultImplInternal
	eventRegisterDefaultImplInternal = func() int {
		return 0
	}
	defer func() {
		eventRegisterDefaultImplInternal = eventRegisterDefaultImplInternalOrg
	}()

	newVirConnectionInternalOrg := newVirConnectionInternal
	newVirConnectionInternal = func(uri string) (libvirt.VirConnection, error) {
		return libvirt.VirConnection{}, nil
	}
	defer func() {
		newVirConnectionInternal = newVirConnectionInternalOrg
	}()

	lookupDomainByNameOrg := lookupDomainByName
	lookupDomainByName = func(virConn *libvirt.VirConnection,
		domainName string) (libvirt.VirDomain, error) {
		return libvirt.VirDomain{}, nil
	}
	defer func() {
		lookupDomainByName = lookupDomainByNameOrg
	}()

	libvirtGoMonitor := NewLibvirtGoMonitor("testURI", "testDomain")
	err := libvirtGoMonitor.Connect()
	if err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	expectedOutput := `{"current":true,"CPU":0,"qom_path":"/machine/unattached/device[0]","pc":33504,"halted":false,"thread_id":25213}],"id":"libvirt-32"}`
	qemuMonitorCommandInternalOrg := qemuMonitorCommandInternal
	qemuMonitorCommandInternal = func(
		domain *libvirt.VirDomain, cmd string) (string, error) {
		return expectedOutput, nil
	}
	defer func() {
		qemuMonitorCommandInternal = qemuMonitorCommandInternalOrg
	}()

	result, err := libvirtGoMonitor.Run([]byte("{\"execute\" : \"query-cpus\"}"))
	if err != nil {
		t.Fatalf("unexpected error: %v\n", err)
	}

	if expectedOutput != string(result) {
		t.Fatalf("Unexpected value. Expected %s and got %s\n", expectedOutput, string(result))
	}
}

func TestLibvirtGoRunNoConnection(t *testing.T) {
	libvirtGoMonitor := NewLibvirtGoMonitor("testURI", "testDomain")
	_, err := libvirtGoMonitor.Run([]byte("{\"execute\" : \"query-cpus\"}"))
	if err == nil {
		t.Fatalf("no connection error expected")
	}
}

func TestLibvirtGoEventsDomainEventRegisterOK(t *testing.T) {
	eventRegisterDefaultImplInternalOrg := eventRegisterDefaultImplInternal
	eventRegisterDefaultImplInternal = func() int {
		return 0
	}
	defer func() {
		eventRegisterDefaultImplInternal = eventRegisterDefaultImplInternalOrg
	}()

	newVirConnectionInternalOrg := newVirConnectionInternal
	newVirConnectionInternal = func(uri string) (libvirt.VirConnection, error) {
		return libvirt.VirConnection{}, nil
	}
	defer func() {
		newVirConnectionInternal = newVirConnectionInternalOrg
	}()

	lookupDomainByNameOrg := lookupDomainByName
	lookupDomainByName = func(virConn *libvirt.VirConnection,
		domainName string) (libvirt.VirDomain, error) {
		return libvirt.VirDomain{}, nil
	}
	defer func() {
		lookupDomainByName = lookupDomainByNameOrg
	}()

	libvirtGoMonitor := NewLibvirtGoMonitor("testURI", "testDomain")
	err := libvirtGoMonitor.Connect()
	if err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	var wg sync.WaitGroup
	expectedEvent := libvirt.DomainLifecycleEvent{
		Event:  libvirt.VIR_DOMAIN_EVENT_STARTED,
		Detail: 0,
	}
	simulateSendingEvents := func(callback libvirt.DomainEventCallback) {
		callback(nil, nil, expectedEvent, func() {})
		wg.Done()
	}

	domainEventRegisterInternalOrg := domainEventRegisterInternal
	domainEventRegisterInternal = func(
		mon *libvirtGoMonitorLinux,
		callback *libvirt.DomainEventCallback,
		fn func()) int {
		wg.Add(1)
		go simulateSendingEvents(*callback)

		return 0
	}
	defer func() {
		domainEventRegisterInternal = domainEventRegisterInternalOrg
	}()

	events, err := libvirtGoMonitor.Events()
	if err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	var resultEvent Event
	go func() {
		wg.Add(1)
		for event := range events {
			resultEvent = event
			wg.Done()
			break
		}
	}()

	wg.Wait()

	detailsEvent, found := resultEvent.Data["details"]
	if !found {
		t.Errorf("Expected at least one Event")
	}
	if expectedEvent != detailsEvent {
		t.Errorf("Unexpected event. Expected %#v and got %#v\n", expectedEvent, detailsEvent)
	}
}

func TestLibvirtGoEventsNoConnection(t *testing.T) {
	libvirtGoMonitor := NewLibvirtGoMonitor("testURI", "testDomain")
	_, err := libvirtGoMonitor.Events()
	if err == nil {
		t.Fatalf("no connection error expected")
	}
}

func TestLibvirtGoEventsDomainEventRegisterFailed(t *testing.T) {
	eventRegisterDefaultImplInternalOrg := eventRegisterDefaultImplInternal
	eventRegisterDefaultImplInternal = func() int {
		return 0
	}
	defer func() {
		eventRegisterDefaultImplInternal = eventRegisterDefaultImplInternalOrg
	}()

	newVirConnectionInternalOrg := newVirConnectionInternal
	newVirConnectionInternal = func(uri string) (libvirt.VirConnection, error) {
		return libvirt.VirConnection{}, nil
	}
	defer func() {
		newVirConnectionInternal = newVirConnectionInternalOrg
	}()

	lookupDomainByNameOrg := lookupDomainByName
	lookupDomainByName = func(virConn *libvirt.VirConnection,
		domainName string) (libvirt.VirDomain, error) {
		return libvirt.VirDomain{}, nil
	}
	defer func() {
		lookupDomainByName = lookupDomainByNameOrg
	}()

	libvirtGoMonitor := NewLibvirtGoMonitor("testURI", "testDomain")
	err := libvirtGoMonitor.Connect()
	if err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	domainEventRegisterInternalOrg := domainEventRegisterInternal
	domainEventRegisterInternal = func(
		mon *libvirtGoMonitorLinux,
		callback *libvirt.DomainEventCallback,
		fn func()) int {
		return -1
	}
	defer func() {
		domainEventRegisterInternal = domainEventRegisterInternalOrg
	}()

	_, err = libvirtGoMonitor.Events()
	if err == nil {
		t.Fatalf("domain register error expected")
	}
}

func TestLibvirtGoGetPollIntervalDefault(t *testing.T) {

	interval := getPollInterval()
	expectedInterval, _ := time.ParseDuration("1s")

	if expectedInterval != interval {
		t.Errorf("Unexpected value. Expected [%s] and got [%s]\n", expectedInterval, interval)
	}
}

func TestLibvirtGoGetPollIntervalDefaultInvalid(t *testing.T) {
	os.Setenv(LibvirtGoEventsInterval, "invalid")
	defer func() {
		os.Unsetenv(LibvirtGoEventsInterval)
	}()
	interval := getPollInterval()
	expectedInterval, _ := time.ParseDuration("1s")

	if expectedInterval != interval {
		t.Errorf("Unexpected value. Expected [%s] and got [%s]\n", expectedInterval, interval)
	}

}

func TestLibvirtGoGetPollIntervalDefault5Second(t *testing.T) {
	os.Setenv(LibvirtGoEventsInterval, "5s")
	defer func() {
		os.Unsetenv(LibvirtGoEventsInterval)
	}()
	interval := getPollInterval()
	expectedInterval, _ := time.ParseDuration("5s")

	if expectedInterval != interval {
		t.Errorf("Unexpected value. Expected [%s] and got [%s]\n", expectedInterval, interval)
	}
}
