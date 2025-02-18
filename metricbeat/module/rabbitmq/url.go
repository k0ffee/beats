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

package rabbitmq

import (
	"github.com/k0ffee/beats/v7/metricbeat/mb/parse"
)

// Subpaths to management plugin endpoints
const (
	ConnectionsPath = "/api/connections"
	ExchangesPath   = "/api/exchanges"
	NodesPath       = "/api/nodes"
	OverviewPath    = "/api/overview"
	QueuesPath      = "/api/queues"
)

const (
	defaultScheme = "http"
	pathConfigKey = "management_path_prefix"
)

var (
	// HostParser parses host urls for RabbitMQ management plugin
	HostParser = parse.URLHostParserBuilder{
		DefaultScheme:   defaultScheme,
		PathConfigKey:   pathConfigKey,
		DefaultUsername: "guest",
		DefaultPassword: "guest",
	}.Build()
)
