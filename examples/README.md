# go-qemu examples

This folder contains example applications that demonstrate the use of the go-qemu and go-qemu/hypervisor packages.

If the program is executed using the **"unix"** named network (i.e locally from the hypervisor host), the user under which the program executes on needs belong to 'libvirtd' group
so the account has access to **'/var/run/libvirt/libvirt-sock'**

In case of executing the program remotely and connecting thru **"tcp"**, you could configure libvirtd on the hypervisor host to allow tcp connections. Of course, this is not secured.

To see the list of posible arguments for each program, use **--help**

#### hypervisor_domain_list

[hypervisor_domain_list](./hypervisor_domain_list) demonstrates how to use the [hypervisor](../hypervisor) package to obtain a list of the domains from the connected hypervisor. The list of domains returned is of type: qemu/Domain.

To run:
```{r, engine='bash', count_lines}
   $ go get github.com/digitalocean/go-qemu
   $ go run examples/hypervisor_domain_list/main.go
   or
   $ go run examples/hypervisor_domain_list/main.go -network=tcp -address="hypervisorhost:16509"
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

[domain_details](./domain_details) 

#### domain_system_powerdown

[domain_system_powerdown](./domain_system_powerdown)

