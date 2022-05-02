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

package redis

import (
	"github.com/k0ffee/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "redis", asset.ModuleFieldsPri, AssetRedis); err != nil {
		panic(err)
	}
}

// AssetRedis returns asset data.
// This is the base64 encoded zlib format compressed contents of module/redis.
func AssetRedis() string {
	return "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"
}
