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

package main

import (
	"flag"
	"fmt"
	"log"

	qemu "github.com/digitalocean/go-qemu"
	"github.com/digitalocean/go-qemu/qmp"
)

var (
	uri        = flag.String("uri", "qemu:///system", "URI to connect to the libvirtd host.")
	domainName = flag.String("domainName", "mydomain", "The domain to run commands against.")
)

func main() {
	flag.Parse()

	mon := qmp.NewLibvirtGoMonitor(*uri, *domainName)

	if err := mon.Connect(); err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	domain, err := qemu.NewDomain(mon, *domainName)
	if err != nil {
		log.Fatalf("failed to create domain object: %v", err)
	}
	defer domain.Close()

	// domain.CPUs will forward QMP commands to the monitor
	cpus, err := domain.CPUs()
	if err != nil {
		log.Fatalf("failed to get cpus: %v", err)
	}

	fmt.Printf("CPUs: %v\n", cpus)

	if err = mon.Disconnect(); err != nil {
		log.Fatalf("Unable to disconnect: %v\n", err)
	}
}
