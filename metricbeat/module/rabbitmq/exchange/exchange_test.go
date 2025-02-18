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

package exchange

import (
	"testing"

	mbtest "github.com/k0ffee/beats/v7/metricbeat/mb/testing"
	"github.com/k0ffee/beats/v7/metricbeat/module/rabbitmq/mtest"
	"github.com/elastic/elastic-agent-libs/mapstr"

	"github.com/stretchr/testify/assert"
)

func TestFetchEventContents(t *testing.T) {
	server := mtest.Server(t, mtest.DefaultServerConfig)
	defer server.Close()

	reporter := &mbtest.CapturingReporterV2{}

	metricSet := mbtest.NewReportingMetricSetV2Error(t, getConfig(server.URL))
	err := metricSet.Fetch(reporter)
	assert.NoError(t, err)

	e := mbtest.StandardizeEvent(metricSet, reporter.GetEvents()[0])
	t.Logf("%s/%s event: %+v", metricSet.Module().Name(), metricSet.Name(), e.Fields.StringToPrint())

	ee, _ := e.Fields.GetValue("rabbitmq.exchange")
	event := ee.(mapstr.M)

	messagesExpected := mapstr.M{
		"publish_in": mapstr.M{
			"count":   int64(100),
			"details": mapstr.M{"rate": float64(0.5)},
		},
		"publish_out": mapstr.M{
			"count":   int64(99),
			"details": mapstr.M{"rate": float64(0.9)},
		},
	}

	assert.Equal(t, "exchange.name", event["name"])
	assert.Equal(t, true, event["durable"])
	assert.Equal(t, false, event["auto_delete"])
	assert.Equal(t, false, event["internal"])
	assert.Equal(t, messagesExpected, event["messages"])
}

func TestData(t *testing.T) {
	server := mtest.Server(t, mtest.DefaultServerConfig)
	defer server.Close()

	ms := mbtest.NewReportingMetricSetV2Error(t, getConfig(server.URL))
	err := mbtest.WriteEventsReporterV2ErrorCond(ms, t, "", func(e mapstr.M) bool {
		hasIn, _ := e.HasKey("rabbitmq.exchange.messages.publish_in")
		hasOut, _ := e.HasKey("rabbitmq.exchange.messages.publish_out")
		return hasIn && hasOut
	})
	if err != nil {
		t.Fatal("write", err)
	}
}

func getConfig(url string) map[string]interface{} {
	return map[string]interface{}{
		"module":     "rabbitmq",
		"metricsets": []string{"exchange"},
		"hosts":      []string{url},
	}
}
