// Copyright 2025 SolarWinds Worldwide, LLC. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package domain

import (
	"github.com/solarwinds/solarwinds-otel-collector-releases/receiver/swohostmetricsreceiver/internal/providers"
	"github.com/solarwinds/solarwinds-otel-collector-releases/receiver/swohostmetricsreceiver/internal/wmi"
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
