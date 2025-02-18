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

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package aerospike

import (
	"github.com/k0ffee/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "aerospike", asset.ModuleFieldsPri, AssetAerospike); err != nil {
		panic(err)
	}
}

// AssetAerospike returns asset data.
// This is the base64 encoded zlib format compressed contents of module/aerospike.
func AssetAerospike() string {
	return "eJzUmMGO4zYMhu95CmIve2nyADkUWLSXHnZRFL0VxYCR6FgdWTQkalIDffhCdjzj2HKSycxuPDrkYEvk/8mkRGYNj9RsAclzqM0jrQDEiKUtfPrSP/u0AtAUlDe1GHZb+HkFAPD8HirW0aalnixhoC3scQVQGLI6bNvJa3BY0amjNKSp03TPsT4+yXg6NTU0l35DjYqe3+RMzprtRs7IKUk/xjKGUpQ15OTk1ZyWC3rS+KW1BkFQwmb0NqdiqESTJaHJ63NqrlA0UNU5APHoAqq0IOSVnlM7VEzes8/O6EVbdvuZCVfoTuNbrHbkgYvjl8pSSIkCBRpLGg5GSkDXicuRDcKQ5aHg6PRiGDyFaIU0GAeY9EGr7zxGiEpRCD8M4uiviPYMz3nJYiriOE677yf50r4nPRo4ylR2L9kT5uLkvZIzmf/oqTlm6FRcCIW7JeFE7SQDP0z+TYJnkdmX3/ArUu/gzXe9GFv7Hz35MhA3Xov3j+cpyyIDem7LZ0L6pdB7MmoczzeXnL+a8AhB2OOeurB9Zd2JT2gs7ixtapXbwk5ZUGhJPxSWMTepYF+hbKEmr6b19BUYaXwlDNFT2kWCyjhTxQoUOzH7yDGAblFTyQ+oPIcAaG37NBwLpr4lmD9LCk/3B/29W5y+GBcdlsIalZGmFZi8gJQmXEMkLGg3u0Yol7Jng7+HmVt8BcqfyXtn4Jnl+ImsZYXpKhUewUB7+KcnrM+QxUB6gWBJFuya65l6nvJQPew8oSppXO10NDtmS+hedwD8VoD4SD8NOvsSA/SO4HNp9uX6gEJ+/VfC+K+iin3z97pW8vliqPXiu0XvdWp9ba2dnluvbZgXmMrdJl1M5tfEv0bBuyXBl4qjkwEZKxVr08V/UvZGOOM0/btMunQHtfLeiBgWzBhIsdPom46UwhtZF3UVHXlvPq3Tb/a4e6TmwH58gl8Q+O3Z88TuSyesaVNyyP8beZtX1gQTkycO3xkzOZwl5N0/pCYF6s2XyB+k2OuZZu3S7VFhEMp1aWfD9IowfGkT/FHfMNbgUBpVAvpUw4p5oqOQTIdzUuD9EKHGjdPk5TyYzZMgXD+0rVD+w76lqBnLMQFU9J6c2Kb9eyYVmQfj9l0vFjar/wMAAP//lvJazw=="
}
