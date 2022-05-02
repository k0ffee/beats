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

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package docker

import (
	"github.com/k0ffee/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "docker", asset.ModuleFieldsPri, AssetDocker); err != nil {
		panic(err)
	}
}

// AssetDocker returns asset data.
// This is the base64 encoded zlib format compressed contents of module/docker.
func AssetDocker() string {
	return "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"
}
