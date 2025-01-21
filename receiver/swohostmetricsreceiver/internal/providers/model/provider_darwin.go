package model

import "github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers"

type provider struct{}

var _ providers.Provider[Model] = (*provider)(nil)

func CreateModelProvider() providers.Provider[Model] {
	return &provider{}
}

// Provide implements Provider.
func (p *provider) Provide() <-chan Model {
	ch := make(chan Model)
	go func() {
		close(ch)
	}()
	return ch
}
