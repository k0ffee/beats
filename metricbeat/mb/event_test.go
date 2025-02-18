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

//go:build !integration
// +build !integration

package mb

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/k0ffee/beats/v7/libbeat/beat/events"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

func TestEventConversionToBeatEvent(t *testing.T) {
	var (
		timestamp = time.Now()
		module    = "docker"
		metricSet = "uptime"
	)

	t.Run("all levels", func(t *testing.T) {
		e := (&Event{
			Timestamp: timestamp,
			RootFields: mapstr.M{
				"type": "docker",
			},
			ModuleFields: mapstr.M{
				"container": mapstr.M{
					"name": "wordpress",
				},
			},
			MetricSetFields: mapstr.M{
				"ms": 1000,
			},
		}).BeatEvent(module, metricSet)

		assert.Equal(t, timestamp, e.Timestamp)
		assert.Equal(t, mapstr.M{
			"type": "docker",
			"docker": mapstr.M{
				"container": mapstr.M{
					"name": "wordpress",
				},
				"uptime": mapstr.M{
					"ms": 1000,
				},
			},
			"service": mapstr.M{
				"type": "docker",
			},
		}, e.Fields)
	})

	t.Run("idempotent", func(t *testing.T) {
		mbEvent := &Event{
			Timestamp: timestamp,
			RootFields: mapstr.M{
				"type": "docker",
			},
			ModuleFields: mapstr.M{
				"container": mapstr.M{
					"name": "wordpress",
				},
			},
			MetricSetFields: mapstr.M{
				"ms": 1000,
			},
		}
		e := mbEvent.BeatEvent(module, metricSet)
		e = mbEvent.BeatEvent(module, metricSet)

		assert.Equal(t, timestamp, e.Timestamp)
		assert.Equal(t, mapstr.M{
			"type": "docker",
			"docker": mapstr.M{
				"container": mapstr.M{
					"name": "wordpress",
				},
				"uptime": mapstr.M{
					"ms": 1000,
				},
			},
			"service": mapstr.M{
				"type": "docker",
			},
		}, e.Fields)
	})

	t.Run("with event modifiers", func(t *testing.T) {
		modifier := func(m, ms string, e *Event) {
			e.RootFields.Put("module", m)
			e.RootFields.Put("metricset", ms)
		}

		e := (&Event{}).BeatEvent(module, metricSet, modifier)
		assert.Equal(t, mapstr.M{
			"module":    module,
			"metricset": metricSet,
			"service": mapstr.M{
				"type": "docker",
			},
		}, e.Fields)
	})

	t.Run("with ID", func(t *testing.T) {
		mbEvent := &Event{
			ID:        "foobar",
			Timestamp: timestamp,
			RootFields: mapstr.M{
				"type": "docker",
			},
			ModuleFields: mapstr.M{
				"container": mapstr.M{
					"name": "wordpress",
				},
			},
			MetricSetFields: mapstr.M{
				"ms": 1000,
			},
		}
		e := mbEvent.BeatEvent(module, metricSet)
		e = mbEvent.BeatEvent(module, metricSet)

		assert.Equal(t, "foobar", e.Meta[events.FieldMetaID])
		assert.Equal(t, timestamp, e.Timestamp)
		assert.Equal(t, mapstr.M{
			"type": "docker",
			"docker": mapstr.M{
				"container": mapstr.M{
					"name": "wordpress",
				},
				"uptime": mapstr.M{
					"ms": 1000,
				},
			},
			"service": mapstr.M{
				"type": "docker",
			},
		}, e.Fields)
	})

	t.Run("error message", func(t *testing.T) {
		msg := "something failed"
		e := (&Event{
			Error: errors.New(msg),
		}).BeatEvent(module, metricSet)

		errorMessage, err := e.Fields.GetValue("error.message")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, msg, errorMessage)
	})
}

func TestAddMetricSetInfo(t *testing.T) {
	const (
		host    = "localhost"
		elapsed = time.Duration(500 * time.Millisecond)
	)

	t.Run("all fields", func(t *testing.T) {
		e := Event{
			Host: host,
			Took: elapsed,
		}

		AddMetricSetInfo(moduleName, metricSetName, &e)

		assert.Equal(t, mapstr.M{
			"event": mapstr.M{
				"module":   moduleName,
				"dataset":  moduleName + "." + metricSetName,
				"duration": time.Duration(500000000),
			},
			"service": mapstr.M{
				"address": host,
			},
			"metricset": mapstr.M{
				"name": metricSetName,
			},
		}, e.RootFields)
	})

	t.Run("no optional fields", func(t *testing.T) {
		e := Event{}

		AddMetricSetInfo(moduleName, metricSetName, &e)

		assert.Equal(t, mapstr.M{
			"event": mapstr.M{
				"module":  moduleName,
				"dataset": moduleName + "." + metricSetName,
			},
			"metricset": mapstr.M{
				"name": metricSetName,
			},
		}, e.RootFields)
	})
}

func TestTransformMapStrToEvent(t *testing.T) {
	var (
		timestamp  = time.Now()
		took       = time.Duration(1)
		moduleData = mapstr.M{
			"container_id": "busybox",
		}
		metricSetData = mapstr.M{
			"uptime": "1 day",
		}
		failure = errors.New("failed")
	)

	m := mapstr.M{
		TimestampKey:  timestamp,
		RTTKey:        took,
		ModuleDataKey: moduleData,
	}
	m.DeepUpdate(metricSetData)

	t.Run("normal", func(t *testing.T) {
		m := m.Clone()
		e := TransformMapStrToEvent("module", m, failure)

		assert.Equal(t, timestamp, e.Timestamp)
		assert.Equal(t, took, e.Took)
		assert.Empty(t, e.RootFields)
		assert.Equal(t, moduleData, e.ModuleFields)
		assert.Equal(t, metricSetData, e.MetricSetFields)
		assert.Equal(t, failure, e.Error)
	})

	t.Run("namespace", func(t *testing.T) {
		const namespace = "foo.bar"

		mapWithNamespace := m.Clone()
		mapWithNamespace.Put(NamespaceKey, namespace)

		e := TransformMapStrToEvent("module", mapWithNamespace, nil)

		assert.Equal(t, timestamp, e.Timestamp)
		assert.Equal(t, took, e.Took)
		assert.Empty(t, e.RootFields)
		assert.Equal(t, moduleData, e.ModuleFields)
		assert.Equal(t, "module."+namespace, e.Namespace)
		assert.Equal(t, metricSetData, e.MetricSetFields)
	})
}
