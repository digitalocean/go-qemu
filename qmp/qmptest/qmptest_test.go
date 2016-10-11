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

package qmptest_test

import (
	"encoding/json"
	"testing"

	"github.com/digitalocean/go-qemu/qmp"
	"github.com/digitalocean/go-qemu/qmp/qmptest"
)

func TestNewMonitor(t *testing.T) {
	tests := []struct {
		cmd qmp.Command
		fn  qmptest.RunFunc
		out string
	}{
		{
			cmd: qmp.Command{
				Execute: "query-version",
			},
			fn: func(cmd qmp.Command) (interface{}, error) {
				var r struct {
					ID     string      `json:"id"`
					Return qmp.Version `json:"return"`
				}
				r.Return.QEMU.Major = 2

				return r, nil
			},
			out: `{"id":"","return":{"package":"","qemu":{"major":2,"micro":0,"minor":0}}}`,
		},
		{
			cmd: qmp.Command{
				Execute: "query-commands",
			},
			fn: func(cmd qmp.Command) (interface{}, error) {
				type n struct {
					Name string `json:"name"`
				}

				var r struct {
					ID     string `json:"id"`
					Return []n    `json:"return"`
				}

				cmds := []string{
					"query-block",
					"query-commands",
					"query-version",
				}

				r.Return = make([]n, 0, len(cmds))
				for _, c := range cmds {
					r.Return = append(r.Return, n{
						Name: c,
					})
				}

				return r, nil
			},
			out: `{"id":"","return":[{"name":"query-block"},{"name":"query-commands"},{"name":"query-version"}]}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.cmd.Execute, func(t *testing.T) {
			m := qmptest.NewMonitor(func(cmd qmp.Command) (interface{}, error) {
				if want, got := tt.cmd.Execute, cmd.Execute; want != got {
					t.Fatalf("unexpected QMP command:\n- want: %q\n-  got: %q",
						want, got)
				}

				return tt.fn(cmd)
			})

			cmdb, err := json.Marshal(tt.cmd)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			out, err := m.Run(cmdb)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if want, got := tt.out, string(out); want != got {
				t.Fatalf("unexpected output:\n- want: %v\n-  got: %v",
					want, got)
			}
		})
	}
}
