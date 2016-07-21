qemu [![GoDoc](http://godoc.org/github.com/digitalocean/go-qemu?status.svg)](http://godoc.org/github.com/digitalocean/go-qemu) [![Build Status](https://travis-ci.org/digitalocean/go-qemu.svg?branch=master)](https://travis-ci.org/digitalocean/go-qemu) [![Report Card](https://goreportcard.com/badge/github.com/digitalocean/go-qemu)](https://goreportcard.com/report/github.com/digitalocean/go-qemu)
====

Package `qemu` provides an interface for interacting with running QEMU instances.

Warning
-------

All packages contained in this repository should be treated as **pre-production
software with an unstable API**.

While these package are reasonably well-tested and have seen some use inside of
DigitalOcean, there may be subtle bugs which could cause the packages to act
in unexpected ways.  Use at your own risk!

In addition, the API is not considered stable at this time.  If you would like
to include package `qemu` in a project, we highly recommend vendoring it into
your project.
