package e2e

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/network"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestSingleSolarWindsOtelCollector(t *testing.T) {
	ctx, _ := context.WithCancel(context.Background())

	net, err := network.New(ctx)
	require.NoError(t, err)
	testcontainers.CleanupNetwork(t, net)

	rContainer, err := runReceivingSolarWindsOTELCollector(ctx, net.Name)
	require.NoError(t, err)
	testcontainers.CleanupContainer(t, rContainer)

	eContainer, err := runEmittingSolarWindsOTELCollector(ctx, net.Name)
	require.NoError(t, err)
	testcontainers.CleanupContainer(t, eContainer)

	<-time.After(20 * time.Second)
	log.Println("***: termination in progress")

	err = eContainer.Terminate(ctx)
	require.NoError(t, err)

	err = rContainer.Terminate(ctx)
	require.NoError(t, err)
}

func runReceivingSolarWindsOTELCollector(
	ctx context.Context,
	networkName string,
) (testcontainers.Container, error) {
	configPath, err := filepath.Abs(filepath.Join(".", "testdata", "receiving_collector.yaml"))
	if err != nil {
		return nil, err
	}

	r, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}

	lc := new(MyLogConsumer)
	lc.Prefix = "receiving"
	req := testcontainers.ContainerRequest{
		Image: "solarwinds-otel-collector:latest",
		LogConsumerCfg: &testcontainers.LogConsumerConfig{
			/*Opts: []testcontainers.LogProductionOption{
				testcontainers.WithLogProductionTimeout(10 * time.Second),
			},*/
			Consumers: []testcontainers.LogConsumer{lc},
		},
		Files: []testcontainers.ContainerFile{
			{
				HostFilePath:      configPath,
				Reader:            r,
				ContainerFilePath: "/opt/default-config.yaml",
				FileMode:          0o440,
			},
		},
		WaitingFor: wait.ForLog("Everything is ready. Begin running and processing data."),
		Networks:   []string{networkName},
		Name:       "receiving",
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	return container, err
}

func runEmittingSolarWindsOTELCollector(
	ctx context.Context,
	networkName string,
) (testcontainers.Container, error) {
	configPath, err := filepath.Abs(filepath.Join(".", "testdata", "emitting_collector.yaml"))
	if err != nil {
		return nil, err
	}

	r, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}

	lc := new(MyLogConsumer)
	lc.Prefix = "emitting"
	req := testcontainers.ContainerRequest{
		Image: "solarwinds-otel-collector:latest",
		LogConsumerCfg: &testcontainers.LogConsumerConfig{
			/*Opts: []testcontainers.LogProductionOption{
				testcontainers.WithLogProductionTimeout(10 * time.Second),
			},*/
			Consumers: []testcontainers.LogConsumer{lc},
		},
		Files: []testcontainers.ContainerFile{
			{
				HostFilePath:      configPath,
				Reader:            r,
				ContainerFilePath: "/opt/default-config.yaml",
				FileMode:          0o440,
			},
		},
		WaitingFor: wait.ForLog("Everything is ready. Begin running and processing data."),
		Networks:   []string{networkName},
		Name:       "emitting",
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	return container, err
}

type MyLogConsumer struct {
	Prefix string
}

func (lc *MyLogConsumer) Accept(l testcontainers.Log) {
	msg := fmt.Sprintf("***%s: %s", lc.Prefix, string(l.Content))
	//fmt.Print(msg)
	log.Print(msg)
}
