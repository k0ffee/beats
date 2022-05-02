// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package suricata

import (
	"github.com/k0ffee/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "suricata", asset.ModuleFieldsPri, AssetSuricata); err != nil {
		panic(err)
	}
}

// AssetSuricata returns asset data.
// This is the base64 encoded zlib format compressed contents of module/suricata.
func AssetSuricata() string {
	return "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"
}
