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
	"net"
	"time"

	"github.com/digitalocean/go-qemu/qemu"
	"github.com/digitalocean/go-qemu/qmp"
)

var (
	network    = flag.String("network", "unix", `Named network used to connect on. For Unix sockets -network=unix, for TCP connections: -network=tcp`)
	address    = flag.String("address", "/var/run/libvirt/libvirt-sock", `Address of the hypervisor. This could be in the form of Unix or TCP sockets. For TCP connections: -address="host:16509"`)
	timeout    = flag.Duration("timeout", 2*time.Second, "Connection timeout. Another valid value could be -timeout=500ms")
	domainName = flag.String("domainName", "mydomain", "This is the domain to get details for.")
)

func main() {
	flag.Parse()

	fmt.Printf("\nConnecting to %s://%s\n", *network, *address)

	c, err := net.DialTimeout(*network, *address, *timeout)
	if err != nil {
		log.Fatalf("failed to connect to hypervisor: %v", err)
	}

	mon := qmp.NewLibvirtRPCMonitor(*domainName, c)

	if err := mon.Connect(); err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	domain, err := qemu.NewDomain(mon, *domainName)
	if err != nil {
		log.Fatalf("failed to create domain object: %v", err)
	}
	defer domain.Close()

	version, err := domain.Version()
	if err != nil {
		log.Fatalf("Error getting Domain Version: %v\n", err)
	}
	fmt.Printf("\nVersion: %s\n", version)

	status, err := domain.Status()
	if err != nil {
		log.Fatalf("Error getting Domain Status: %v\n", err)
	}
	fmt.Printf("\nStatus: %s\n", status)

	displayPCIDevices(domain)

	displayBlockDevices(domain)
}

func displayPCIDevices(domain *qemu.Domain) {
	pciDevices, err := domain.PCIDevices()
	if err != nil {
		log.Fatalf("Error getting PCIDevices: %v\n", pciDevices)
	}
	fmt.Printf("\n[ PCIDevices ]\n")
	fmt.Printf("======================================\n")
	fmt.Printf("%10s %20s\n", "[ID]", "[Description]")
	fmt.Printf("======================================\n")
	for _, pciDevice := range pciDevices {
		fmt.Printf("[%10s] [%20s]\n", pciDevice.QdevID, pciDevice.ClassInfo.Desc)
	}
}

func displayBlockDevices(domain *qemu.Domain) {
	blockDevices, err := domain.BlockDevices()
	if err != nil {
		log.Fatalf("Error getting blockDevices: %v\n", blockDevices)
	}
	fmt.Printf("\n[ BlockDevices ]\n")
	fmt.Printf("========================================================================\n")
	fmt.Printf("%20s %8s %30s\n", "Device", "Driver", "File")
	fmt.Printf("========================================================================\n")
	for _, blockDevice := range blockDevices {
		fmt.Printf("%20s %8s %30s\n",
			blockDevice.Device, blockDevice.Inserted.Driver, blockDevice.Inserted.File)
	}
}
