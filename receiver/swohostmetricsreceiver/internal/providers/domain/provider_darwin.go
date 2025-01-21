package domain

import "github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers"

type provider struct{}

var _ providers.Provider[Domain] = (*provider)(nil)

// Provide implements Provider.
func (dp *provider) Provide() <-chan Domain {
	ch := make(chan Domain)
	go dp.provideInternal(ch)
	return ch
}

func CreateDomainProvider() providers.Provider[Domain] {
	return &provider{}
}

func (dp *provider) provideInternal(ch chan Domain) {
	defer close(ch)
}
