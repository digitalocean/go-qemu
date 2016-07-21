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

package hypervisor

import (
	"fmt"
	"net"
	"testing"

	"github.com/digitalocean/go-libvirt/libvirttest"
)

var hvConn = func() (net.Conn, error) {
	m := libvirttest.New()
	return m, nil
}

func TestRPCDriverDomainNames(t *testing.T) {
	h := NewRPCDriver(hvConn)

	domains, err := h.DomainNames()
	if err != nil {
		t.Error(err)
	}

	wantLen := 2
	gotLen := len(domains)
	if gotLen != wantLen {
		t.Errorf("expected %d domains to be returned, got %d", wantLen, gotLen)
	}

	for i, name := range domains {
		wantName := fmt.Sprintf("aaaaaaa-%d", i+1)
		if name != wantName {
			t.Errorf("expected domain name %q, got %q", wantName, name)
		}
	}
}

func TestRPCDriverVersion(t *testing.T) {
	h := NewRPCDriver(hvConn)

	version, err := h.Version()
	if err != nil {
		t.Error(err)
	}

	expected := "1.3.4"
	if version != expected {
		t.Errorf("expected version %q, got %q", expected, version)
	}
}
