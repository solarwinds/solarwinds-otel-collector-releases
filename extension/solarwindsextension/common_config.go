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

package solarwindsextension

import (
	"go.opentelemetry.io/collector/config/configopaque"

	"github.com/solarwinds/solarwinds-otel-collector-releases/extension/solarwindsextension/internal"
)

type CommonConfig interface {
	Url() (string, error)
	Token() configopaque.String
	CollectorName() string
	WithoutEntity() bool
}

type commonConfig struct{ cfg *internal.Config }

var _ CommonConfig = (*commonConfig)(nil)

func newCommonConfig(cfg *internal.Config) *commonConfig {
	return &commonConfig{cfg: cfg}
}

func (c *commonConfig) Url() (string, error) {
	return c.cfg.EndpointUrl()
}

func (c *commonConfig) Token() configopaque.String {
	return c.cfg.IngestionToken
}

func (c *commonConfig) CollectorName() string {
	return c.cfg.CollectorName
}

func (c *commonConfig) WithoutEntity() bool { return c.cfg.WithoutEntity }
