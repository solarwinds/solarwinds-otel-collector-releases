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
	"testing"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest/plogtest"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest/pmetrictest"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest/ptracetest"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
)

const (
	collectorNameAttributeName            = "sw.otelcol.collector.name"
	collectorNameAttributeValue           = "testing_collector_name"
	collectorEntityCreationAttributeName  = "sw.otelcol.collector.entity_creation"
	collectorEntityCreationAttributeValue = "on"
)

func TestMetricStream(t *testing.T) {
	ctx := context.Background()
	rContainer := startCollectorContainers(t, ctx, "emitting_collector.yaml", Metrics, collectorRunningPeriod)
	evaluateMetricsStream(t, ctx, rContainer)
}

func TestTracesStream(t *testing.T) {
	ctx := context.Background()
	rContainer := startCollectorContainers(t, ctx, "emitting_collector.yaml", Traces, collectorRunningPeriod)
	evaluateTracesStream(t, ctx, rContainer)
}

func TestLogsStream(t *testing.T) {
	ctx := context.Background()
	rContainer := startCollectorContainers(t, ctx, "emitting_collector.yaml", Logs, collectorRunningPeriod)
	evaluateLogsStream(t, ctx, rContainer)
}

func evaluateMetricsStream(
	t *testing.T,
	ctx context.Context,
	container testcontainers.Container,
) {
	lines, err := loadResultFile(ctx, container, receivingContainerResultsPath)
	require.NoError(t, err)

	heartbeatMetrics := getHeartbeatMetrics(lines)
	assertHeartbeatMetrics(t, heartbeatMetrics, "emitting_collector_heartbeat.json")

	nonHeartbeatMetrics := getNonHeartbeatMetrics(lines)
	expectedMetrics := loadExpectedMetrics(t, "test_metrics_stream.json")

	require.Equal(t, len(expectedMetrics), len(nonHeartbeatMetrics), "expected metrics count doesn't match")
	for i, metric := range nonHeartbeatMetrics {
		err := pmetrictest.CompareMetrics(expectedMetrics[i], metric, pmetrictest.IgnoreTimestamp(), pmetrictest.IgnoreMetricValues())
		require.NoError(t, err)
	}
}

func evaluateTracesStream(
	t *testing.T,
	ctx context.Context,
	container testcontainers.Container,
) {
	lines, err := loadResultFile(ctx, container, receivingContainerResultsPath)
	require.NoError(t, err)

	metrics := getMetrics(lines)
	assertHeartbeatMetrics(t, metrics, "emitting_collector_heartbeat.json")

	traces := getTraces(lines)
	expectedTraces := loadExpectedTraces(t, "test_traces_stream.json")

	require.Equal(t, len(expectedTraces), len(traces), "expected traces count doesn't match")
	for i, trace := range traces {
		maskParentSpanID(expectedTraces[i])
		maskParentSpanID(trace)
		err := ptracetest.CompareTraces(expectedTraces[i], trace, ptracetest.IgnoreStartTimestamp(), ptracetest.IgnoreEndTimestamp(), ptracetest.IgnoreSpanID(), ptracetest.IgnoreTraceID())
		require.NoError(t, err)
	}
}

func evaluateLogsStream(
	t *testing.T,
	ctx context.Context,
	container testcontainers.Container,
) {
	lines, err := loadResultFile(ctx, container, receivingContainerResultsPath)
	require.NoError(t, err)

	metrics := getMetrics(lines)
	assertHeartbeatMetrics(t, metrics, "emitting_collector_heartbeat.json")

	logs := getLogs(lines)
	expectedLogs := loadExpectedLogs(t, "test_logs_stream.json")

	require.Equal(t, len(expectedLogs), len(logs), "expected logs count doesn't match")
	for i, log := range logs {
		err := plogtest.CompareLogs(expectedLogs[i], log, plogtest.IgnoreTimestamp())
		require.NoError(t, err)
	}
}
