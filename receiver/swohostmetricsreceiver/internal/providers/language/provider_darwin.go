package language

import "github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers"

type provider struct{}

var _ providers.Provider[Language] = (*provider)(nil)

// Provide implements Provider.
func (dp *provider) Provide() <-chan Language {
	ch := make(chan Language)
	go dp.provideInternal(ch)
	return ch
}

func CreateLanguageProvider() providers.Provider[Language] {
	return &provider{}
}

func (dp *provider) provideInternal(ch chan Language) {
	defer close(ch)
}
