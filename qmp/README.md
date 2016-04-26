QMP
===

Package `qmp` enables interaction with QEMU instances via the QEMU Machine Protocol (QMP).

## Available Drivers

### Libvirt

Accessing an instance's monitor socket is available by proxying requests through the libvirt daemon.
Support for direct interaction with an instance's unix socket can be added should the need arise.

## Examples

The following examples provide a brief overview of QMP usage.

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

monitor, _ := NewLibvirtMonitor("qemu:///system", "example")
monitor.Connect()

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
monitor, _ := NewLibvirtMonitor("qemu:///system", "example")
monitor.Connect()

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
