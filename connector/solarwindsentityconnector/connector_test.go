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
	"github.com/solarwinds/solarwinds-otel-collector-releases/connector/solarwindsentityconnector/internal"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
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

func createTestConnector(t *testing.T) *solarwindsentity {
	return &solarwindsentity{
		entities: internal.NewEntities(map[string]internal.Entity{
			"Snowflake": internal.NewEntity(
				"Snowflake",
				[]string{"id1", "id2"},
				[]string{"attr1", "attr2"}),
		}),
		logsConsumer: &testLogsConsumer{},
	}
}

func TestConsumeMetrics(t *testing.T) {
	connector := createTestConnector(t)

	// Create a test metrics object
	metrics := pmetric.NewMetrics()
	resourceMetrics := metrics.ResourceMetrics().AppendEmpty()
	resourceMetrics.Resource().Attributes().PutStr("id1", "value1")
	resourceMetrics.Resource().Attributes().PutStr("id2", "value2")

	// TODO: not sure how to do the check
	err := connector.ConsumeMetrics(context.Background(), metrics)

	// Assert that there were no errors
	if err != nil {
		t.Errorf("ConsumeMetrics() error = %v", err)
	}
}
