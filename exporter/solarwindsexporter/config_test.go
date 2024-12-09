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
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/confmap/confmaptest"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

// loadConfigTestdata is a helper function to load a testdata
// file by its name.
func loadConfigTestdata(t *testing.T, name string) *confmap.Conf {
	t.Helper()

	filename := fmt.Sprintf("%s.yaml", name)
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", filename))
	require.NoError(t, err)

	return cm
}

// TestConfigUnmarshalFull tries to parse a configuration file
// with all values provided and verifies the configuration.
func TestConfigUnmarshalFull(t *testing.T) {
	cfgFile := loadConfigTestdata(t, "full")

	// Parse configuration.
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()
	require.NoError(t, cfgFile.Unmarshal(&cfg))

	// Verify the values.
	assert.Equal(t, &Config{
		ExtensionName: "swo",
		BackoffSettings: configretry.BackOffConfig{
			Enabled:             false,
			InitialInterval:     15000000000,
			RandomizationFactor: 0.7,
			Multiplier:          2.4,
			MaxInterval:         40000000000,
			MaxElapsedTime:      400000000000,
		},
		QueueSettings: exporterhelper.QueueConfig{
			Enabled:      true,
			NumConsumers: 20,
			QueueSize:    2000,
		},
		Timeout: exporterhelper.TimeoutConfig{
			Timeout: 20000000000,
		},
	}, cfg)
}

// TestConfigValidateOK verifies that a configuration
// file containing only the mandatory values successfully
// validates.
func TestConfigValidateOK(t *testing.T) {
	cfgFile := loadConfigTestdata(t, "minimal")

	// Parse configuration.
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()
	require.NoError(t, cfgFile.Unmarshal(&cfg))

	// Try to validate it.
	assert.NoError(t, cfg.(*Config).Validate())
}

// TestConfigTokenRedacted checks that the configuration
// type doesn't leak its secret token unless it is accessed explicitly.
func TestConfigTokenRedacted(t *testing.T) {
	cfg := &Config{
		ingestionToken: "SECRET",
	}
	// This is the only way of accessing the actual token.
	require.Equal(t, "SECRET", string(cfg.ingestionToken))

	// It is redacted when printed.
	assert.Equal(t, "[REDACTED]", cfg.ingestionToken.String())
}
