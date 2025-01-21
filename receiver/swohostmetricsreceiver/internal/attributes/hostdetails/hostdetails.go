package hostdetails

import (
	"sync"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/attributes/shared"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/synchronization"
)

type attributeGenerator struct {
	DomainAttributeGenerator   shared.AttributesGenerator
	ModelAttributeGenerator    shared.AttributesGenerator
	TimeZoneAttributeGenerator shared.AttributesGenerator
}

var _ shared.AttributesGenerator = (*attributeGenerator)(nil)

func CreateHostDetailsAttributesGenerator(
	domain shared.AttributesGenerator,
	model shared.AttributesGenerator,
	timezone shared.AttributesGenerator,
) shared.AttributesGenerator {
	return &attributeGenerator{
		DomainAttributeGenerator:   domain,
		ModelAttributeGenerator:    model,
		TimeZoneAttributeGenerator: timezone,
	}
}

// Generate implements shared.AttributesGenerator.
func (g *attributeGenerator) Generate() shared.AttributesChannel {
	ch := make(shared.AttributesChannel)
	go g.generateInternal(ch)
	return ch
}

func (g *attributeGenerator) generateInternal(ch shared.AttributesChannel) {
	defer close(ch)

	const attGeneratorsCount = 3
	var wg sync.WaitGroup
	wg.Add(attGeneratorsCount)

	terminationCh := synchronization.ActivateSupervisingRoutine(&wg)

	domainCh := g.DomainAttributeGenerator.Generate()
	modelCh := g.ModelAttributeGenerator.Generate()
	timezoneCh := g.TimeZoneAttributeGenerator.Generate()

	allAttsCount := domainCount + modelCount + timeZoneCount
	containerAtts := make(shared.Attributes, allAttsCount)

loop:
	for {
		select {
		case dAtts, opened := <-domainCh:
			shared.ProcessReceivedAttributes(dAtts, containerAtts, opened, &domainCh, &wg)
		case mAtts, opened := <-modelCh:
			shared.ProcessReceivedAttributes(mAtts, containerAtts, opened, &modelCh, &wg)
		case tzAtts, opened := <-timezoneCh:
			shared.ProcessReceivedAttributes(tzAtts, containerAtts, opened, &timezoneCh, &wg)
		case <-terminationCh:
			break loop
		}
	}

	ch <- containerAtts
}
