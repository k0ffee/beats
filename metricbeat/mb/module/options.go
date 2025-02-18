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

package module

import (
	"time"

	"github.com/k0ffee/beats/v7/metricbeat/mb"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

// Option specifies some optional arguments used for configuring the behavior
// of a module Wrapper.
type Option func(wrapper *Wrapper)

// WithMaxStartDelay specifies the upper bound for the random startup delay
// for each MetricSet in the module. By default there is no delay.
func WithMaxStartDelay(delay time.Duration) Option {
	return func(w *Wrapper) {
		w.maxStartDelay = delay
	}
}

// WithEventModifier attaches an EventModifier that will be executed for each
// event generated by the MetricSets of the module. Multiple EventModifiers can
// be added and they will be executed in the order in which they were added.
func WithEventModifier(modifier mb.EventModifier) Option {
	return func(w *Wrapper) {
		w.eventModifiers = append(w.eventModifiers, modifier)
	}
}

// WithMetricSetInfo attaches an EventModifier that adds information about the
// MetricSet that generated the event. It will always add the metricset and
// module names. And it will add the host and rtt (round-trip time in
// microseconds) values if they are non-zero values.
//
//     "metricset": {
//       "host": "apache",
//       "module": "apache",
//       "name": "status",
//       "rtt": 115
//     }
func WithMetricSetInfo() Option {
	return WithEventModifier(mb.AddMetricSetInfo)
}

// WithServiceName sets the `service.name` field of the event with the value
// given to the `service.name` setting in the module configuration.
func WithServiceName() Option {
	return func(w *Wrapper) {
		if w.Module == nil {
			return
		}
		serviceName := w.Module.Config().ServiceName
		if serviceName == "" {
			return
		}
		modifier := func(_, _ string, event *mb.Event) {
			if event == nil {
				return
			}
			if event.RootFields == nil {
				event.RootFields = mapstr.M{}
			} else if current, err := event.RootFields.GetValue("service.name"); err == nil && current != "" {
				// Already set by the metricset, don't overwrite
				return
			}
			event.RootFields.Put("service.name", serviceName)
		}
		w.eventModifiers = append(w.eventModifiers, modifier)
	}
}
