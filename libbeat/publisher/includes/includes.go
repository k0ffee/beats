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

package includes

import (
	// import queue types
	_ "github.com/k0ffee/beats/v7/libbeat/outputs/codec/format"
	_ "github.com/k0ffee/beats/v7/libbeat/outputs/codec/json"
	_ "github.com/k0ffee/beats/v7/libbeat/outputs/console"
	_ "github.com/k0ffee/beats/v7/libbeat/outputs/elasticsearch"
	_ "github.com/k0ffee/beats/v7/libbeat/outputs/fileout"
	_ "github.com/k0ffee/beats/v7/libbeat/outputs/kafka"
	_ "github.com/k0ffee/beats/v7/libbeat/outputs/logstash"
	_ "github.com/k0ffee/beats/v7/libbeat/outputs/redis"
	_ "github.com/k0ffee/beats/v7/libbeat/outputs/shipper"
	_ "github.com/k0ffee/beats/v7/libbeat/publisher/queue/diskqueue"
	_ "github.com/k0ffee/beats/v7/libbeat/publisher/queue/memqueue"
)
