// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package cockroachdb

import (
	"github.com/k0ffee/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "cockroachdb", asset.ModuleFieldsPri, AssetCockroachdb); err != nil {
		panic(err)
	}
}

// AssetCockroachdb returns asset data.
// This is the base64 encoded zlib format compressed contents of module/cockroachdb.
func AssetCockroachdb() string {
	return "eJxczjsOwjAQhOHepxi5SUMusAUFcAtE4dhLYsWxLe+myO0RLykw5Uif9PeYeSP44udWnJ/CYACNmpjQnb/v5dQZoHFiJ0wYWJ0BAotvsWosmXA0ALATWEpYExtAWDXmUQhXK5LsAXZSrfZmgHvkFIReuEd2C//HPKdbZcLYylo/z9697W/dIwAA//+zakEu"
}
