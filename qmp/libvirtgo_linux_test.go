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
		t.Fatalf("expected %s and got %s\n", expectedOutput, string(result))
	}
}

func TestLibvirtGoRunNoConnection(t *testing.T) {
	libvirtGoMonitor := NewLibvirtGoMonitor("testURI", "testDomain")
	_, err := libvirtGoMonitor.Run([]byte("{\"execute\" : \"query-cpus\"}"))
	if err == nil {
		t.Fatalf("no connection error expected")
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
