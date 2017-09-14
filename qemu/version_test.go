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
	"testing"

	"github.com/digitalocean/go-qemu/qmp"
	"github.com/digitalocean/go-qemu/qmp/raw"
)

func TestVersion(t *testing.T) {
	result := qmp.Version{}
	result.QEMU.Major = 2
	result.QEMU.Minor = 5
	result.QEMU.Micro = 0

	d, done := testDomain(t, func(cmd qmp.Command) (interface{}, error) {
		if want, got := "query-version", cmd.Execute; want != got {
			t.Fatalf("unexpected QMP command:\n- want: %q\n-  got: %q",
				want, got)
		}

		return success{
			Return: result,
		}, nil
	})
	defer done()

	v, err := d.Version()
	if err != nil {
		t.Error(err)
	}

	expected := "2.5.0"
	if v != expected {
		t.Errorf("expected version %q, instead got %q", expected, v)
	}
}

func TestPackageVersion(t *testing.T) {
	result := raw.VersionInfo{}
	result.Qemu.Major = 2
	result.Qemu.Minor = 8
	result.Qemu.Micro = 0
	result.Package = "(Debian 1:2.8+dfsg-3ubuntu2.4)"

	d, done := testDomain(t, func(cmd qmp.Command) (interface{}, error) {
		if want, got := "query-version", cmd.Execute; want != got {
			t.Fatalf("unexpected QMP command:\n- want: %q\n-  got: %q",
				want, got)
		}

		return success{
			Return: result,
		}, nil
	})
	defer done()

	v, err := d.PackageVersion()
	if err != nil {
		t.Error(err)
	}

	expected := "(Debian 1:2.8+dfsg-3ubuntu2.4)"
	if v != expected {
		t.Errorf("expected package version %q, instead got %q", expected, v)
	}
}
