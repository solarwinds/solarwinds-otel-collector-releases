package firewall

import (
	"fmt"
	"testing"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers/firewall"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

type firewallProviderMock struct {
	fc firewall.Container
}

var _ providers.Provider[firewall.Container] = (*firewallProviderMock)(nil)

func CreateFirewallProviderMock(
	fc firewall.Container,
) providers.Provider[firewall.Container] {
	return &firewallProviderMock{
		fc: fc,
	}
}

// Provide implements providers.Provider.
func (m *firewallProviderMock) Provide() <-chan firewall.Container {
	ch := make(chan firewall.Container)
	go func() {
		defer close(ch)
		ch <- m.fc
	}()
	return ch
}

func produceFirewallProfileFromDataPoint(
	t *testing.T,
	dp pmetric.NumberDataPoint,
) firewall.Profile {
	v, b := dp.Attributes().Get(akFirewallProfile)
	assert.True(t, b, fmt.Sprintf("attribute %s must be available", akFirewallProfile))
	return firewall.Profile{Name: v.AsString(), Enabled: int(dp.IntValue())}
}
