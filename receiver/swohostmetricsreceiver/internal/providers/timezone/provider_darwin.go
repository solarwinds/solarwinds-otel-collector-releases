package timezone

import "github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers"

type provider struct{}

var _ providers.Provider[TimeZone] = (*provider)(nil)

// Provide implements Provider.
func (dp *provider) Provide() <-chan TimeZone {
	ch := make(chan TimeZone)
	go dp.provideInternal(ch)
	return ch
}

func CreateTimeZoneProvider() providers.Provider[TimeZone] {
	return &provider{}
}

func (dp *provider) provideInternal(ch chan TimeZone) {
	defer close(ch)
}
