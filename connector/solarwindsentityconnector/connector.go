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
	"github.com/solarwinds/solarwinds-otel-collector-releases/connector/solarwindsentityconnector/internal"
	"go.opentelemetry.io/collector/connector"

	"context"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

type solarwindsentity struct {
	logsConsumer consumer.Logs
	entities     *internal.Entities

	component.StartFunc
	component.ShutdownFunc
}

var _ connector.Metrics = (*solarwindsentity)(nil)
var _ connector.Logs = (*solarwindsentity)(nil)

func (s *solarwindsentity) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}

func (s *solarwindsentity) ConsumeMetrics(ctx context.Context, metrics pmetric.Metrics) error {
	logs := plog.NewLogs()
	events := internal.NewEvents(logs)
	for i := range metrics.ResourceMetrics().Len() {
		resourceMetric := metrics.ResourceMetrics().At(i)
		resourceAttrs := resourceMetric.Resource().Attributes()

		// This will be replaced with actual logic when conditions are introduced
		events.AppendEntityUpdateEvent(s.entities.GetEntity("Snowflake"), resourceAttrs)
	}

	if logs.LogRecordCount() == 0 {
		return nil
	}

	err := s.logsConsumer.ConsumeLogs(ctx, logs)
	if err != nil {
		return err
	}
	return nil
}

func (s *solarwindsentity) ConsumeLogs(ctx context.Context, logs plog.Logs) error {
	newLogs := plog.NewLogs()
	events := internal.NewEvents(newLogs)

	for i := range logs.ResourceLogs().Len() {
		resourceLog := logs.ResourceLogs().At(i)
		resourceAttrs := resourceLog.Resource().Attributes()

		// This will be replaced with actual logic when conditions are introduced
		events.AppendEntityUpdateEvent(s.entities.GetEntity("Snowflake"), resourceAttrs)
	}

	if newLogs.LogRecordCount() == 0 {
		return nil
	}

	err := s.logsConsumer.ConsumeLogs(ctx, newLogs)
	if err != nil {
		return err
	}
	return nil
}
