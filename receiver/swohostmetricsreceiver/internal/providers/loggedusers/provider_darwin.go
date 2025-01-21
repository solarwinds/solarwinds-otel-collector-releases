package loggedusers

import "github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers"

type provider struct{}

var _ providers.Provider[Data] = (*provider)(nil)

// Provide implements Provider.
func (dp *provider) Provide() <-chan Data {
	ch := make(chan Data)
	go dp.provideInternal(ch)
	return ch
}

func CreateProvider() providers.Provider[Data] {
	return &provider{}
}

func (dp *provider) provideInternal(ch chan Data) {
	defer close(ch)
}
