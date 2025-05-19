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
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest/pmetrictest"
	"github.com/solarwinds/solarwinds-otel-collector-contrib/pkg/version"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

const (
	resourceAttributeName  = "resource.attributes.testing_attribute"
	resourceAttributeValue = "testing_value"
	heartbeatMetricName    = "sw.otelcol.uptime"
)

func isHeartbeatMetrics(ms pmetric.Metrics) bool {
	return ms.ResourceMetrics().Len() == 1 &&
		ms.ResourceMetrics().At(0).ScopeMetrics().Len() == 1 &&
		ms.ResourceMetrics().At(0).ScopeMetrics().At(0).Metrics().Len() == 1 &&
		ms.ResourceMetrics().At(0).ScopeMetrics().At(0).Metrics().At(0).Name() == heartbeatMetricName
}

func loadResultFile(
	ctx context.Context,
	container testcontainers.Container,
	resultFilePath string,
) ([]string, error) {
	r, err := container.CopyFileFromContainer(ctx, resultFilePath)
	if err != nil {
		return make([]string, 0), err
	}
	defer r.Close()

	content, err := io.ReadAll(r)
	if err != nil {
		return make([]string, 0), err
	}

	log.Print("*** raw result content:\n" + string(content) + "\n")
	lines := strings.Split(string(content), "\n")
	return lines, nil
}

func loadExpectedResultFile(
	filename string,
) ([]string, error) {
	expectedDataPath, err := filepath.Abs(filepath.Join(".", "testdata", "expected", filename))
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(expectedDataPath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	return lines, nil
}

func getMetricsWithFilter(lines []string, filter func(pmetric.Metrics) bool) []pmetric.Metrics {
	res := make([]pmetric.Metrics, 0, len(lines))
	jum := new(pmetric.JSONUnmarshaler)
	for _, line := range lines {
		metric, err := jum.UnmarshalMetrics([]byte(line))
		if err == nil && metric.ResourceMetrics().Len() > 0 && filter(metric) {
			res = append(res, metric)
		}
	}
	return res
}

func getMetrics(lines []string) []pmetric.Metrics {
	return getMetricsWithFilter(lines, func(pmetric.Metrics) bool { return true })
}

func getHeartbeatMetrics(lines []string) []pmetric.Metrics {
	return getMetricsWithFilter(lines, isHeartbeatMetrics)
}

func getNonHeartbeatMetrics(lines []string) []pmetric.Metrics {
	return getMetricsWithFilter(lines, func(m pmetric.Metrics) bool {
		return !isHeartbeatMetrics(m)
	})
}

func getLogs(lines []string) []plog.Logs {
	res := make([]plog.Logs, 0, len(lines))
	jum := new(plog.JSONUnmarshaler)
	for _, line := range lines {
		log, err := jum.UnmarshalLogs([]byte(line))
		if err == nil && log.ResourceLogs().Len() > 0 {
			res = append(res, log)
		}
	}
	return res
}

func getTraces(lines []string) []ptrace.Traces {
	res := make([]ptrace.Traces, 0, len(lines))
	jum := new(ptrace.JSONUnmarshaler)
	for _, line := range lines {
		trace, err := jum.UnmarshalTraces([]byte(line))
		if err == nil && trace.ResourceSpans().Len() > 0 {
			res = append(res, trace)
		}
	}
	return res
}

// maskSpanID replaces all parent span IDs with an empty value. It is a workaround for missing ptracetest.IgnoreParentSpanID().
func maskParentSpanID(traces ptrace.Traces) {
	spanID := pcommon.NewSpanIDEmpty()
	for i := 0; i < traces.ResourceSpans().Len(); i++ {
		rs := traces.ResourceSpans().At(i)
		for j := 0; j < rs.ScopeSpans().Len(); j++ {
			ss := rs.ScopeSpans().At(j)
			for k := 0; k < ss.Spans().Len(); k++ {
				span := ss.Spans().At(k)
				span.SetParentSpanID(spanID)
			}
		}
	}
}

func loadExpectedMetrics(t *testing.T, filename string) []pmetric.Metrics {
	lines, err := loadExpectedResultFile(filename)
	require.NoError(t, err)
	return getMetrics(lines)
}

func loadExpectedLogs(t *testing.T, filename string) []plog.Logs {
	lines, err := loadExpectedResultFile(filename)
	require.NoError(t, err)
	return getLogs(lines)
}

func loadExpectedTraces(t *testing.T, filename string) []ptrace.Traces {
	lines, err := loadExpectedResultFile(filename)
	require.NoError(t, err)
	return getTraces(lines)
}

func assertHeartbeatMetrics(
	t *testing.T,
	metrics []pmetric.Metrics,
	expectedHeartbeatMetricFilename string,
) {
	expectedMetric := loadExpectedMetrics(t, expectedHeartbeatMetricFilename)[0]
	require.True(t, isHeartbeatMetrics(expectedMetric), "metric must be a heartbeat metric")
	if v, ok := expectedMetric.ResourceMetrics().At(0).Resource().Attributes().Get("sw.otelcol.collector.version"); ok {
		v.SetStr(version.Version)
	}
	expectedMetric.ResourceMetrics().At(0).ScopeMetrics().At(0).Scope().SetVersion(version.Version)

	require.GreaterOrEqual(t, len(metrics), 1, "there must be at least one heartbeat metric")
	for _, metric := range metrics {
		err := pmetrictest.CompareMetrics(expectedMetric, metric, pmetrictest.IgnoreTimestamp(), pmetrictest.IgnoreMetricValues())
		require.NoError(t, err)
	}
}
