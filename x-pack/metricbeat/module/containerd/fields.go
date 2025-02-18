// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package containerd

import (
	"github.com/k0ffee/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "containerd", asset.ModuleFieldsPri, AssetContainerd); err != nil {
		panic(err)
	}
}

// AssetContainerd returns asset data.
// This is the base64 encoded zlib format compressed contents of module/containerd.
func AssetContainerd() string {
	return "eJzUmE1v8zYMx+/5FEQvAwY0u+cwYC1QoBi6Fdt6LmSZTrXoxdBL0/TTP5Bsp44t20qeBnF0KAq/kD/Sf5JSbmGDuxVQJS1hEnW+ALDMclzBzf3+4s0CQCNHYnAFGVqyAMjRUM1Ky5Rcwe8LAICvF8BYYg1QxTlSizkUWolDLwVDnptVePEWJBHYwfDL7kpcwVorV9ZXIm79epSF0oL4y0Bk5Z8Zy6gBkilnW6Z/MaCdlEyuvy6aZW2pTdUm839NSSju7zRwG9xt1R54BLGTob7FxlfGN0z1/LST4Ff3eyR4v+OKbuDxt79BoNWM7qOORd4myvGdHYQ+Fv4EhF9/EYGgCsgCUMR641gj6ZqOZyPB6R+UOuE48XL0dg3kTnsV2DcEzopA5P/fq6JjIpahNqwqTe9ew8uVXEduTiCHXDmRofZsJ0G3ZLWzeDRgVVXDLycEcOdfDfDHsTfcW81sXHs/KYJg+OpUcBr1XGTg6S3K05RgnBBE787XEPzkuE5V+KauStRhAl6tOkKTaD7CUTLpDVFauskRmjao/3HSMoFw//wSm5tDc3hsnpqdsSiWVlnCo2rOlct4t+lNpPE/bw2cQV3tgIKPQB3wTYnSApNgkCqZH8TwReYMWac32ynVb1BL7EY4bnLMbNu0jAkNxhPYrAQ9+uVT92eIoEoLSCJVnb3BkH3+rzngF6+f9HBjEr6qeG1dNUcFvCypHQzaUMIxfy24IrGHmmZZoqYoY08k0D9XL3to3xJDDPtKl94DZ5+YQ7YL/VLu54R/iCodadOHNTu3CH1cTNZwEM5NaXEuRyt1pmGGJv4tQdLSLX9dRmupClJl/2M0B9WN14lqSwjRU/bqy0dp35gJgL3hLVCog33etx2Be5bHxvRW6Q2Ta4M2opNJjYzrYyJxT4GzIQCDtrFD1hgf3dp0v/Lgpm1sw5a0y6jSCBoNy/22wvMZ9jlARqhl7/jAeM/Z+Qgrn1AwjpWROBol9C2+3zkHVPA2hsPkJXLVeJ3M1nG7wwnfL6En1EoKvxd229jU9lKQjwscVZ7IR0Md8jHcfC89XOpK4FzRcMitqYcaCSTv686X3IP2MpHegjC+pMpF05OU5QSgB8I4BCeoh1E4E2yY4oxabCWqgogXbvQgdnLl1oei6yvdBjy5gi9VBjWonVs11FzFrIviUJ4jVWG2pKv8k2vi3y0pr7AiKuzZ10PAnF01BKp510Jblp1KWPwIAAD//wv4XH8="
}
