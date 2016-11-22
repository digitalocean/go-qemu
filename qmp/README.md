QMP
===

Package `qmp` enables interaction with QEMU instances via the QEMU Machine Protocol (QMP).

## Available Drivers

### Libvirt

If your environment is managed by Libvirt, QMP interaction must be proxied through the Libvirt daemon. This can be done through three available drivers:

#### RPC

The RPC driver provides a pure Go implementation of Libvirt's RPC protocol.

```go
//conn, err := net.DialTimeout("unix", "/var/run/libvirt/libvirt-sock", 2*time.Second)
conn, err := net.DialTimeout("tcp", "192.168.1.1:16509", 2*time.Second)
monitor := libvirtrpc.New("stage-lb-1", conn)
```

### Libvirt-go

This monitor provides communication with the Libvirt daemon using the [libvirt-go](https://github.com/rgbkrk/libvirt-go) package as its internal implementation.
At the moment, libvirt-go only supports `Linux`.

```go
monitor, err := qmp.NewLibvirtGoMonitor("qemu:///system", "centos7")
```

#### virsh

A connection to the monitor socket is provided by proxing requests through the `virsh` executable.

```go
monitor, err := qmp.NewLibvirtMonitor("qemu:///system", "stage-lb-1")
```

### Socket

If your QEMU instances are not managed by libvirt, direct communication over its UNIX socket is available.

```go
monitor, err := qmp.NewSocketMonitor("unix", "/var/lib/qemu/example.monitor", 2*time.Second)
```

## Examples

Using the above to establish a new `qmp.Monitor`, the following examples provide a brief overview of QMP usage.

_error checking omitted for the sake of brevity._

### Command Execution
```go
type StatusResult struct {
	ID     string `json:"id"`
	Return struct {
		Running    bool   `json:"running"`
		Singlestep bool   `json:"singlestep"`
		Status     string `json:"status"`
	} `json:"return"`
}

monitor.Connect()
defer monitor.Disconnect()

cmd := []byte(`{ "execute": "query-status" }`)
raw, _ := monitor.Run(cmd)

var result StatusResult
json.Unmarshal(raw, &result)

fmt.Println(result.Return.Status)
```

```
running
```

### Event Monitor

```go
monitor.Connect()
defer monitor.Disconnect()

stream, _ := monitor.Events()
for e := range stream {
	log.Printf("EVENT: %s", e.Event)
}

```

```
$ virsh reboot example
Domain example is being rebooted
```

```
EVENT: POWERDOWN
EVENT: SHUTDOWN
EVENT: STOP
EVENT: RESET
EVENT: RESUME
EVENT: RESET
...
```

## More information

* [QEMU QMP Wiki](http://wiki.qemu.org/QMP)
* [QEMU QMP Intro](http://git.qemu.org/?p=qemu.git;a=blob_plain;f=docs/qmp-intro.txt;hb=HEAD)
* [QEMU QMP Events](http://git.qemu.org/?p=qemu.git;a=blob_plain;f=docs/qmp-events.txt;hb=HEAD)
* [QEMU QMP Spec](http://git.qemu.org/?p=qemu.git;a=blob_plain;f=docs/qmp-spec.txt;hb=HEAD)
