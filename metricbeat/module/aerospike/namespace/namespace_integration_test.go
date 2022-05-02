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

//go:build integration
// +build integration

package namespace

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/k0ffee/beats/v7/libbeat/tests/compose"
	mbtest "github.com/k0ffee/beats/v7/metricbeat/mb/testing"
)

func TestData(t *testing.T) {
	service := compose.EnsureUp(t, "aerospike")

	f := mbtest.NewReportingMetricSetV2Error(t, getConfig(service.Host()))
	if err := mbtest.WriteEventsReporterV2Error(f, t, ""); err != nil {
		t.Fatal("write", err)
	}
}

func TestFetch(t *testing.T) {
	service := compose.EnsureUp(t, "aerospike")

	f := mbtest.NewReportingMetricSetV2Error(t, getConfig(service.Host()))
	events, errs := mbtest.ReportingFetchV2Error(f)
	if len(errs) > 0 {
		t.Fatalf("Expected 0 error, had %d. %v\n", len(errs), errs)
	}
	assert.NotEmpty(t, events)
	event := events[0].MetricSetFields

	t.Logf("%s/%s event: %+v", f.Module().Name(), f.Name(), event)
}

func getConfig(host string) map[string]interface{} {
	return map[string]interface{}{
		"module":     "aerospike",
		"metricsets": []string{"namespace"},
		"hosts":      []string{host},
	}
}
