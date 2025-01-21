package firewall

import (
	"errors"
	"fmt"
	"sync"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/registry"
	"go.uber.org/zap"
)

const (
	firewallProfileKey = `SYSTEM\CurrentControlSet\services\SharedAccess\Parameters\FirewallPolicy`
	enabledValueName   = "EnableFirewall"
)

type provider struct {
	getRegistryValue registry.GetKeyUIntValueTypeFunc
}

var _ (providers.Provider[Container]) = (*provider)(nil)

func CreateFirewallProvider() providers.Provider[Container] {
	return &provider{
		getRegistryValue: registry.GetKeyUIntValue,
	}
}

// Provide implements providers.Provider.
func (fp *provider) Provide() <-chan Container {
	ch := make(chan Container)
	go func() {
		defer close(ch)
		result, err := fp.getFirewallProfiles()
		if err != nil {
			zap.L().Error("providing firewall profiles failed", zap.Error(err))
		}
		fc := Container{
			FirewallProfiles: result,
			Error:            err,
		}
		ch <- fc
	}()
	return ch
}

func getFirewallProfilesMapping() map[string]string {
	return map[string]string{
		"StandardProfile": "Private",
		"DomainProfile":   "Domain",
		"PublicProfile":   "Public",
	}
}

// getFirewallProfiles fetches individual profiles states from registry.
// Mapping of final profile names is based upon Get-NetFirewallProfile PS
// cmdlet outputs.
func (fp *provider) getFirewallProfiles() ([]Profile, error) {
	profileKeys := getFirewallProfilesMapping()
	errCh := make(chan error, len(profileKeys))
	profilesCh := make(chan Profile, len(profileKeys))

	wg := sync.WaitGroup{}
	for profileKey, profileType := range profileKeys {
		pk := profileKey
		pt := profileType
		wg.Add(1)
		go func() {
			defer wg.Done()
			value, err := fp.getRegistryValue(registry.LocalMachineKey, firewallProfileKey, pk, enabledValueName)

			if err == nil {
				profilesCh <- Profile{Name: pt, Enabled: int(value)}
			} else {
				errCh <- fmt.Errorf("failed to get profile '%s' value: %w", pk, err)
			}
		}()
	}
	wg.Wait()
	close(errCh)
	close(profilesCh)

	profiles, err := processFirewallProfilesChannels(errCh, profilesCh)
	return profiles, err
}

func processFirewallProfilesChannels(errCh chan error, profilesCh chan Profile) ([]Profile, error) {
	var err error
	for e := range errCh {
		err = errors.Join(e, err)
	}

	var profiles []Profile
	for p := range profilesCh {
		profiles = append(profiles, p)
	}

	return profiles, err
}
