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

//go:build e2e

package e2e

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	containerName = "standalone"
)

func TestConfigIsNotProvided_FailsToRun(t *testing.T) {
	ctx := context.Background()

	lc := new(logConsumer)
	lc.Prefix = containerName

	req := testcontainers.ContainerRequest{
		Image: "solarwinds-otel-collector:latest",
		LogConsumerCfg: &testcontainers.LogConsumerConfig{
			Consumers: []testcontainers.LogConsumer{lc},
		},
		WaitingFor: wait.ForLog("open /opt/default-config.yaml: no such file or directory"),
		Name:       containerName,
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err)
	testcontainers.CleanupContainer(t, container)
}

func TestInvalidConfigProvided_FailsToRun(t *testing.T) {
	configPath, err := filepath.Abs(filepath.Join(".", "testdata", "invalid_setup.yaml"))
	require.NoError(t, err)

	files := []testcontainers.ContainerFile{
		{
			HostFilePath:      configPath,
			ContainerFilePath: "/opt/default-config.yaml",
			FileMode:          0o440,
		},
	}

	ctx := context.Background()

	lc := new(logConsumer)
	lc.Prefix = containerName

	req := testcontainers.ContainerRequest{
		Image: "solarwinds-otel-collector:latest",
		LogConsumerCfg: &testcontainers.LogConsumerConfig{
			Consumers: []testcontainers.LogConsumer{lc},
		},
		Files:      files,
		WaitingFor: wait.ForLog("failed to get config: cannot unmarshal the configuration: decoding failed"),
		Name:       containerName,
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err)
	testcontainers.CleanupContainer(t, container)
}
