package firewall

import (
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers"
	"go.uber.org/zap"
)

type provider struct{}

var _ (providers.Provider[Container]) = (*provider)(nil)

func CreateFirewallProvider() providers.Provider[Container] {
	return &provider{}
}

// Provide implements providers.Provider.
func (*provider) Provide() <-chan Container {
	ch := make(chan Container)
	go func() {
		defer close(ch)
		zap.L().Warn("This provider is not supported on Darwin")
	}()
	return ch
}
