package hypervisor

import (
	"net/url"
	"strings"

	"github.com/digitalocean/go-qemu/qmp"
)

var _ Driver = &LibvirtDriver{}

// A LibvirtDriver is a QEMU QMP monitor driver which utilizes libvirt by
// shelling out to 'virsh'.
type LibvirtDriver struct {
	uri *url.URL
}

// NewMonitor creates a new qmp.Monitor using libvirt with 'virsh'.
func (d *LibvirtDriver) NewMonitor(domain string) (qmp.Monitor, error) {
	return qmp.NewLibvirtMonitor(d.uri.String(), domain)
}

// DomainNames retrieves all hypervisor domain names using libvirt with
// 'virsh'.
func (d *LibvirtDriver) DomainNames() ([]string, error) {
	return virshList(d.uri.String())
}

// NewLibvirtDriver configures a LibvirtDriver using the provided hypervisor URI.
//
// Hypervisor URIs may be local or remote, e.g.,
//	qemu:///system
//	qemu+ssh://libvirt@example.com/system
func NewLibvirtDriver(uri string) (*LibvirtDriver, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	return &LibvirtDriver{
		uri: u,
	}, nil
}

// virshList shells out to 'virsh list --all --name' to produce a list of domain names.
func virshList(uri string) ([]string, error) {
	out, err := qmp.Virsh(uri, "list", "--all", "--name")
	if err != nil {
		return nil, err
	}

	// Remove blank newline at end of domains list
	domains := strings.Split(string(out), "\n")
	domains = domains[:len(domains)-2]
	return domains, nil
}
