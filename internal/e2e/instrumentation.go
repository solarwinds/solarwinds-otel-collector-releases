package e2e

import (
	"context"
	"log"
	"path/filepath"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	receivingContainer     = "receiver"
	testedContainer        = "sut"
	generatingContainer    = "generator"
	port                   = 17016
	collectorRunningPeriod = 35 * time.Second
	samplesCount           = 10
)

func runReceivingSolarWindsOTELCollector(
	ctx context.Context,
	networkName string,
) (testcontainers.Container, error) {
	configPath, err := filepath.Abs(filepath.Join(".", "testdata", "receiving_collector.yaml"))
	if err != nil {
		return nil, err
	}

	container, err := runSolarWindsOTELCollector(ctx, networkName, receivingContainer, configPath)
	return container, err
}

func runTestedSolarWindsOTELCollector(
	ctx context.Context,
	networkName string,
) (testcontainers.Container, error) {
	configPath, err := filepath.Abs(filepath.Join(".", "testdata", "emitting_collector.yaml"))
	if err != nil {
		return nil, err
	}

	container, err := runSolarWindsOTELCollector(ctx, networkName, testedContainer, configPath)
	return container, err
}

func runSolarWindsOTELCollector(
	ctx context.Context,
	networkName string,
	containerName string,
	configPath string,
) (testcontainers.Container, error) {
	lc := new(logConsumer)
	lc.Prefix = containerName
	req := testcontainers.ContainerRequest{
		Image: "solarwinds-otel-collector:latest",
		LogConsumerCfg: &testcontainers.LogConsumerConfig{
			Consumers: []testcontainers.LogConsumer{lc},
		},
		Files: []testcontainers.ContainerFile{
			{
				HostFilePath:      configPath,
				ContainerFilePath: "/opt/default-config.yaml",
				FileMode:          0o440,
			},
		},
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

type logConsumer struct {
	Prefix string
}

func (lc *logConsumer) Accept(l testcontainers.Log) {
	log.Printf("***%s: %s", lc.Prefix, string(l.Content))
}
