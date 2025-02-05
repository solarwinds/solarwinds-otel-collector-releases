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
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/network"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func Test_SwohostmetricsreceiverRuns(t *testing.T) {
	expectedMetrics := []string{
		"swo.hostinfo.uptime",
		"os.cpu.numcores",
		"swo.hardwareinventory.cpu",
	}

	ctx := context.Background()

	net, err := network.New(ctx)
	require.NoError(t, err)
	testcontainers.CleanupNetwork(t, net)

	configName := "swohostmetricsreceiver.yaml"
	rContainer, err := runConnectedSolarWindsOTELCollectors(t, ctx, net.Name, configName)
	require.NoError(t, err)

	<-time.After(15 * time.Second)

	evaluateSWOHostMetrics(t, ctx, rContainer, expectedMetrics)
}

func Test_SwohostmetricsreceiverDefaultConfig(t *testing.T) {
	expectedDefaultMetrics := []string{
		"swo.hostinfo.uptime",
		"swo.hardwareinventory.cpu",
	}

	ctx := context.Background()

	net, err := network.New(ctx)
	require.NoError(t, err)
	testcontainers.CleanupNetwork(t, net)

	configName := "swohostmetricsreceiver_default.yaml"
	rContainer, err := runConnectedSolarWindsOTELCollectors(t, ctx, net.Name, configName)
	require.NoError(t, err)

	<-time.After(9 * time.Second)

	evaluateSWOHostMetrics(t, ctx, rContainer, expectedDefaultMetrics)
}

func evaluateSWOHostMetrics(
	t *testing.T,
	ctx context.Context,
	rContainer testcontainers.Container,
	expectedMetrics []string,
) {
	heartbeatMetricName := "sw.otelcol.uptime"

	lines, err := loadResultFile(ctx, rContainer, "/tmp/result.json")
	require.NoError(t, err)

	expectedMetricsMap := initiateFocusedMetricMap(expectedMetrics)

	jum := new(pmetric.JSONUnmarshaler)
	for _, line := range lines {
		m, err := jum.UnmarshalMetrics([]byte(line))
		if err != nil {
			continue
		}

		rmCount := m.ResourceMetrics().Len()
		for rmi := 0; rmi < rmCount; rmi++ {

			smCount := m.ResourceMetrics().At(rmi).ScopeMetrics().Len()
			for smi := 0; smi < smCount; smi++ {

				mCount := m.ResourceMetrics().At(rmi).ScopeMetrics().At(smi).Metrics().Len()
				for mi := 0; mi < mCount; mi++ {
					mName := m.ResourceMetrics().At(rmi).ScopeMetrics().At(smi).Metrics().At(mi).Name()

					// Heart beat metric is filtered out.
					if mName == heartbeatMetricName {
						continue
					}

					// Only occurrence on focussed metrics is recorded.
					if _, found := expectedMetricsMap[mName]; found {
						expectedMetricsMap[mName]++
					}
				}
			}
		}
	}

	var somethingMissed bool
	for k, v := range expectedMetricsMap {
		if v == 0 {
			somethingMissed = true
			log.Printf("*** evaluation: metric '%s' hasn't arrived\n", k)
		}
	}
	require.False(t, somethingMissed, "all required metrics must arrive")
}

func initiateFocusedMetricMap(
	expectedMetrics []string,
) map[string]int {
	mm := make(map[string]int, len(expectedMetrics))
	for _, m := range expectedMetrics {
		mm[m] = 0
	}
	return mm
}
