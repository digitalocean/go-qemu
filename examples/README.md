# go-qemu examples

This folder contains example applications that demonstrate the use of the 
 go-qemu and go-qemu/hypervisor packages.

If the program is executed using the `unix` named network (i.e locally 
 from the hypervisor host), the user under which the program executes on
 needs belong to `libvirtd` group
so the account has access to `/var/run/libvirt/libvirt-sock`.

In case of executing the program remotely and connecting through `tcp`, 
 you could configure libvirtd on the hypervisor host to allow tcp connections.
Of course, this is not secure.

To see the list of posible arguments for each program, use `--help`.

#### hypervisor_domain_list

[hypervisor_domain_list](./hypervisor_domain_list) demonstrates how to use 
 the [hypervisor](https://godoc.org/github.com/digitalocean/go-qemu/hypervisor) 
 package to obtain a list of the domains from the connected hypervisor. 
 The list of domains returned is of type: 
[go-qemu/Domain](https://godoc.org/github.com/digitalocean/go-qemu#Domain).

To run:
```{r, engine='bash', count_lines}
   $ go get github.com/digitalocean/go-qemu/...
   $ go run examples/hypervisor_domain_list/main.go -network=tcp \
                                          -address="hypervisorhost:16509"
```


You should have an output similar to this:
```{r, engine='bash', count_lines}
Connecting to unix:///var/run/libvirt/libvirt-sock

**********Domains**********
centos7
ubuntu14.04
debian8

***************************
```


#### domain_details

[domain_details](./domain_details) domanstrates how to use the 
[go-qemu](https://godoc.org/github.com/digitalocean/go-qemu)
 to connect to a hypervisor host using 
 [qmp.NewLibvirtRPCMonitor](https://godoc.org/github.com/digitalocean/go-qemu/qmp#LibvirtRPCMonitor) 
 and get the details for a specified domain.

To run:
```{r, engine='bash', count_lines}
   $ go get github.com/digitalocean/go-qemu/...
   $ go run examples/domain_details/main.go
   or
   $ go run examples/domain_details/main.go -network=tcp \
              -address="hypervisorhost:16509" -domainName="ubuntu14.04"
```


You should have an output similar to this:
```{r, engine='bash', count_lines}

Connecting to Connecting to unix:///var/run/libvirt/libvirt-sock

Version: 1.5.3

Status: running

[ PCIDevices ]
======================================
      [ID]        [Description]
======================================
[          ] [         Host bridge]
[          ] [          ISA bridge]
[          ] [      IDE controller]
[          ] [              Bridge]
[          ] [      VGA controller]
[      net0] [ Ethernet controller]
[    sound0] [    Audio controller]
[virtio-serial0] [                    ]
[          ] [      USB controller]
[          ] [      USB controller]
[          ] [      USB controller]
[       usb] [      USB controller]
[virtio-disk0] [     SCSI controller]
[  balloon0] [                    ]

[ BlockDevices ]
========================================================================
              Device   Driver                           File
========================================================================
  drive-virtio-disk0    qcow2 /var/lib/libvirt/images/ubuntu14.04.qcow2
      drive-ide0-0-0                                        


```

#### domain_system_powerdown

[domain_system_powerdown](./domain_system_powerdown) demonstrates how to use 
 the [hypervisor](https://godoc.org/github.com/digitalocean/go-qemu/hypervisor) 
 package to shut off the specified domain.

To run:
```{r, engine='bash', count_lines}
   $ go get github.com/digitalocean/go-qemu/...
   $ go run examples/domain_system_powerdown/main.go -domainName="ubuntu14.04"
   or
   $ go run examples/domain_system_powerdown/main.go -network=tcp \
    -address="hypervisorhost:16509"  -domainName="ubuntu14.04"
```


You should have an output similar to this:
```{r, engine='bash', count_lines}
Connecting to unix:///var/run/libvirt/libvirt-sock
Domain should be shut off now
```
