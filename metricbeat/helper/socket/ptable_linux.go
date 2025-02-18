// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

//go:build linux
// +build linux

package socket

import (
	"github.com/k0ffee/beats/v7/libbeat/common"
)

var requiredCapabilities = []string{"sys_ptrace", "dac_read_search"}

// isPrivileged checks if this process has privileges to read sockets
// of all users
func isPrivileged() (bool, error) {
	capabilities, err := common.GetCapabilities()
	if err != nil {
		return false, err
	}
	return capabilities.Check(requiredCapabilities), nil
}
