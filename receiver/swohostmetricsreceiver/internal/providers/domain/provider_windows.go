package domain

import (
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/wmi"
)

type provider struct {
	wmi wmi.Executor
}

var _ providers.Provider[Domain] = (*provider)(nil)

// Win32_ComputerSystem represents actual Computer System WMI Object
// with subset of fields required for scraping.
type Win32_ComputerSystem struct {
	Name        string
	Domain      string
	DNSHostName string
	DomainRole  uint16
	Workgroup   string
}

// Provide implements DomainProvider.
func (dp *provider) Provide() <-chan Domain {
	ch := make(chan Domain)
	go func() {
		defer close(ch)
		result, err := wmi.QuerySingleResult[Win32_ComputerSystem](dp.wmi)
		if err == nil {
			ch <- Domain{
				Domain:     result.Domain,
				FQDN:       result.DNSHostName + "." + result.Domain,
				DomainRole: int(result.DomainRole),
				Workgroup:  result.Workgroup,
			}
		}
	}()
	return ch
}

func CreateDomainProvider() providers.Provider[Domain] {
	return &provider{
		wmi: wmi.NewExecutor(),
	}
}
