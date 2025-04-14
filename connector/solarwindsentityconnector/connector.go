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
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

type solarwindsentity struct {
	logsConsumer consumer.Logs

	component.StartFunc
	component.ShutdownFunc
}

func (s *solarwindsentity) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}

func (s *solarwindsentity) ConsumeMetrics(ctx context.Context, metrics pmetric.Metrics) error {
	logs := plog.NewLogs()
	err := s.logsConsumer.ConsumeLogs(ctx, logs)
	if err != nil {
		return err
	}
	return nil
}

func (s *solarwindsentity) ConsumeLogs(ctx context.Context, logs plog.Logs) error {
	newLogs := plog.NewLogs()
	err := s.logsConsumer.ConsumeLogs(ctx, newLogs)
	if err != nil {
		return err
	}
	return nil
}
