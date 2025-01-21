package hostdetails

import (
	"strconv"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/attributes/shared"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers/domain"
)

// Following keys will be used in metric as attributes.
const (
	domainObject    = "hostdetails.domain"
	domainFqdn      = "hostdetails.domain.fqdn"
	domainRole      = "hostdetails.domain.role"
	domainWorkgroup = "hostdetails.domain.workgroup"
	domainCount     = 5
)

type domainAttributesGenerator struct {
	DomainProvider providers.Provider[domain.Domain]
}

var _ shared.AttributesGenerator = (*domainAttributesGenerator)(nil)

func CreateDomainAttributesGenerator(
	dp providers.Provider[domain.Domain],
) shared.AttributesGenerator {
	return &domainAttributesGenerator{
		DomainProvider: dp,
	}
}

// Generate implements shared.AttributesGenerator.
func (g *domainAttributesGenerator) Generate() shared.AttributesChannel {
	ch := make(shared.AttributesChannel)
	go g.generateInternal(ch)
	return ch
}

func (g *domainAttributesGenerator) generateInternal(ch shared.AttributesChannel) {
	defer close(ch)

	// activate provider and store its channel
	pCh := g.DomainProvider.Provide()

	// receive data and wait until provider's channel is done
	for d := range pCh {
		atts := processDomainAttributes(d)
		// when attributes are ready send it
		ch <- atts
	}
}

func processDomainAttributes(d domain.Domain) shared.Attributes {
	m := make(shared.Attributes, domainCount)
	if len(d.Domain) > 0 {
		m[domainObject] = d.Domain
	}
	if len(d.FQDN) > 0 {
		m[domainFqdn] = d.FQDN
	}

	if len(d.Workgroup) > 0 {
		m[domainWorkgroup] = d.Workgroup
	}
	// when domain is missing, no role is assigned to attribute for windows
	if d.Domain != "" && d.DomainRole >= 0 {
		m[domainRole] = strconv.Itoa(d.DomainRole)
	}

	return m
}
