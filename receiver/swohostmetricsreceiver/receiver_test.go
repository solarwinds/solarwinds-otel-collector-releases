//go:build !integration

package swohostmetricsreceiver

import (
	"context"
	"testing"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
)

const (
	expectedReceiverType = "swohostmetrics"
	expectedStability    = component.StabilityLevelDevelopment
)

func Test_NewFactory_ReceiverFactoryHasProperNameAndStabilityLevel(t *testing.T) {
	sut := NewFactory()

	if sut.Type().String() != expectedReceiverType {
		t.Fatalf("Receiver has incorrect name [%s]", sut.Type())
	}

	if sut.MetricsStability() != expectedStability {
		t.Fatalf(
			"Receiver has incorrect stability level for metrics receiver [%s]",
			sut.MetricsStability(),
		)
	}
}

func Test_NewFactory_ReceiverSupportMetrics(t *testing.T) {
	sut := NewFactory()
	metricReceiver, err := sut.CreateMetrics(
		context.TODO(),
		receivertest.NewNopSettings(), // for testing purposes provided receiver settings object
		&ReceiverConfig{
			ControllerConfig: scraperhelper.ControllerConfig{
				CollectionInterval: 10,
			},
			Scrapers: make(map[string]component.Config, 0),
		},
		new(consumertest.MetricsSink), // for testing purposes provided consumer object
	)
	if err != nil {
		t.Fatal("Metric receiver creation failed")
	}

	if metricReceiver == nil {
		t.Fatal("Metric receiver creation failed")
	}
}
