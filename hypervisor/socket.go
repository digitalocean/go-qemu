package hypervisor

import (
	"fmt"
	"time"

	"github.com/digitalocean/go-qemu/qmp"
)

var _ Driver = &SocketDriver{}

// A SocketDriver is a QEMU QMP monitor driver which communicates directly
// with a QEMU monitor socket.
type SocketDriver struct {
	addrs []SocketAddress
}

// A SocketAddress is a QEMU QMP monitor socket address used to configure
// a SocketDriver.
type SocketAddress struct {
	// Network should be one of "unix", "tcp", etc.
	Network string

	// Address should be a host:port address or UNIX socket filepath,
	// such "8.8.8.8:4444" or "/var/lib/qemu/example.socket".
	Address string

	// Timeout specifies how long the monitor socket will attempt to be
	// reached before timing out.
	Timeout time.Duration
}

// NewMonitor creates a new qmp.Monitor using a QEMU monitor socket at
// the specified address.
func (d *SocketDriver) NewMonitor(addr string) (qmp.Monitor, error) {
	for _, a := range d.addrs {
		if a.Address == addr {
			return qmp.NewSocketMonitor(
				a.Network,
				a.Address,
				a.Timeout,
			)
		}
	}

	return nil, fmt.Errorf("address not known to SocketDriver: %q", addr)
}

// DomainNames retrieves all hypervisor domain names known to the SocketDriver.
func (d *SocketDriver) DomainNames() ([]string, error) {
	names := make([]string, 0, len(d.addrs))
	for _, a := range d.addrs {
		names = append(names, a.Address)
	}

	return names, nil
}

// NewSocketDriver configures a SocketDriver using one or more SocketAddress
// structures for configuration.
func NewSocketDriver(addrs []SocketAddress) *SocketDriver {
	return &SocketDriver{
		addrs: addrs,
	}
}
