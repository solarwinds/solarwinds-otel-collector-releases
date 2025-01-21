package timezone

import (
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/wmi"
)

type provider struct {
	wmi wmi.Executor
}

var _ providers.Provider[TimeZone] = (*provider)(nil)

func CreateTimeZoneProvider() providers.Provider[TimeZone] {
	return &provider{
		wmi: wmi.NewExecutor(),
	}
}

// Win32_TimeZone represents actual Time Zone WMI Object
// with subset of fields required for scraping.
type Win32_TimeZone struct {
	Bias         int32
	Caption      string
	StandardName string
}

// Provide implements Provider.
func (tp *provider) Provide() <-chan TimeZone {
	ch := make(chan TimeZone)
	go func() {
		defer close(ch)
		result, err := wmi.QuerySingleResult[Win32_TimeZone](tp.wmi)
		if err == nil {
			ch <- TimeZone{Bias: int(result.Bias), StandardName: result.StandardName, Caption: result.Caption}
		}
	}()
	return ch
}
