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

//go:build e2e

package e2e

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/network"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func TestWithoutEntity(t *testing.T) {
	ctx := context.Background()

	net, err := network.New(ctx)
	require.NoError(t, err)
	testcontainers.CleanupNetwork(t, net)

	certPath := t.TempDir()
	_, err = generateCertificates(receivingContainer, certPath)
	require.NoError(t, err)

	rContainer, err := runReceivingSolarWindsOTELCollector(ctx, certPath, net.Name)
	require.NoError(t, err)
	testcontainers.CleanupContainer(t, rContainer)

	eContainer, err := runTestedSolarWindsOTELCollector(ctx, certPath, net.Name, "emitting_collector_without_entity.yaml")
	require.NoError(t, err)
	testcontainers.CleanupContainer(t, eContainer)

	cmd := []string{
		"metrics",
		"--metrics", strconv.Itoa(samplesCount),
		"--otlp-insecure",
		"--otlp-endpoint", fmt.Sprintf("%s:%d", testedContainer, port),
		"--otlp-attributes", fmt.Sprintf("%s=\"%s\"", resourceAttributeName, resourceAttributeValue),
	}

	gContainer, err := runGeneratorContainer(ctx, net.Name, cmd)
	require.NoError(t, err)
	testcontainers.CleanupContainer(t, gContainer)

	<-time.After(collectorRunningPeriod)

	ms := pmetric.NewMetrics()
	mum := new(pmetric.JSONUnmarshaler)
	lines, err := loadResultFile(ctx, rContainer, "/tmp/result.json")
	for _, line := range lines {
		// Metrics to process.
		m, err := mum.UnmarshalMetrics([]byte(line))
		if err == nil && m.ResourceMetrics().Len() != 0 {
			m.ResourceMetrics().MoveAndAppendTo(ms.ResourceMetrics())
			continue
		}
	}
	evaluateHeartbeatMetricHasEntityCreationAsOff(t, ms)
}

func evaluateHeartbeatMetricHasEntityCreationAsOff(
	t *testing.T,
	ms pmetric.Metrics,
) {
	require.GreaterOrEqual(t, ms.ResourceMetrics().Len(), 1, "there must be at least one metric")
	atts := ms.ResourceMetrics().At(0).Resource().Attributes()

	v, available := atts.Get("sw.otelcol.collector.entity_creation")
	require.True(t, available, "sw.otelcol.collector.entity_creation resource attribute must be available")
	require.Equal(t, "off", v.AsString(), "attribute value must be the same")
}
