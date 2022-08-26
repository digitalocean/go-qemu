go-qemu [![GoDoc](http://godoc.org/github.com/digitalocean/go-qemu?status.svg)](http://godoc.org/github.com/digitalocean/go-qemu) ![Build Status](https://github.com/digitalocean/go-qemu/actions/workflows/goqemu.yml/badge.svg?branch=master) [![Report Card](https://goreportcard.com/badge/github.com/digitalocean/go-qemu)](https://goreportcard.com/report/github.com/digitalocean/go-qemu)
=======

`go-qemu` is a collection of Go packages for interacting with running QEMU
instances.  Apache 2.0 Licensed.

Feel free to join us in [`#go-qemu` on libera chat](https://web.libera.chat/)
if you'd like to discuss the project.

Installation
------------

Use `go get` to retrieve all of the packages in `go-qemu`:

```shell
$ go get github.com/digitalocean/go-qemu/...
```

Overview
--------

Here is a quick overview of each top-level package, and what they can be used for:

- `hypervisor`: Package hypervisor provides management facilities for one or
more QEMU virtual machines on a hypervisor.
  - Provides easier access for managing groups of VMs than the `qemu` package.
  - Provides access to individual `qemu.Domain` types.
- `qemu`: Package qemu provides an interface for interacting with running QEMU instances.
  - Typically used for managing a single VM.
  - Good for quick experiments.
- `qmp`: Package qmp enables interaction with QEMU instances via the QEMU Machine
Protocol (QMP).
  - Typically not used by consumers outside of this repository.
  - Wraps code-generated types with friendlier APIs.

Details
-------

Package `qemu` is used in production at DigitalOcean, alongside package
[`libvirt`](https://github.com/digitalocean/go-libvirt).  This being said, it is
possible that there may still be subtle bugs which could cause the packages to act
in unexpected ways.

If you encounter any problems, please [look at the open issues](https://github.com/digitalocean/go-qemu/issues)
and if your problem does not match any of the ones listed, file a new issue with as
much detail as you are willing to provide.

The API is not considered stable at this time.  We do not anticipate making major
changes to the API, but it is possible that the API may change over time, if deemed
necessary by the project maintainers.  If you would like to include package
`qemu` in a project, we highly recommend vendoring it into your project.
