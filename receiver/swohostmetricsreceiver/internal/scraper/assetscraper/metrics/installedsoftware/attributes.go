package installedsoftware

import "github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers/installedsoftware"

// GetAttributes implements types.AttributesProvider.
func getAttributes(is installedsoftware.InstalledSoftware) map[string]any {
	const (
		name                    = "installedsoftware.name"
		date                    = "installedsoftware.date"
		publisher               = "installedsoftware.publisher"
		version                 = "installedsoftware.version"
		maximumSizeOfAttributes = 4
	)

	attributes := make(map[string]any, maximumSizeOfAttributes)
	if len(is.Name) > 0 {
		attributes[name] = is.Name
	}
	if len(is.InstallDate) > 0 {
		attributes[date] = is.InstallDate
	}
	if len(is.Publisher) > 0 {
		attributes[publisher] = is.Publisher
	}
	if len(is.Version) > 0 {
		attributes[version] = is.Version
	}
	return attributes
}
