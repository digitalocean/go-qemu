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

package qemu

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/digitalocean/go-qemu/qmp"
)

var mockSuccessJSON = []byte(`{"return":{},"id":"libvirt-00"}`)

type mockMonitor struct {
	activeJobs        bool
	alwaysFail        bool
	eventErrors       bool
	eventTimeout      bool
	invalidJSON       bool
	poweredOff        bool
	eventsUnsupported bool
	disconnected      bool
}

func (mm *mockMonitor) Connect() error {
	if mm.alwaysFail {
		return errors.New("fail")
	}

	return nil
}

func (mm *mockMonitor) Disconnect() error {
	if mm.alwaysFail {
		return errors.New("fail")
	}

	mm.disconnected = true
	return nil
}

func (mm *mockMonitor) Run(cmd []byte) ([]byte, error) {
	if mm.alwaysFail {
		return []byte(""), errors.New("fail")
	}

	if mm.invalidJSON {
		return []byte("this is invalid json"), nil
	}

	if strings.Contains(string(cmd), "blockdev-snapshot-sync") {
		return mockSuccessJSON, nil
	}

	if strings.Contains(string(cmd), "screendump") {
		return mockSuccessJSON, nil
	}

	if strings.Contains(string(cmd), "block-commit") {
		if mm.activeJobs {
			return []byte(`{"id":"libvirt-00","error":{"class":"GenericError","desc":""}}`), errors.New("no")
		}

		return mockSuccessJSON, nil
	}

	if strings.Contains(string(cmd), "system_powerdown") || strings.Contains(string(cmd), "system_reset") {
		return mockSuccessJSON, nil
	}

	if out, ok := mm.runBlockCommand(cmd); ok {
		return out, nil
	}

	if out, ok := mm.runQueryCommand(cmd); ok {
		return out, nil
	}

	return []byte("{}"), fmt.Errorf("invalid command %q", string(cmd))
}

func (mm *mockMonitor) runBlockCommand(cmd []byte) ([]byte, bool) {
	switch string(cmd) {
	case `{"execute":"block-job-complete","arguments":{"device":"drive-virtio-disk0"}}`:
		return mockSuccessJSON, true
	case `{"execute":"block-job-cancel","arguments":{"device":"drive-virtio-disk0"}}`:
		return mockSuccessJSON, true
	case `{"execute":"drive-mirror","arguments":{"device":"drive-virtio-disk0","format":"raw","mode":"absolute-paths","sync":"full","target":"/tmp/foo.img"}}`:
		return mockSuccessJSON, true
	case `{"execute":"query-block"}`:
		return []byte(`{"return":[{"io-status":"ok","device":"drive-virtio-disk0","locked":false,"removable":false,"inserted":{"iops_rd":0,"detect_zeroes":"off","image":{"virtual-size":10737418240,"filename":"/var/lib/libvirt/images/example-00000.qcow2","cluster-size":65536,"format":"qcow2","actual-size":2569801728,"format-specific":{"type":"qcow2","data":{"compat":"1.1","lazy-refcounts":true,"refcount-bits":16,"corrupt":false}},"dirty-flag":false},"iops_wr":0,"ro":false,"node-name":"#block114","backing_file_depth":0,"drv":"qcow2","iops":0,"bps_wr":0,"write_threshold":0,"encrypted":false,"bps":0,"bps_rd":0,"cache":{"no-flush":false,"direct":false,"writeback":true},"file":"/var/lib/libvirt/images/example-00000.qcow2","encryption_key_missing":false},"type":"unknown"},{"io-status":"ok","device":"drive-ide0-0-0","locked":false,"removable":true,"tray_open":false,"type":"unknown"}],"id":"libvirt-34"}`), true
	case `{"execute":"query-blockstats"}`:
		return []byte(`{"return":[{"stats":{"account_failed":false,"account_invalid":true,"idle_time_ns":2953431879,"wr_merged":0,"rd_merged":0,"flush_total_times_ns":49653,"wr_highest_offset":2821110784,"wr_bytes":9786368,"wr_operations":692,"rd_bytes":122739200,"rd_operations":36604,"flush_operations":51,"wr_total_times_ns":313253456,"rd_total_times_ns":3465673657},"parent":{"stats":{"account_failed":false,"account_invalid":true,"idle_time_ns":2953431879,"wr_merged":0,"rd_merged":0,"flush_operations":61,"wr_highest_offset":3686448128,"wr_bytes":9786368,"wr_operations":751,"rd_bytes":122567168,"rd_operations":36772,"wr_total_times_ns":313253456,"rd_total_times_ns":3465673657,"flush_total_times_ns":49653}},"device":"ide0-hd0"},{"stats":{"account_failed":false,"account_invalid":false,"wr_merged":0,"rd_merged":0,"flush_total_times_ns":0,"wr_highest_offset":0,"wr_bytes":0,"wr_operations":0,"rd_bytes":0,"rd_operations":0,"flush_operations":0,"wr_total_times_ns":0,"rd_total_times_ns":0},"device":"ide1-cd0"},{"stats":{"account_failed":false,"account_invalid":false,"wr_merged":0,"rd_merged":0,"flush_total_times_ns":0,"wr_highest_offset":0,"wr_bytes":0,"wr_operations":0,"rd_bytes":0,"rd_operations":0,"flush_operations":0,"wr_total_times_ns":0,"rd_total_times_ns":0},"device":"floppy0"},{"stats":{"account_failed":false,"account_invalid":false,"wr_merged":0,"rd_merged":0,"flush_total_times_ns":0,"wr_highest_offset":0,"wr_bytes":0,"wr_operations":0,"rd_bytes":0,"rd_operations":0,"flush_operations":0,"wr_total_times_ns":0,"rd_total_times_ns":0},"device":"sd0"}]}`), true
	case `{"execute":"query-block-jobs"}`:
		if mm.activeJobs {
			return []byte(`{"return":[{"io-status":"ok","device":"drive-virtio-disk0","busy":true,"len":10737418240,"offset":1845493760,"paused":false,"speed":0,"ready":false,"type":"mirror"}],"id":"libvirt-35"}`), true
		}

		return mockSuccessJSON, true
	}

	return nil, false
}

func (mm *mockMonitor) runQueryCommand(cmd []byte) ([]byte, bool) {
	switch string(cmd) {
	case `{"execute":"query-commands"}`:
		return []byte(`{"return":[{"name":"query-rocker-of-dpa-groups"},{"name":"query-rocker-of-dpa-flows"},{"name":"query-rocker-ports"},{"name":"query-rocker"},{"name":"block-set-write-threshold"},{"name":"x-input-send-event"},{"name":"trace-event-set-state"},{"name":"trace-event-get-state"},{"name":"rtc-reset-reinjection"},{"name":"query-acpi-ospm-status"},{"name":"query-memory-devices"},{"name":"query-memdev"},{"name":"blockdev-change-medium"},{"name":"query-named-block-nodes"},{"name":"x-blockdev-insert-medium"},{"name":"x-blockdev-remove-medium"},{"name":"blockdev-close-tray"},{"name":"blockdev-open-tray"},{"name":"x-blockdev-del"},{"name":"blockdev-add"},{"name":"query-rx-filter"},{"name":"chardev-remove"},{"name":"chardev-add"},{"name":"query-tpm-types"},{"name":"query-tpm-models"},{"name":"query-tpm"},{"name":"query-target"},{"name":"query-cpu-definitions"},{"name":"query-machines"},{"name":"device-list-properties"},{"name":"qom-list-types"},{"name":"change-vnc-password"},{"name":"nbd-server-stop"},{"name":"nbd-server-add"},{"name":"nbd-server-start"},{"name":"qom-get"},{"name":"qom-set"},{"name":"qom-list"},{"name":"query-block-jobs"},{"name":"query-balloon"},{"name":"query-migrate-parameters"},{"name":"migrate-set-parameters"},{"name":"query-migrate-capabilities"},{"name":"migrate-set-capabilities"},{"name":"query-migrate"},{"name":"query-command-line-options"},{"name":"query-uuid"},{"name":"query-name"},{"name":"query-spice"},{"name":"query-vnc-servers"},{"name":"query-vnc"},{"name":"query-mice"},{"name":"query-status"},{"name":"query-kvm"},{"name":"query-pci"},{"name":"query-iothreads"},{"name":"query-cpus"},{"name":"query-blockstats"},{"name":"query-block"},{"name":"query-chardev-backends"},{"name":"query-chardev"},{"name":"query-qmp-schema"},{"name":"query-events"},{"name":"query-commands"},{"name":"query-version"},{"name":"human-monitor-command"},{"name":"qmp_capabilities"},{"name":"add_client"},{"name":"expire_password"},{"name":"set_password"},{"name":"block_set_io_throttle"},{"name":"block_passwd"},{"name":"query-fdsets"},{"name":"remove-fd"},{"name":"add-fd"},{"name":"closefd"},{"name":"getfd"},{"name":"set_link"},{"name":"balloon"},{"name":"change-backing-file"},{"name":"drive-mirror"},{"name":"blockdev-snapshot-delete-internal-sync"},{"name":"blockdev-snapshot-internal-sync"},{"name":"blockdev-snapshot"},{"name":"blockdev-snapshot-sync"},{"name":"block-dirty-bitmap-clear"},{"name":"block-dirty-bitmap-remove"},{"name":"block-dirty-bitmap-add"},{"name":"transaction"},{"name":"block-job-complete"},{"name":"block-job-resume"},{"name":"block-job-pause"},{"name":"block-job-cancel"},{"name":"block-job-set-speed"},{"name":"blockdev-backup"},{"name":"drive-backup"},{"name":"block-commit"},{"name":"block-stream"},{"name":"block_resize"},{"name":"object-del"},{"name":"object-add"},{"name":"netdev_del"},{"name":"netdev_add"},{"name":"query-dump-guest-memory-capability"},{"name":"dump-guest-memory"},{"name":"client_migrate_info"},{"name":"migrate_set_downtime"},{"name":"migrate_set_speed"},{"name":"query-migrate-cache-size"},{"name":"migrate-start-postcopy"},{"name":"migrate-set-cache-size"},{"name":"migrate-incoming"},{"name":"migrate_cancel"},{"name":"migrate"},{"name":"xen-set-global-dirty-log"},{"name":"xen-save-devices-state"},{"name":"ringbuf-read"},{"name":"ringbuf-write"},{"name":"inject-nmi"},{"name":"pmemsave"},{"name":"memsave"},{"name":"cpu-add"},{"name":"cpu"},{"name":"send-key"},{"name":"device_del"},{"name":"device_add"},{"name":"system_powerdown"},{"name":"system_reset"},{"name":"system_wakeup"},{"name":"cont"},{"name":"stop"},{"name":"screendump"},{"name":"change"},{"name":"eject"},{"name":"quit"}],"id":"libvirt-36"}`), true
	case `{"execute":"query-status"}`:
		if mm.poweredOff == true {
			return []byte(`{"return":{"status":"shutdown","singlestep":false,"running":false},"id":"libvirt-00"}`), true
		}

		return []byte(`{"return":{"status":"running","singlestep":false,"running":true},"id":"libvirt-39"}`), true
	case `{"execute":"query-version"}`:
		return []byte(`{"return":{"qemu":{"micro":0,"minor":5,"major":2}}}`), true
	case `{"execute":"query-pci"}`:
		return []byte(`{"return":[{"bus":0,"devices":[{"bus":0}]},{"bus":1,"devices":[{"bus":1,"class_info":{"desc":"Intel Ethernet controller"}}]}]}`), true
	}

	return nil, false
}

func (mm *mockMonitor) Events() (<-chan qmp.Event, error) {
	if mm.alwaysFail {
		return nil, errors.New("fail")
	}

	if mm.eventsUnsupported {
		return nil, qmp.ErrEventsNotSupported
	}

	c := make(chan qmp.Event)
	go func() {
		events := []string{blockJobReady, blockJobCompleted}

		i := 0
		for {
			if mm.eventTimeout {
				<-time.After(1 * time.Second)
				continue
			}

			if mm.eventErrors {
				c <- qmp.Event{Event: blockJobError}
				<-time.After(1 * time.Second)
				continue
			}

			c <- qmp.Event{Event: events[i]}
			<-time.After(1 * time.Second)

			i = (i + 1) % len(events)
		}
	}()

	return c, nil
}
