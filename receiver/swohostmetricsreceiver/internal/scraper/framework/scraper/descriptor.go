package scraper

import (
	"go.opentelemetry.io/collector/component"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/scraper/framework/scope"
)

// Descriptor represents description of scraper.
type Descriptor struct {
	// Type is name of described scraper.
	Type component.Type
	// Scopes belong to these scrapers. Map keys represent names of scopes.
	ScopeDescriptors map[string]scope.Descriptor
}
