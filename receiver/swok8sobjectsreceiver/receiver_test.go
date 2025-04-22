// Copyright 2025 SolarWinds Worldwide, LLC. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Source: https://github.com/open-telemetry/opentelemetry-collector-contrib
// Changes customizing the original source code

package swok8sobjectsreceiver

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	apiWatch "k8s.io/apimachinery/pkg/watch"
)

func TestNewReceiver(t *testing.T) {
	t.Parallel()

	rCfg := createDefaultConfig().(*Config)
	rCfg.makeDynamicClient = newMockDynamicClient().getMockDynamicClient
	r, err := newReceiver(
		receivertest.NewNopSettings(receivertest.NopType),
		rCfg,
		consumertest.NewNop(),
	)

	require.NoError(t, err)
	require.NotNil(t, r)
	require.NoError(t, r.Start(context.Background(), componenttest.NewNopHost()))
	assert.NoError(t, r.Shutdown(context.Background()))
}

func TestPullObject(t *testing.T) {
	t.Parallel()

	mockClient := newMockDynamicClient()
	mockClient.createPods(
		generatePod("pod1", "default", map[string]any{
			"environment": "production",
		}, "1"),
		generatePod("pod2", "default", map[string]any{
			"environment": "test",
		}, "2"),
		generatePod("pod3", "default_ignore", map[string]any{
			"environment": "production",
		}, "3"),
	)

	rCfg := createDefaultConfig().(*Config)
	rCfg.makeDynamicClient = mockClient.getMockDynamicClient
	rCfg.makeDiscoveryClient = getMockDiscoveryClient

	rCfg.Objects = []*K8sObjectsConfig{
		{
			Name:          "pods",
			Mode:          PullMode,
			Interval:      time.Second * 30,
			LabelSelector: "environment=production",
		},
	}

	err := rCfg.Validate()
	require.NoError(t, err)

	consumer := newMockLogConsumer()
	r, err := newReceiver(
		receivertest.NewNopSettings(receivertest.NopType),
		rCfg,
		consumer,
	)
	require.NoError(t, err)
	require.NotNil(t, r)
	require.NoError(t, r.Start(context.Background(), componenttest.NewNopHost()))
	time.Sleep(time.Second)
	assert.Len(t, consumer.Logs(), 1)
	assert.Equal(t, 2, consumer.Count())
	assert.NoError(t, r.Shutdown(context.Background()))
}

func TestWatchObject(t *testing.T) {
	t.Parallel()

	mockClient := newMockDynamicClient()

	mockClient.createPods(
		generatePod("pod1", "default", map[string]any{
			"environment": "production",
		}, "1"),
	)

	rCfg := createDefaultConfig().(*Config)
	rCfg.makeDynamicClient = mockClient.getMockDynamicClient
	rCfg.makeDiscoveryClient = getMockDiscoveryClient

	rCfg.Objects = []*K8sObjectsConfig{
		{
			Name:       "pods",
			Mode:       WatchMode,
			Namespaces: []string{"default"},
		},
	}

	err := rCfg.Validate()
	require.NoError(t, err)

	consumer := newMockLogConsumer()
	r, err := newReceiver(
		receivertest.NewNopSettings(receivertest.NopType),
		rCfg,
		consumer,
	)

	ctx := context.Background()
	require.NoError(t, err)
	require.NotNil(t, r)
	require.NoError(t, r.Start(ctx, componenttest.NewNopHost()))

	time.Sleep(time.Millisecond * 100)
	assert.Len(t, consumer.Logs(), 1)
	assert.Equal(t, 1, consumer.Count())

	mockClient.createPods(
		generatePod("pod2", "default", map[string]any{
			"environment": "test",
		}, "2"),
		generatePod("pod3", "default_ignore", map[string]any{
			"environment": "production",
		}, "3"),
		generatePod("pod4", "default", map[string]any{
			"environment": "production",
		}, "4"),
	)
	time.Sleep(time.Millisecond * 100)
	assert.Len(t, consumer.Logs(), 3)
	assert.Equal(t, 3, consumer.Count())

	mockClient.deletePods(
		generatePod("pod2", "default", map[string]any{
			"environment": "test",
		}, "2"),
	)
	time.Sleep(time.Millisecond * 100)
	assert.Len(t, consumer.Logs(), 3)
	assert.Equal(t, 3, consumer.Count())

	assert.NoError(t, r.Shutdown(ctx))
}

func TestExcludeDeletedTrue(t *testing.T) {
	t.Parallel()

	mockClient := newMockDynamicClient()

	mockClient.createPods(
		generatePod("pod1", "default", map[string]any{
			"environment": "production",
		}, "1"),
	)

	rCfg := createDefaultConfig().(*Config)
	rCfg.makeDynamicClient = mockClient.getMockDynamicClient
	rCfg.makeDiscoveryClient = getMockDiscoveryClient

	rCfg.Objects = []*K8sObjectsConfig{
		{
			Name:       "pods",
			Mode:       WatchMode,
			Namespaces: []string{"default"},
			ExcludeWatchType: []apiWatch.EventType{
				apiWatch.Deleted,
			},
		},
	}

	err := rCfg.Validate()
	require.NoError(t, err)

	consumer := newMockLogConsumer()
	r, err := newReceiver(
		receivertest.NewNopSettings(receivertest.NopType),
		rCfg,
		consumer,
	)

	ctx := context.Background()
	require.NoError(t, err)
	require.NotNil(t, r)
	require.NoError(t, r.Start(ctx, componenttest.NewNopHost()))

	time.Sleep(time.Millisecond * 100)
	assert.Len(t, consumer.Logs(), 1)
	assert.Equal(t, 1, consumer.Count())

	mockClient.deletePods(
		generatePod("pod1", "default", map[string]any{
			"environment": "test",
		}, "1"),
	)
	time.Sleep(time.Millisecond * 100)
	assert.Len(t, consumer.Logs(), 1)
	assert.Equal(t, 1, consumer.Count())

	assert.NoError(t, r.Shutdown(ctx))
}

func TestGetObjectHashesRemovesTimestamps(t *testing.T) {
	t.Parallel()

	objWithTimestamp := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]interface{}{
				"name": "test-configmap",
				"annotations": map[string]interface{}{
					"time": "2025-03-04T09:11:59Z",
				},
			},
			"spec": map[string]interface{}{
				"dummy": "value",
			},
			"status": map[string]interface{}{
				"dummy": "value",
			},
		},
	}

	objWithoutTimestamp := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]interface{}{
				"name": "test-configmap",
				"annotations": map[string]interface{}{
					"time": "",
				},
			},
			"spec": map[string]interface{}{
				"dummy": "value",
			},
			"status": map[string]interface{}{
				"dummy": "value",
			},
		},
	}

	hashWith, err := getObjectHashes(objWithTimestamp)
	require.NoError(t, err)
	hashWithout, err := getObjectHashes(objWithoutTimestamp)
	require.NoError(t, err)

	assert.Equal(t, hashWithout.Metadata, hashWith.Metadata)
	assert.Equal(t, hashWithout.Spec, hashWith.Spec)
	assert.Equal(t, hashWithout.Status, hashWith.Status)
	assert.Equal(t, hashWithout.Other, hashWith.Other)
}
