// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package envoyproxy

import (
	"github.com/k0ffee/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "envoyproxy", asset.ModuleFieldsPri, AssetEnvoyproxy); err != nil {
		panic(err)
	}
}

// AssetEnvoyproxy returns asset data.
// This is the base64 encoded zlib format compressed contents of module/envoyproxy.
func AssetEnvoyproxy() string {
	return "eJysk8Fu2zAMhu9+iv8BlhfwYcCQZcAOuyzY2dAs2hEii5pEZdOevpDluKnbokBaHpnw+37T9A5nyi3IXTj7wP9yA4gRSy0Otz1NsQ/Gi2HX4nMDAD9YJ0sYOOCknLbGjbA8RvjAOvWk8TtXcAMMhqyO7Ty4g1MTbaSlJHtqMQZOfum8oC31baZhCDxVCmZM1atBKMBxmJQ1/1WZXQZvQzzGsDx2xbzSa4wz5b8c9Np9JQrqngplHvy0mG3Gl/3+cDw2G12g6NlF6garxniv9OdCQaVsHMlHCaSmLlK4mJ46MdP2+Sy7cW0NJbO00CncLgwwzifprj875ThSz07HtyP+WkJgCYESAsY9oTxbzp9EUTqj713M96/gAXKiK2urUElOHIzk973venArrN7W1jX/6QOOq8rqeUnvUT45Ed+UeggAAP//R+geWQ=="
}
