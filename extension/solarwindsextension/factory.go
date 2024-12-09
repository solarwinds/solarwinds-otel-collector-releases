package solarwindsextension

import (
	"context"
	"fmt"
	"github.com/solarwinds/solarwinds-otel-collector/extension/solarwindsextension/internal"
	"github.com/solarwinds/solarwinds-otel-collector/extension/solarwindsextension/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

func NewFactory() extension.Factory {
	return extension.NewFactory(
		metadata.Type,
		internal.NewDefaultConfig,
		createExtension,
		metadata.ExtensionStability)
}

func createExtension(ctx context.Context, set extension.Settings, cfg component.Config) (extension.Extension, error) {
	extCfg, ok := cfg.(*internal.Config)
	if !ok {
		return nil, fmt.Errorf("unexpected config type: %T", cfg)
	}
	return newExtension(ctx, set, extCfg)
}
