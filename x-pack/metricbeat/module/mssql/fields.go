// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package mssql

import (
	"github.com/k0ffee/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "mssql", asset.ModuleFieldsPri, AssetMssql); err != nil {
		panic(err)
	}
}

// AssetMssql returns asset data.
// This is the base64 encoded zlib format compressed contents of module/mssql.
func AssetMssql() string {
	return "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"
}
