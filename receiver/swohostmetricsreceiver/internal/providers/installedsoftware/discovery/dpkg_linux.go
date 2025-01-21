package discovery

import (
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/cli"
)

type dpkgDiscovery struct {
	bash cli.CommandLineExecutor
}

var _ (Discovery) = (*dpkgDiscovery)(nil)

func NewDpkgDiscovery() Discovery {
	return createDpkgDiscovery(
		cli.NewBashCliExecutor(),
	)
}

func createDpkgDiscovery(
	bash cli.CommandLineExecutor,
) Discovery {
	return &dpkgDiscovery{
		bash: bash,
	}
}

// Discover implements InstalledSoftwareDiscovery.
func (d *dpkgDiscovery) Discover() bool {
	command := "dpkg --version"
	_, _, err := d.bash.ExecuteCommand(command)
	return err == nil
}
