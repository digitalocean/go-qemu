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
	"reflect"
	"testing"

	"github.com/digitalocean/go-qemu/qmp"
)

func TestHypervisorDomains(t *testing.T) {
	wantDomains := []string{"ubuntu", "arch", "gentoo"}

	var seenDomains []string
	hv := testHypervisor(
		t,
		wantDomains,
		func(domain string) (qmp.Monitor, error) {
			seenDomains = append(seenDomains, domain)
			return &noopMonitor{}, nil
		},
	)

	domains, err := hv.Domains()
	if err != nil {
		t.Fatalf("failed to retrieve domains: %v", err)
	}

	gotDomains := make([]string, 0, len(domains))
	for _, d := range domains {
		gotDomains = append(gotDomains, d.Name)
	}

	if want, got := wantDomains, gotDomains; !reflect.DeepEqual(want, got) {
		t.Fatalf("unexpected domains:\n- want: %v\n-  got: %v",
			want, got)
	}

	if want, got := wantDomains, seenDomains; !reflect.DeepEqual(want, got) {
		t.Fatalf("did not open monitors for all domains:\n- want: %v\n-  got: %v",
			want, got)
	}
}

func TestHypervisorDomainOK(t *testing.T) {
	const wantDomain = "ubuntu"

	hv := testHypervisor(
		t,
		[]string{wantDomain},
		noopMonitorFunc(),
	)

	gotDomain, err := hv.Domain(wantDomain)
	if err != nil {
		t.Fatalf("failed to retrieve domain: %v", err)
	}

	if want, got := wantDomain, gotDomain.Name; want != got {
		t.Fatalf("unexpected domain:\n- want: %v\n-  got: %v",
			want, got)
	}
}

func TestHypervisorDomainNotFound(t *testing.T) {
	hv := testHypervisor(
		t,
		nil,
		noopMonitorFunc(),
	)

	_, err := hv.Domain("foo")
	if want, got := `domain "foo" not found`, err.Error(); want != got {
		t.Fatalf("unexpected error:\n- want: %v\n-  got: %v",
			want, got)
	}
}

func TestHypervisorDomainNames(t *testing.T) {
	wantDomains := []string{"ubuntu", "arch", "gentoo"}

	hv := testHypervisor(
		t,
		wantDomains,
		noopMonitorFunc(),
	)

	gotDomains, err := hv.DomainNames()
	if err != nil {
		t.Fatalf("failed to retrieve domain names: %v", err)
	}

	if want, got := wantDomains, gotDomains; !reflect.DeepEqual(want, got) {
		t.Fatalf("unexpected domain names:\n- want: %v\n-  got: %v",
			want, got)
	}
}

// A monitorFunc is a function which can create a qmp.Monitor.
type monitorFunc func(domain string) (qmp.Monitor, error)

func noopMonitorFunc() monitorFunc {
	return func(_ string) (qmp.Monitor, error) {
		return &noopMonitor{}, nil
	}
}

var _ qmp.Monitor = &noopMonitor{}

type noopMonitor struct{}

func (noopMonitor) Connect() error                    { return nil }
func (noopMonitor) Disconnect() error                 { return nil }
func (noopMonitor) Run(_ []byte) ([]byte, error)      { return nil, nil }
func (noopMonitor) Events() (<-chan qmp.Event, error) { return nil, nil }

type testConnectMonitor struct {
	connected bool
	noopMonitor
}

func (m *testConnectMonitor) Connect() error {
	m.connected = true
	return nil
}

func (m *testConnectMonitor) Disconnect() error {
	m.connected = false
	return nil
}

func testHypervisor(
	t *testing.T,
	domains []string,
	newMonitor monitorFunc,
) *Hypervisor {
	return New(&testDriver{
		newMonitor: newMonitor,
		domains:    domains,
	})
}

type testDriver struct {
	newMonitor monitorFunc
	domains    []string
}

func (d *testDriver) NewMonitor(domain string) (qmp.Monitor, error) {
	return d.newMonitor(domain)
}

func (d *testDriver) DomainNames() ([]string, error) {
	return d.domains, nil
}
