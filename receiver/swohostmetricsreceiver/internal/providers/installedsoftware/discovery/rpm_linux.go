package discovery

import (
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/cli"
)

type rpmDiscovery struct {
	bash cli.CommandLineExecutor
}

var _ (Discovery) = (*rpmDiscovery)(nil)

func NewRpmDiscovery() Discovery {
	return createRpmDiscovery(
		cli.NewBashCliExecutor(),
	)
}

func createRpmDiscovery(
	bash cli.CommandLineExecutor,
) Discovery {
	return &rpmDiscovery{
		bash: bash,
	}
}

// Discover implements InstalledSoftwareDiscovery.
func (d *rpmDiscovery) Discover() bool {
	command := "rpm --version"
	_, _, err := d.bash.ExecuteCommand(command)
	return err == nil
}
