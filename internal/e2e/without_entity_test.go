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
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"testing"
)

func TestWithoutEntity(t *testing.T) {
	ctx := context.Background()
	rContainer := startCollectorContainers(t, ctx, "emitting_collector_without_entity.yaml", "metrics")
	ms := pmetric.NewMetrics()
	mum := new(pmetric.JSONUnmarshaler)
	lines, _ := loadResultFile(ctx, rContainer, "/tmp/result.json")
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
