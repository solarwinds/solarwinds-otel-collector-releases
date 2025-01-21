package language

import (
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/cli"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers"
)

const (
	localeCommand = "localectl"
	language      = "LANGUAGE"
)

type provider struct {
	cli cli.CommandLineExecutor
}

var _ providers.Provider[Language] = (*provider)(nil)

// Provide implements Provider.
func (lp *provider) Provide() <-chan Language {
	ch := make(chan Language)
	go func() {
		defer close(ch)
		stdout, err := cli.ProcessCommand(lp.cli, localeCommand)
		if err == nil {
			// localeCommand output is in format LANGUAGE=en_US
			parsedOutput := providers.ParseKeyValue(stdout, "=", []string{language})
			ch <- Language{
				Name: parsedOutput[language],
			}
		}
	}()
	return ch
}

func CreateLanguageProvider() providers.Provider[Language] {
	return &provider{
		cli: &cli.BashCli{},
	}
}
