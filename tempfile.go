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
	"io"
	"os"
)

// newRemoveReadFileCloser opens a temporary file in the filesystem with the
// input name and returns an io.ReadCloser which removes the temporary file
// when Close is called.
func newRemoveFileReadCloser(name string) (io.ReadCloser, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	return &removeFileReadCloser{File: f}, nil
}

var _ io.ReadCloser = &removeFileReadCloser{}

// A removeFileReadCloser is a *os.File wrapper which removes the embedded file
// when its Close method is called.
type removeFileReadCloser struct {
	*os.File
}

// Close closes the file handle and removes the file from the filesystem.
func (c *removeFileReadCloser) Close() error {
	if err := c.File.Close(); err != nil {
		_ = os.Remove(c.File.Name())
		return err
	}

	return os.Remove(c.File.Name())
}
