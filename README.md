qemu [![GoDoc](http://godoc.org/github.com/digitalocean/go-qemu?status.svg)](http://godoc.org/github.com/digitalocean/go-qemu) [![Build Status](https://travis-ci.org/digitalocean/go-qemu.svg?branch=master)](https://travis-ci.org/digitalocean/go-qemu) [![Report Card](https://goreportcard.com/badge/github.com/digitalocean/go-qemu)](https://goreportcard.com/report/github.com/digitalocean/go-qemu)
====

Package `qemu` provides an interface for interacting with running QEMU instances.

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
