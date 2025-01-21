package timezone

import (
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/cli"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers"
)

const (
	timeZoneCommand = "timedatectl | grep 'Time zone'"
	timeZone        = "Time zone"
)

type provider struct {
	cli cli.CommandLineExecutor
}

var _ providers.Provider[TimeZone] = (*provider)(nil)

func CreateTimeZoneProvider() providers.Provider[TimeZone] {
	return &provider{
		cli: &cli.BashCli{},
	}
}

// Provide implements Provider.
func (dp *provider) Provide() <-chan TimeZone {
	ch := make(chan TimeZone)
	go func() {
		defer close(ch)
		stdout, err := cli.ProcessCommand(dp.cli, timeZoneCommand)
		// Bias field has to be set to invalid value because it
		// defaults to zero, which is valid in this context
		result := TimeZone{
			Bias: -1,
		}
		if err == nil {
			// The output contains result in format Time zone: xxx
			parsedOutput := providers.ParseKeyValue(stdout, ": ", []string{timeZone})
			result.Caption = parsedOutput[timeZone]
		}
		ch <- result
	}()
	return ch
}
