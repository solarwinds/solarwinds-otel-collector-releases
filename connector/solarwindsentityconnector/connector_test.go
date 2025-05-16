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
	"path/filepath"
	"testing"

	"github.com/solarwinds/solarwinds-otel-collector-releases/connector/solarwindsentityconnector/config"

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

// Test configuration for the entites, relationships and events.
var (
	configuredEntities = []config.Entity{
		{Type: "Snowflake", IDs: []string{"snowflake.id"}, Attributes: []string{"attr1"}},
		{Type: "AWS EC2", IDs: []string{"aws.ec2.id", "aws.ec2.name"}, Attributes: []string{"attr2"}},
	}

	configuredRelationships = []config.Relationship{
		{
			Type:        "MemberOf",
			Source:      "Snowflake",
			Destination: "AWS EC2",
			Attributes:  []string{},
		},
		{
			Type:        "TestRelationshipType",
			Source:      "AWS EC2",
			Destination: "AWS EC2",
			Attributes:  []string{},
		},
	}

	configuredEvents = config.Events{
		Relationships: configuredRelationships,
	}
)

func TestLogsToLogs(t *testing.T) {

	testCases := []struct {
		name         string
		inputFile    string
		expectedFile string
	}{
		{
			name:         "when relationship for different type is inferred log event is sent",
			inputFile:    "relationship/input-log-different-type-relationship.yaml",
			expectedFile: "relationship/expected-log-different-type-relationship.yaml",
		},
		{
			name:         "when relationship for same type is inferred log event is sent",
			inputFile:    "relationship/input-log-same-type-relationship.yaml",
			expectedFile: "relationship/expected-log-same-type-relationship.yaml",
		},
		{
			name:         "when relationship for same type is not inferred no log is sent",
			inputFile:    "relationship/input-log-same-type-relationship-nomatch.yaml",
			expectedFile: "relationship/expected-log-same-type-relationship-nomatch.yaml",
		},
		{
			name:         "when entity is inferred log event is sent",
			inputFile:    "entity/input-log.yaml",
			expectedFile: "entity/expected-log.yaml",
		},
		{
			name:      "when entity is not inferred no log is sent",
			inputFile: "entity/input-log-nomatch.yaml",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			factory := NewFactory()
			sink := &consumertest.LogsSink{}
			conn, err := factory.CreateLogsToLogs(context.Background(),
				connectortest.NewNopSettings(metadata.Type), &Config{
					Schema: config.Schema{
						Entities: configuredEntities,
						Events:   configuredEvents,
					},
					SourcePrefix:      "src.",
					DestinationPrefix: "dst.",
				}, sink)
			require.NoError(t, err)
			require.NotNil(t, conn)

			require.NoError(t, conn.Start(context.Background(), componenttest.NewNopHost()))
			defer func() {
				assert.NoError(t, conn.Shutdown(context.Background()))
			}()

			testLogs, err := golden.ReadLogs(filepath.Join("testdata", "logsToLogs", tc.inputFile))

			assert.NoError(t, err)
			assert.NoError(t, conn.ConsumeLogs(context.Background(), testLogs))

			allLogs := sink.AllLogs()
			if len(tc.expectedFile) == 0 {
				assert.Len(t, allLogs, 0)
				return
			}

			expected, err := golden.ReadLogs(filepath.Join("testdata", "logsToLogs", tc.expectedFile))

			assert.NoError(t, err)
			assert.Equal(t, allLogs[0].LogRecordCount(), expected.LogRecordCount())
			assert.NoError(t, plogtest.CompareLogs(expected, allLogs[0], plogtest.IgnoreObservedTimestamp()))
		})
	}
}

func TestMetricsToLogs(t *testing.T) {
	testCases := []struct {
		name         string
		inputFile    string
		expectedFile string
	}{
		{
			name:         "when entity is inferred, log event is sent",
			inputFile:    "entity/input-metric.yaml",
			expectedFile: "entity/expected-log.yaml",
		},
		{
			name:      "when entity is not inferred, no log is sent",
			inputFile: "entity/input-metric-nomatch.yaml",
		},
		{
			name:         "when relationship for different type is inferred log event is sent",
			inputFile:    "relationship/input-metric-different-type-relationship.yaml",
			expectedFile: "relationship/expected-metric-different-type-relationship.yaml",
		},
		{
			name:         "when relationship for same type is inferred log event is sent",
			inputFile:    "relationship/input-metric-same-type-relationship.yaml",
			expectedFile: "relationship/expected-metric-same-type-relationship.yaml",
		},
		{
			name:         "when relationship for same type is not inferred no log is sent",
			inputFile:    "relationship/input-metric-same-type-relationship-nomatch.yaml",
			expectedFile: "relationship/expected-metric-same-type-relationship-nomatch.yaml",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			factory := NewFactory()
			sink := &consumertest.LogsSink{}
			conn, err := factory.CreateMetricsToLogs(context.Background(),
				connectortest.NewNopSettings(metadata.Type), &Config{
					Schema: config.Schema{
						Entities: configuredEntities,
						Events:   configuredEvents,
					},
					SourcePrefix:      "src.",
					DestinationPrefix: "dst.",
				}, sink)
			require.NoError(t, err)
			require.NotNil(t, conn)

			require.NoError(t, conn.Start(context.Background(), componenttest.NewNopHost()))
			defer func() {
				assert.NoError(t, conn.Shutdown(context.Background()))
			}()

			testMetrics, err := golden.ReadMetrics(filepath.Join("testdata", "metricsToLogs", tc.inputFile))
			assert.NoError(t, err)
			assert.NoError(t, conn.ConsumeMetrics(context.Background(), testMetrics))

			allLogs := sink.AllLogs()
			if len(tc.expectedFile) == 0 {
				assert.Len(t, allLogs, 0)
				return
			}

			expected, err := golden.ReadLogs(filepath.Join("testdata", "metricsToLogs", tc.expectedFile))

			assert.NoError(t, err)
			assert.Equal(t, allLogs[0].LogRecordCount(), expected.LogRecordCount())
			assert.NoError(t, plogtest.CompareLogs(expected, allLogs[0], plogtest.IgnoreObservedTimestamp()))
		})
	}
}
