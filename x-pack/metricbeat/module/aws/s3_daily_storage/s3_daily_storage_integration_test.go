// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

//go:build integration && aws
// +build integration,aws

package s3_daily_storage

import (
	"testing"

	_ "github.com/k0ffee/beats/v7/libbeat/processors/actions"
	mbtest "github.com/k0ffee/beats/v7/metricbeat/mb/testing"
	"github.com/k0ffee/beats/v7/x-pack/metricbeat/module/aws/mtest"
)

func TestData(t *testing.T) {
	config := mtest.GetConfigForTest(t, "s3_daily_storage", "86400s")

	metricSet := mbtest.NewFetcher(t, config)
	metricSet.WriteEvents(t, "/")
}
