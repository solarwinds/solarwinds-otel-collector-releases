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

package solarwindsentityconnector

import (
	"context"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest/plogtest"
	"github.com/solarwinds/solarwinds-otel-collector-releases/connector/solarwindsentityconnector/internal/metadata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/connector/connectortest"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/pdata/plog"
	"path/filepath"
	"testing"
)

type testLogsConsumer struct {
}

func (t *testLogsConsumer) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}

func (t *testLogsConsumer) ConsumeLogs(ctx context.Context, ld plog.Logs) error {
	return nil
}

var _ consumer.Logs = (*testLogsConsumer)(nil)

var (
	expectedEntities = []Entity{
		NewEntity("Snowflake", []string{"id1"}, []string{"attr1"}),
	}
)

func NewEntity(entityType string, ids []string, attributes []string) Entity {
	return Entity{
		Type:       entityType,
		IDs:        ids,
		Attributes: attributes,
	}
}

func TestLogsToLogs(t *testing.T) {
	factory := NewFactory()
	sink := &consumertest.LogsSink{}
	conn, err := factory.CreateLogsToLogs(context.Background(),
		connectortest.NewNopSettings(metadata.Type), &Config{Schema{Entities: expectedEntities}}, sink)
	require.NoError(t, err)
	require.NotNil(t, conn)

	require.NoError(t, conn.Start(context.Background(), componenttest.NewNopHost()))
	defer func() {
		assert.NoError(t, conn.Shutdown(context.Background()))
	}()

	testLogs, err := golden.ReadLogs(filepath.Join("testdata", "logsToLogs", "input-log.yaml"))
	assert.NoError(t, err)
	assert.NoError(t, conn.ConsumeLogs(context.Background(), testLogs))

	allLogs := sink.AllLogs()
	assert.Len(t, allLogs, 1)

	expected, err := golden.ReadLogs(filepath.Join("testdata", "logsToLogs", "expected-log.yaml"))
	assert.NoError(t, err)

	assert.NoError(t, plogtest.CompareLogs(expected, allLogs[0], plogtest.IgnoreObservedTimestamp()))
}

func TestMetricsToLogs(t *testing.T) {
	factory := NewFactory()
	sink := &consumertest.LogsSink{}
	conn, err := factory.CreateMetricsToLogs(context.Background(),
		connectortest.NewNopSettings(metadata.Type), &Config{Schema{Entities: expectedEntities}}, sink)
	require.NoError(t, err)
	require.NotNil(t, conn)

	require.NoError(t, conn.Start(context.Background(), componenttest.NewNopHost()))
	defer func() {
		assert.NoError(t, conn.Shutdown(context.Background()))
	}()

	testMetrics, err := golden.ReadMetrics(filepath.Join("testdata", "metricsToLogs", "input-metric.yaml"))
	assert.NoError(t, err)
	assert.NoError(t, conn.ConsumeMetrics(context.Background(), testMetrics))

	allLogs := sink.AllLogs()
	assert.Len(t, allLogs, 1)

	expected, err := golden.ReadLogs(filepath.Join("testdata", "metricsToLogs", "expected-log.yaml"))
	assert.NoError(t, err)

	assert.NoError(t, plogtest.CompareLogs(expected, allLogs[0], plogtest.IgnoreObservedTimestamp()))
}
