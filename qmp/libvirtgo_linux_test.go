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
	"testing"

	libvirt "github.com/rgbkrk/libvirt-go"
)

func TestConnetOK(t *testing.T) {

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
		t.Errorf("Unable to connect: %v\n", err)
	}

}
