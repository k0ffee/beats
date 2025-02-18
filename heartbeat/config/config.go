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

// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import (
	"github.com/k0ffee/beats/v7/libbeat/autodiscover"
	conf "github.com/elastic/elastic-agent-libs/config"
)

// Config defines the structure of heartbeat.yml.
type Config struct {
	RunOnce         bool                 `config:"run_once"`
	Monitors        []*conf.C            `config:"monitors"`
	ConfigMonitors  *conf.C              `config:"config.monitors"`
	Scheduler       Scheduler            `config:"scheduler"`
	Autodiscover    *autodiscover.Config `config:"autodiscover"`
	SyntheticSuites []*conf.C            `config:"synthetic_suites"`
	Jobs            map[string]JobLimit  `config:"jobs"`
}

type JobLimit struct {
	Limit int64 `config:"limit" validate:"min=0"`
}

// Scheduler defines the syntax of a heartbeat.yml scheduler block.
type Scheduler struct {
	Limit    int64  `config:"limit"  validate:"min=0"`
	Location string `config:"location"`
}

// DefaultConfig is the canonical instantiation of Config.
var DefaultConfig = Config{}
