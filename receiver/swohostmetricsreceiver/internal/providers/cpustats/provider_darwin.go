package cpustats

import (
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers"
	"go.uber.org/zap"
)

type provider struct{}

func CreateProvider() providers.Provider[Container] {
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
