// Copyright 2024 SolarWinds Worldwide, LLC. All rights reserved.
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

package solarwindsexporter

import (
	"context"
	"errors"
	"fmt"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/otlpexporter"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type exporterType int

const (
	metricsExporterType exporterType = iota
	logsExporterType
	tracesExporterType
)

var extensionType = component.MustNewType("solarwinds")

type solarwindsExporter struct {
	exporterType
	config   *Config
	settings exporter.Settings
	metrics  exporter.Metrics
	logs     exporter.Logs
	traces   exporter.Traces
}

func newExporter(
	cfg *Config,
	settings exporter.Settings,
	typ exporterType,
) (*solarwindsExporter, error) {

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("validation of configuration failed: %w", err)
	}

	swiExporter := &solarwindsExporter{
		exporterType: typ,
		config:       cfg,
		settings:     settings,
	}

	return swiExporter, nil
}

func (swiExporter *solarwindsExporter) initExporterType(
	ctx context.Context,
	settings exporter.Settings,
	host component.Host,
	typ exporterType,
) error {
	swiExporter.exporterType = typ
	var extensionID *component.ID
	if swiExporter.config.ExtensionName != "" {
		id := component.NewIDWithName(extensionType, swiExporter.config.ExtensionName)
		extensionID = &id
	}

	swiExtension := findExtension(host.GetExtensions(), extensionID)
	if swiExtension == nil {
		return errors.New("solarwinds extension not found")
	}

	// Get token from the extensions.
	token := swiExtension.Token()
	swiExporter.config.ingestionToken = token

	// Get URl from the extension.
	url, err := swiExtension.Url()
	if err != nil {
		return fmt.Errorf(": %w", err)
	}
	swiExporter.config.endpointURL = url

	otlpExporter := otlpexporter.NewFactory()
	otlpCfg, err := swiExporter.config.OTLPConfig()
	if err != nil {
		return err
	}

	switch typ {
	case metricsExporterType:
		swiExporter.metrics, err = otlpExporter.CreateMetrics(ctx, settings, otlpCfg)
		return err
	case logsExporterType:
		swiExporter.logs, err = otlpExporter.CreateLogs(ctx, settings, otlpCfg)
		return err
	case tracesExporterType:
		swiExporter.traces, err = otlpExporter.CreateTraces(ctx, settings, otlpCfg)
		return err
	default:
		return fmt.Errorf("unknown exporter type: %v", typ)
	}

}

type EndpointConfigProvider interface {
	Url() (string, error)
	Token() configopaque.String
}

func findExtension(extensions map[component.ID]component.Component, cfgExtensionID *component.ID) EndpointConfigProvider {
	foundExtensions := make([]EndpointConfigProvider, 0)

	for foundExtensionID, ext := range extensions {
		if swiExtension, ok := ext.(EndpointConfigProvider); ok {
			// If configured extension ID is found, return it.
			if cfgExtensionID != nil && *cfgExtensionID == foundExtensionID {
				return swiExtension
			}

			// Otherwise, store it to the slice.
			foundExtensions = append(foundExtensions, swiExtension)
		}
	}

	// If no extension name configured and there is only one
	// found matching the type, return it.
	if len(foundExtensions) == 1 && cfgExtensionID != nil {
		return foundExtensions[0]
	}

	return nil
}

func (swiExporter *solarwindsExporter) start(ctx context.Context, host component.Host) error {
	if err := swiExporter.initExporterType(ctx, swiExporter.settings, host, swiExporter.exporterType); err != nil {
		return fmt.Errorf("failed to initialiaze exporter: %w", err)
	}

	switch swiExporter.exporterType {
	case metricsExporterType:
		return swiExporter.metrics.Start(ctx, host)
	case logsExporterType:
		return swiExporter.logs.Start(ctx, host)
	case tracesExporterType:
		return swiExporter.traces.Start(ctx, host)
	default:
		return fmt.Errorf("unknown exporter type: %v", swiExporter.exporterType)
	}
}

func (swiExporter *solarwindsExporter) shutdown(ctx context.Context) error {
	switch swiExporter.exporterType {
	case metricsExporterType:
		return swiExporter.metrics.Shutdown(ctx)
	case logsExporterType:
		return swiExporter.logs.Shutdown(ctx)
	case tracesExporterType:
		return swiExporter.traces.Shutdown(ctx)
	default:
		return fmt.Errorf("unknown exporter type: %v", swiExporter.exporterType)
	}
}

func (swiExporter *solarwindsExporter) pushMetrics(ctx context.Context, metrics pmetric.Metrics) error {
	if metrics.MetricCount() == 0 {
		return nil
	}

	return swiExporter.metrics.ConsumeMetrics(ctx, metrics)
}

func (swiExporter *solarwindsExporter) pushLogs(ctx context.Context, logs plog.Logs) error {
	if logs.LogRecordCount() == 0 {
		return nil
	}

	return swiExporter.logs.ConsumeLogs(ctx, logs)
}

func (swiExporter *solarwindsExporter) pushTraces(ctx context.Context, traces ptrace.Traces) error {
	if traces.SpanCount() == 0 {
		return nil
	}

	return swiExporter.traces.ConsumeTraces(ctx, traces)
}
