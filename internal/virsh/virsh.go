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

// Package virsh provides a wrapper for shelling out to the libvirt virsh
// executable.
package virsh

import (
	"errors"
	"os/exec"
	"strings"

	"github.com/digitalocean/go-qemu/internal/shellexec"
)

// Virsh is a wrapper for shelling out to the 'virsh' executable.
func Virsh(prep shellexec.Preparer, url string, args ...string) ([]byte, error) {
	args = append([]string{"-c", url}, args...)
	cmd := prep.Prepare("virsh", args...)

	stdout, err := cmd.Output()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			msg := strings.TrimPrefix(string(ee.Stderr), "error: ")
			return stdout, errors.New(strings.ToLower(msg))
		}

		return stdout, err
	}

	return stdout, nil
}
