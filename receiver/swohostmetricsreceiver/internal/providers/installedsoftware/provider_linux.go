package installedsoftware

import (
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers/installedsoftware/discovery"
	"go.uber.org/zap"
)

type linuxProvider struct{}

var _ Provider = (*linuxProvider)(nil)

func NewInstalledSoftwareProvider() Provider {
	return createInstalledSoftwareProvider(
		map[string]linuxInstalledSoftwareContainer{
			"rpm": {
				Discovery: discovery.NewRpmDiscovery(),
				Provider:  NewRpmProvider(),
			},
			"dpkg": {
				Discovery: discovery.NewDpkgDiscovery(),
				Provider:  NewDpkgProvider(),
			},
		},
		getDefaultProvider(),
	)
}

func createInstalledSoftwareProvider(
	discoverableProviders map[string]linuxInstalledSoftwareContainer,
	fallbackProvider Provider,
) Provider {
	provider := discoverProvider(
		discoverableProviders,
		fallbackProvider,
	)
	return provider
}

func getDefaultProvider() Provider {
	return new(linuxProvider)
}

// GetSoftware implements Provider.
func (p *linuxProvider) GetSoftware() ([]InstalledSoftware, error) {
	zap.L().Debug("unable to provide installed software via linuxProvider")
	return make([]InstalledSoftware, 0), nil
}

type linuxInstalledSoftwareContainer struct {
	Discovery discovery.Discovery
	Provider  Provider
}

func discoverProvider(
	discoverableProviders map[string]linuxInstalledSoftwareContainer,
	fallbackProvider Provider,
) Provider {
	// go through providers and select the most prioritized one
	for _, container := range discoverableProviders {
		// discovery successful => use its provider
		if container.Discovery.Discover() {
			return container.Provider
		}
	}

	zap.L().Warn("default installed software provider for linux will be used")
	return fallbackProvider
}
