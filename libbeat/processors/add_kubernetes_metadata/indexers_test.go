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

package add_kubernetes_metadata

import (
	"fmt"
	"testing"

	"github.com/k0ffee/beats/v7/libbeat/common/kubernetes/metadata"
	"github.com/elastic/elastic-agent-libs/config"
	"github.com/elastic/elastic-agent-libs/mapstr"

	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"github.com/k0ffee/beats/v7/libbeat/common/kubernetes"
)

var addResourceMetadata = metadata.GetDefaultResourceMetadataConfig()
var metagen = metadata.NewPodMetadataGenerator(config.NewConfig(), nil, nil, nil, nil, addResourceMetadata)

func TestPodIndexer(t *testing.T) {
	var testConfig = config.NewConfig()

	podIndexer, err := NewPodNameIndexer(*testConfig, metagen)
	assert.NoError(t, err)

	podName := "testpod"
	uid := "005f3b90-4b9d-12f8-acf0-31020a840133"
	ns := "testns"
	nodeName := "testnode"
	pod := kubernetes.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      podName,
			UID:       types.UID(uid),
			Namespace: ns,
			Labels: map[string]string{
				"labelkey": "labelvalue",
			},
		},
		Spec: v1.PodSpec{
			NodeName: nodeName,
		},
		Status: v1.PodStatus{PodIP: "127.0.0.5"},
	}

	indexers := podIndexer.GetMetadata(&pod)
	assert.Equal(t, len(indexers), 1)
	assert.Equal(t, indexers[0].Index, fmt.Sprintf("%s/%s", ns, podName))

	expected := mapstr.M{
		"kubernetes": mapstr.M{
			"pod": mapstr.M{
				"name": "testpod",
				"uid":  "005f3b90-4b9d-12f8-acf0-31020a840133",
				"ip":   "127.0.0.5",
			},
			"labels": mapstr.M{
				"labelkey": "labelvalue",
			},
			"namespace": "testns",
			"node": mapstr.M{
				"name": "testnode",
			},
		},
	}

	assert.Equal(t, expected.String(), indexers[0].Data.String())

	indices := podIndexer.GetIndexes(&pod)
	assert.Equal(t, len(indices), 1)
	assert.Equal(t, indices[0], fmt.Sprintf("%s/%s", ns, podName))
}

func TestPodUIDIndexer(t *testing.T) {
	var testConfig = config.NewConfig()

	metaGenWithPodUID := metadata.NewPodMetadataGenerator(config.NewConfig(), nil, nil, nil, nil, addResourceMetadata)

	podUIDIndexer, err := NewPodUIDIndexer(*testConfig, metaGenWithPodUID)
	assert.NoError(t, err)

	podName := "testpod"
	ns := "testns"
	uid := "005f3b90-4b9d-12f8-acf0-31020a840133"
	nodeName := "testnode"
	pod := kubernetes.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      podName,
			Namespace: ns,
			UID:       types.UID(uid),
			Labels: map[string]string{
				"labelkey": "labelvalue",
			},
		},
		Spec: v1.PodSpec{
			NodeName: nodeName,
		},
		Status: v1.PodStatus{PodIP: "127.0.0.5"},
	}

	indexers := podUIDIndexer.GetMetadata(&pod)
	assert.Equal(t, len(indexers), 1)
	assert.Equal(t, indexers[0].Index, uid)

	expected := mapstr.M{
		"kubernetes": mapstr.M{
			"pod": mapstr.M{
				"name": "testpod",
				"uid":  "005f3b90-4b9d-12f8-acf0-31020a840133",
				"ip":   "127.0.0.5",
			},
			"namespace": "testns",
			"node": mapstr.M{
				"name": "testnode",
			},
			"labels": mapstr.M{
				"labelkey": "labelvalue",
			},
		},
	}

	assert.Equal(t, expected.String(), indexers[0].Data.String())

	indices := podUIDIndexer.GetIndexes(&pod)
	assert.Equal(t, len(indices), 1)
	assert.Equal(t, indices[0], uid)
}

func TestContainerIndexer(t *testing.T) {
	var testConfig = config.NewConfig()

	conIndexer, err := NewContainerIndexer(*testConfig, metagen)
	assert.NoError(t, err)

	podName := "testpod"
	ns := "testns"
	uid := "005f3b90-4b9d-12f8-acf0-31020a840133"
	container := "container"
	containerImage := "containerimage"
	initContainerImage := "initcontainerimage"
	initContainer := "initcontainer"
	ephemeralContainerImage := "ephemeralcontainerimage"
	ephemeralContainer := "ephemeralcontainer"
	nodeName := "testnode"

	pod := kubernetes.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      podName,
			UID:       types.UID(uid),
			Namespace: ns,
			Labels: map[string]string{
				"labelkey": "labelvalue",
			},
		},
		Status: v1.PodStatus{PodIP: "127.0.0.5"},
		Spec:   v1.PodSpec{},
	}

	indexers := conIndexer.GetMetadata(&pod)
	indices := conIndexer.GetIndexes(&pod)
	assert.Equal(t, len(indexers), 0)
	assert.Equal(t, len(indices), 0)
	expected := mapstr.M{
		"kubernetes": mapstr.M{
			"pod": mapstr.M{
				"name": "testpod",
				"uid":  "005f3b90-4b9d-12f8-acf0-31020a840133",
				"ip":   "127.0.0.5",
			},
			"namespace": "testns",
			"node": mapstr.M{
				"name": "testnode",
			},
			"labels": mapstr.M{
				"labelkey": "labelvalue",
			},
		},
	}
	container1 := "docker://abcde"
	pod.Spec.NodeName = nodeName
	pod.Status.ContainerStatuses = []kubernetes.PodContainerStatus{
		{
			Name:        container,
			Image:       containerImage,
			ContainerID: container1,
		},
	}
	container2 := "docker://fghij"
	pod.Status.InitContainerStatuses = []kubernetes.PodContainerStatus{
		{
			Name:        initContainer,
			Image:       initContainerImage,
			ContainerID: container2,
		},
	}
	container3 := "docker://klmno"
	pod.Status.EphemeralContainerStatuses = []kubernetes.PodContainerStatus{
		{
			Name:        ephemeralContainer,
			Image:       ephemeralContainerImage,
			ContainerID: container3,
		},
	}

	indexers = conIndexer.GetMetadata(&pod)
	assert.Equal(t, len(indexers), 3)
	assert.Equal(t, indexers[0].Index, "abcde")
	assert.Equal(t, indexers[1].Index, "fghij")
	assert.Equal(t, indexers[2].Index, "klmno")

	indices = conIndexer.GetIndexes(&pod)
	assert.Equal(t, len(indices), 3)
	assert.Equal(t, indices[0], "abcde")
	assert.Equal(t, indices[1], "fghij")
	assert.Equal(t, indices[2], "klmno")

	expected.Put("kubernetes.container",
		mapstr.M{
			"name":    container,
			"image":   containerImage,
			"id":      "abcde",
			"runtime": "docker",
		})
	assert.Equal(t, expected.String(), indexers[0].Data.String())

	expected.Put("kubernetes.container",
		mapstr.M{
			"name":    initContainer,
			"image":   initContainerImage,
			"id":      "fghij",
			"runtime": "docker",
		})
	assert.Equal(t, expected.String(), indexers[1].Data.String())

	expected.Put("kubernetes.container",
		mapstr.M{
			"name":    ephemeralContainer,
			"image":   ephemeralContainerImage,
			"id":      "klmno",
			"runtime": "docker",
		})
	assert.Equal(t, expected.String(), indexers[2].Data.String())
}

func TestFilteredGenMeta(t *testing.T) {
	var testConfig = config.NewConfig()

	podIndexer, err := NewPodNameIndexer(*testConfig, metagen)
	assert.NoError(t, err)

	podName := "testpod"
	ns := "testns"
	pod := kubernetes.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      podName,
			Namespace: ns,
			Labels: map[string]string{
				"foo": "bar",
				"x":   "y",
			},
			Annotations: map[string]string{
				"a": "b",
				"c": "d",
			},
		},
		Spec: v1.PodSpec{},
	}

	indexers := podIndexer.GetMetadata(&pod)
	assert.Equal(t, len(indexers), 1)

	rawLabels, _ := indexers[0].Data.GetValue("kubernetes.labels")
	assert.NotNil(t, rawLabels)

	labelMap, ok := rawLabels.(mapstr.M)
	assert.Equal(t, ok, true)
	assert.Equal(t, len(labelMap), 2)

	rawAnnotations, _ := indexers[0].Data.GetValue("kubernetes.annotations")
	assert.Nil(t, rawAnnotations)

	config, err := config.NewConfigFrom(map[string]interface{}{
		"include_annotations": []string{"a"},
		"include_labels":      []string{"foo"},
	})
	assert.NoError(t, err)

	filteredGen := metadata.NewPodMetadataGenerator(config, nil, nil, nil, nil, addResourceMetadata)

	podIndexer, err = NewPodNameIndexer(*testConfig, filteredGen)
	assert.NoError(t, err)

	indexers = podIndexer.GetMetadata(&pod)
	assert.Equal(t, len(indexers), 1)

	rawLabels, _ = indexers[0].Data.GetValue("kubernetes.labels")
	assert.NotNil(t, rawLabels)

	labelMap, ok = rawLabels.(mapstr.M)
	assert.Equal(t, ok, true)
	assert.Equal(t, len(labelMap), 1)

	ok, _ = labelMap.HasKey("foo")
	assert.Equal(t, ok, true)

	rawAnnotations, _ = indexers[0].Data.GetValue("kubernetes.annotations")
	assert.NotNil(t, rawAnnotations)
	annotationsMap, ok := rawAnnotations.(mapstr.M)

	assert.Equal(t, ok, true)
	assert.Equal(t, len(annotationsMap), 1)

	ok, _ = annotationsMap.HasKey("a")
	assert.Equal(t, ok, true)
}

func TestFilteredGenMetaExclusion(t *testing.T) {
	var testConfig = config.NewConfig()

	config, err := config.NewConfigFrom(map[string]interface{}{
		"exclude_labels": []string{"x"},
	})
	assert.NoError(t, err)

	filteredGen := metadata.NewPodMetadataGenerator(config, nil, nil, nil, nil, addResourceMetadata)

	podIndexer, err := NewPodNameIndexer(*testConfig, filteredGen)
	assert.NoError(t, err)

	podName := "testpod"
	ns := "testns"
	pod := kubernetes.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      podName,
			Namespace: ns,
			Labels: map[string]string{
				"foo": "bar",
				"x":   "y",
			},
			Annotations: map[string]string{
				"a": "b",
				"c": "d",
			},
		},
		Spec: v1.PodSpec{},
	}

	assert.NoError(t, err)

	indexers := podIndexer.GetMetadata(&pod)
	assert.Equal(t, len(indexers), 1)

	rawLabels, _ := indexers[0].Data.GetValue("kubernetes.labels")
	assert.NotNil(t, rawLabels)

	labelMap, ok := rawLabels.(mapstr.M)
	assert.Equal(t, ok, true)
	assert.Equal(t, len(labelMap), 1)

	ok, _ = labelMap.HasKey("foo")
	assert.Equal(t, ok, true)

	ok, _ = labelMap.HasKey("x")
	assert.Equal(t, ok, false)
}

func TestIpPortIndexer(t *testing.T) {
	var testConfig = config.NewConfig()

	ipIndexer, err := NewIPPortIndexer(*testConfig, metagen)
	assert.NoError(t, err)

	podName := "testpod"
	ns := "testns"
	uid := "005f3b90-4b9d-12f8-acf0-31020a840133"
	container := "container"
	containerImage := "containerimage"
	ip := "1.2.3.4"
	port := int32(80)
	pod := kubernetes.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      podName,
			UID:       types.UID(uid),
			Namespace: ns,
			Labels: map[string]string{
				"labelkey": "labelvalue",
			},
		},
		Spec: v1.PodSpec{
			Containers: make([]kubernetes.Container, 0),
		},

		Status: v1.PodStatus{
			PodIP:             ip,
			ContainerStatuses: make([]kubernetes.PodContainerStatus, 0),
		},
	}

	indexers := ipIndexer.GetMetadata(&pod)
	indices := ipIndexer.GetIndexes(&pod)

	assert.Equal(t, 1, len(indexers))
	assert.Equal(t, 1, len(indices))
	assert.Equal(t, ip, indices[0])
	assert.Equal(t, ip, indexers[0].Index)

	// Meta doesn't have container info
	_, err = indexers[0].Data.GetValue("kubernetes.container.name")
	assert.Error(t, err)

	expected := mapstr.M{
		"kubernetes": mapstr.M{
			"pod": mapstr.M{
				"name": "testpod",
				"uid":  "005f3b90-4b9d-12f8-acf0-31020a840133",
				"ip":   "1.2.3.4",
			},
			"namespace": "testns",
			"node": mapstr.M{
				"name": "testnode",
			},
			"labels": mapstr.M{
				"labelkey": "labelvalue",
			},
		},
	}

	pod.Spec.Containers = []v1.Container{
		{
			Name:  container,
			Image: containerImage,
			Ports: []v1.ContainerPort{
				{
					Name:          container,
					ContainerPort: port,
				},
			},
		},
	}
	pod.Status.ContainerStatuses = []kubernetes.PodContainerStatus{
		{
			Name:        container,
			Image:       containerImage,
			ContainerID: "docker://foobar",
		},
	}

	nodeName := "testnode"
	pod.Spec.NodeName = nodeName

	indexers = ipIndexer.GetMetadata(&pod)
	assert.Equal(t, 2, len(indexers))
	assert.Equal(t, ip, indexers[0].Index)
	assert.Equal(t, fmt.Sprintf("%s:%d", ip, port), indexers[1].Index)

	indices = ipIndexer.GetIndexes(&pod)
	assert.Equal(t, 2, len(indices))
	assert.Equal(t, ip, indices[0])
	assert.Equal(t, fmt.Sprintf("%s:%d", ip, port), indices[1])

	assert.Equal(t, expected.String(), indexers[0].Data.String())
	expected.Put("kubernetes.container",
		mapstr.M{
			"name":    container,
			"image":   containerImage,
			"id":      "foobar",
			"runtime": "docker"})
	assert.Equal(t, expected.String(), indexers[1].Data.String())
}
