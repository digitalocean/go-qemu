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

	"github.com/digitalocean/go-qemu/qmp"
)

var (
	uri        = flag.String("uri", "qemu:///system", `URI to connect to the libvirtd host.`)
	domainName = flag.String("domainName", "mydomain", "This is the domain to run commands against.")
)

func main() {
	flag.Parse()

	libvirtGoMonitor := qmp.NewLibvirtGoMonitor(*uri, *domainName)

	err := libvirtGoMonitor.Connect()
	if err != nil {
		log.Fatalf("Unable to connect: %v\n", err)
	}

	command := []byte("{\"execute\" : \"query-cpus\"}")
	cpus, err := libvirtGoMonitor.Run(command)
	if err != nil {
		log.Fatalf("Unable to run command: %v\n", err)
	}
	fmt.Printf("query-cpus: %s\n", string(cpus))

	if err = libvirtGoMonitor.Disconnect(); err != nil {
		log.Fatalf("Unable to disconnect: %v\n", err)
	}
}
