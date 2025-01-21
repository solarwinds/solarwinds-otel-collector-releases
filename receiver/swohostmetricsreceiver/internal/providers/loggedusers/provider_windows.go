package loggedusers

import (
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/registry"
)

type provider struct {
	getRegistryValues registry.GetKeyValuesTypeFunc
}

var _ providers.Provider[Data] = (*provider)(nil)

func CreateProvider() providers.Provider[Data] {
	return &provider{getRegistryValues: registry.GetKeyValues}
}

func (dp *provider) Provide() <-chan Data {
	ch := make(chan Data)
	go dp.provideInternal(ch)
	return ch
}

func (dp *provider) provideInternal(ch chan Data) {
	defer close(ch)
	values, err := dp.getRegistryValues(
		registry.LocalMachineKey,
		`SOFTWARE\Microsoft\Windows\CurrentVersion\Authentication`,
		`LogonUI`,
		[]string{`LastLoggedOnUser`, `LastLoggedOnDisplayName`},
	)
	if err != nil {
		ch <- Data{Error: err}
		return
	}

	ch <- Data{
		Users: []User{{
			Name:        values[`LastLoggedOnUser`],
			DisplayName: values[`LastLoggedOnDisplayName`],
		}},
	}
}
