package solarwindsextension

import (
	"context"
	"errors"
	"github.com/solarwinds/solarwinds-otel-collector/extension/solarwindsextension/internal"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.uber.org/zap"
)

type SolarwindsExtension struct {
	logger       *zap.Logger
	config       *internal.Config
	uptimeMetric *internal.UptimeMetric
	heartbeat    *internal.Heartbeat
	exporter     *internal.Exporter
}

func newExtension(ctx context.Context, set extension.Settings, cfg *internal.Config) (*SolarwindsExtension, error) {
	set.Logger.Info("Creating SolarwindsExtension")
	set.Logger.Info("Config", zap.Any("config", cfg))
	e := &SolarwindsExtension{
		logger: set.Logger,
		config: cfg,
	}
	var err error
	e.exporter, err = internal.NewExporter(ctx, set, cfg, e.decorateResourceAttributes)
	if err != nil {
		return nil, err
	}

	e.uptimeMetric = internal.NewUptimeMetric(set.Logger)
	e.heartbeat = internal.NewHeartbeat(e.logger, e.exporter.PushMetrics, e.uptimeMetric.AddUptimeMetric)

	return e, nil
}

func (e *SolarwindsExtension) GetEndpointConfig() EndpointConfig { return newEndpointConfig(e.config) }

func (e *SolarwindsExtension) Start(ctx context.Context, host component.Host) error {
	e.logger.Info("Starting SolarwindsExtension")
	err := e.exporter.Start(ctx, host)
	if err != nil {
		return err
	}
	return e.heartbeat.Start()
}

func (e *SolarwindsExtension) Shutdown(ctx context.Context) error {
	e.logger.Info("Shutting down SolarwindsExtension")
	// Everything must be shut down, regardless of the failure.
	return errors.Join(
		e.heartbeat.Shutdown(),
		e.exporter.Shutdown(ctx))
}

func (e *SolarwindsExtension) decorateResourceAttributes(resource pcommon.Resource) error {
	if e.config.CollectorName != "" {
		resource.Attributes().PutStr("collector_name", e.config.CollectorName)
	}
	return nil
}
