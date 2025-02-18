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

package redis

import (
	"time"

	rd "github.com/gomodule/redigo/redis"

	"github.com/k0ffee/beats/v7/filebeat/channel"
	"github.com/k0ffee/beats/v7/filebeat/harvester"
	"github.com/k0ffee/beats/v7/filebeat/input"
	"github.com/k0ffee/beats/v7/filebeat/input/file"
	"github.com/k0ffee/beats/v7/libbeat/common/cfgwarn"
	"github.com/k0ffee/beats/v7/libbeat/logp"
	conf "github.com/elastic/elastic-agent-libs/config"
)

func init() {
	err := input.Register("redis", NewInput)
	if err != nil {
		panic(err)
	}
}

// Input is a input for redis
type Input struct {
	started  bool
	outlet   channel.Outleter
	config   config
	cfg      *conf.C
	registry *harvester.Registry
}

// NewInput creates a new redis input
func NewInput(cfg *conf.C, connector channel.Connector, context input.Context) (input.Input, error) {
	cfgwarn.Experimental("Redis slowlog input is enabled.")

	config := defaultConfig

	err := cfg.Unpack(&config)
	if err != nil {
		return nil, err
	}

	out, err := connector.Connect(cfg)
	if err != nil {
		return nil, err
	}

	p := &Input{
		started:  false,
		outlet:   out,
		config:   config,
		cfg:      cfg,
		registry: harvester.NewRegistry(),
	}

	return p, nil
}

// LoadStates loads the states
func (p *Input) LoadStates(states []file.State) error {
	return nil
}

// Run runs the input
// Note: Filebeat is required to call the redis input's Run() method multiple times. It is expected to be called
// once initially when the input starts up and then again periodically, where the period is determined
// by the value of the `scan_frequency` setting.
// Also see https://www.elastic.co/guide/en/beats/filebeat/master/filebeat-input-redis.html#redis-scan_frequency.
func (p *Input) Run() {
	logp.Debug("redis", "Run redis input with hosts: %+v", p.config.Hosts)

	if len(p.config.Hosts) == 0 {
		logp.Err("No redis hosts configured")
		return
	}

	forwarder := harvester.NewForwarder(p.outlet)
	for _, host := range p.config.Hosts {
		pool := CreatePool(host, p.config.Password, p.config.Network,
			p.config.MaxConn, p.config.IdleTimeout, p.config.IdleTimeout)

		h, err := NewHarvester(pool.Get())
		if err != nil {
			logp.Err("Failed to create harvester: %v", err)
			continue
		}
		h.forwarder = forwarder

		if err := p.registry.Start(h); err != nil {
			logp.Err("Harvester start failed: %s", err)
		}
	}
}

// Stop stops the input and all its harvesters
func (p *Input) Stop() {
	p.registry.Stop()
	p.outlet.Close()
}

// Wait waits for the input to be completed. Not implemented.
func (p *Input) Wait() {}

// CreatePool creates a redis connection pool
// NOTE: This code is copied from the redis pool handling in metricbeat
func CreatePool(
	host, password, network string,
	maxConn int,
	idleTimeout, connTimeout time.Duration,
) *rd.Pool {
	return &rd.Pool{
		MaxIdle:     maxConn,
		IdleTimeout: idleTimeout,
		Dial: func() (rd.Conn, error) {
			c, err := rd.Dial(network, host,
				rd.DialConnectTimeout(connTimeout),
				rd.DialReadTimeout(connTimeout),
				rd.DialWriteTimeout(connTimeout))
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
	}
}
