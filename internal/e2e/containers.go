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
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/mdelapenya/tlscert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/network"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	receivingContainer            = "receiver"
	testedContainer               = "sut"
	generatingContainer           = "generator"
	port                          = 17016
	collectorRunningPeriod        = 35 * time.Second
	samplesCount                  = 10
	receivingContainerResultsPath = "/tmp/result.json"
)

func runReceivingSolarWindsOTELCollector(
	ctx context.Context,
	certDir string,
	networkName string,
) (testcontainers.Container, error) {
	configPath, err := filepath.Abs(filepath.Join(".", "testdata", "receiving_collector.yaml"))
	if err != nil {
		return nil, err
	}

	// Used by the OTLP/gRPC Receiver for TLS (see its config).
	additionalFiles := []testcontainers.ContainerFile{
		{
			HostFilePath:      filepath.Join(certDir, "cert-server.pem"),
			ContainerFilePath: "/opt/cert-server.pem",
			FileMode:          0o644,
		},
		{
			HostFilePath:      filepath.Join(certDir, "key-server.pem"),
			ContainerFilePath: "/opt/key-server.pem",
			FileMode:          0o644,
		},
	}

	return runSolarWindsOTELCollector(
		ctx,
		networkName,
		receivingContainer,
		configPath,
		additionalFiles,
	)
}

func runTestedSolarWindsOTELCollector(
	ctx context.Context,
	certDir string,
	networkName string,
	configName string,
) (testcontainers.Container, error) {
	configPath, err := filepath.Abs(filepath.Join(".", "testdata", configName))
	if err != nil {
		return nil, err
	}

	// Add the root certificate for the self-signed certs as trusted.
	// Warning: This actually replaces all root certificates in the container.
	additionalFiles := []testcontainers.ContainerFile{
		{
			HostFilePath:      filepath.Join(certDir, "cert-ca.pem"),
			ContainerFilePath: "/etc/ssl/certs/ca-certificates.crt",
			FileMode:          0o644,
		},
	}

	return runSolarWindsOTELCollector(
		ctx,
		networkName,
		testedContainer,
		configPath,
		additionalFiles,
	)
}

func runSolarWindsOTELCollector(
	ctx context.Context,
	networkName string,
	containerName string,
	configPath string,
	additionalFiles []testcontainers.ContainerFile,
) (testcontainers.Container, error) {
	lc := new(logConsumer)
	lc.Prefix = containerName

	files := []testcontainers.ContainerFile{
		{
			HostFilePath:      configPath,
			ContainerFilePath: "/opt/default-config.yaml",
			FileMode:          0o440,
		},
	}
	files = append(files, additionalFiles...)

	req := testcontainers.ContainerRequest{
		Image: "solarwinds-otel-collector:latest",
		LogConsumerCfg: &testcontainers.LogConsumerConfig{
			Consumers: []testcontainers.LogConsumer{lc},
		},
		Files:      files,
		WaitingFor: wait.ForLog("Everything is ready. Begin running and processing data."),
		Networks:   []string{networkName},
		Name:       containerName,
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	return container, err
}

type CertPaths struct {
	CaCertFile string
	CertFile   string
	KeyFile    string
}

// generateCertificates generates a new CA certificate and a server
// key and certificate derived from it for a given `host`.
// All files are stored in a `path`. All paths of files written are
// returned in a CertPaths struct.
func generateCertificates(host, path string) (*CertPaths, error) {
	caCert := tlscert.SelfSignedFromRequest(tlscert.Request{
		Name:      "ca",
		Host:      host,
		IsCA:      true,
		ParentDir: path,
	})
	if caCert == nil {
		return nil, errors.New("failed to generate ca certificate")
	}

	cert := tlscert.SelfSignedFromRequest(tlscert.Request{
		Name:      "server",
		Host:      host,
		IsCA:      true,
		Parent:    caCert,
		ParentDir: path,
	})
	if cert == nil {
		return nil, errors.New("failed to generate server certificate")
	}

	return &CertPaths{
		CaCertFile: caCert.CertPath,
		CertFile:   cert.CertPath,
		KeyFile:    cert.KeyPath,
	}, nil
}

func runGeneratorContainer(
	ctx context.Context,
	networkName string,
	cmd []string,
) (testcontainers.Container, error) {
	containerName := generatingContainer

	lc := new(logConsumer)
	lc.Prefix = containerName

	req := testcontainers.ContainerRequest{
		Image: "ghcr.io/open-telemetry/opentelemetry-collector-contrib/telemetrygen:latest",
		LogConsumerCfg: &testcontainers.LogConsumerConfig{
			Consumers: []testcontainers.LogConsumer{lc},
		},
		Networks: []string{networkName},
		Name:     containerName,
		Cmd:      cmd,
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	return container, err
}

// Starts the receiving collector, and the test collector.
// Test collector container is started with the supplied config file to be tested.
// Returns the receiving collector container instance, so tests can check the ingested data is as expected.
func startCollectorContainers(
	t *testing.T,
	ctx context.Context,
	config string,
	signalType SignalType,
	waitTime time.Duration,
) testcontainers.Container {

	net, err := network.New(ctx)
	require.NoError(t, err)
	testcontainers.CleanupNetwork(t, net)

	rContainer, err := runConnectedSolarWindsOTELCollectors(t, ctx, net.Name, config)
	require.NoError(t, err)

	cmd := []string{
		signalType.String(),
		fmt.Sprintf("--%s", signalType), strconv.Itoa(samplesCount),
		"--otlp-insecure",
		"--otlp-endpoint", fmt.Sprintf("%s:%d", testedContainer, port),
		"--otlp-attributes", fmt.Sprintf("%s=\"%s\"", resourceAttributeName, resourceAttributeValue),
	}

	gContainer, err := runGeneratorContainer(ctx, net.Name, cmd)
	require.NoError(t, err)
	testcontainers.CleanupContainer(t, gContainer)

	<-time.After(waitTime)

	return rContainer
}

func runConnectedSolarWindsOTELCollectors(
	t *testing.T,
	ctx context.Context,
	networkName string,
	emittingCollectorConfig string,
) (testcontainers.Container, error) {
	certPath := t.TempDir()
	_, err := generateCertificates(receivingContainer, certPath)
	if err != nil {
		return nil, err
	}

	rContainer, err := runReceivingSolarWindsOTELCollector(ctx, certPath, networkName)
	if err != nil {
		return nil, err
	}
	testcontainers.CleanupContainer(t, rContainer)

	eContainer, err := runTestedSolarWindsOTELCollector(ctx, certPath, networkName, emittingCollectorConfig)
	if err != nil {
		return nil, err
	}
	testcontainers.CleanupContainer(t, eContainer)

	return rContainer, nil
}

type logConsumer struct {
	Prefix string
}

func (lc *logConsumer) Accept(l testcontainers.Log) {
	log.Printf("***%s: %s", lc.Prefix, string(l.Content))
}

type SignalType int

const (
	Logs SignalType = iota
	Metrics
	Traces
)

func (s SignalType) String() string {
	switch s {
	case Logs:
		return "logs"
	case Metrics:
		return "metrics"
	case Traces:
		return "traces"
	default:
		panic("unexpected signal type")
	}
}
