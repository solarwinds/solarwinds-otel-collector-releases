package osdetails

import (
	"sync"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/attributes/shared"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/synchronization"
)

type attributesGenerator struct {
	InfoStatAttributesGenerator shared.AttributesGenerator
	LanguageAttributesGenerator shared.AttributesGenerator
}

// Generate implements shared.AttributesGenerator.
func (g *attributesGenerator) Generate() shared.AttributesChannel {
	ch := make(shared.AttributesChannel)
	go g.generateInternal(ch)
	return ch
}

var _ shared.AttributesGenerator = (*attributesGenerator)(nil)

func CreateOsDetailsAttributesGenerator(
	infostat shared.AttributesGenerator,
	language shared.AttributesGenerator,
) shared.AttributesGenerator {
	return &attributesGenerator{
		InfoStatAttributesGenerator: infostat,
		LanguageAttributesGenerator: language,
	}
}

func (g *attributesGenerator) generateInternal(ch shared.AttributesChannel) {
	defer close(ch)

	// prepare for language attributes generator
	const attGeneratorsCount = 2
	var wg sync.WaitGroup
	wg.Add(attGeneratorsCount)

	terminationCh := synchronization.ActivateSupervisingRoutine(&wg)

	infoStatCh := g.InfoStatAttributesGenerator.Generate()
	languageCh := g.LanguageAttributesGenerator.Generate()

	allAttsCount := infoStatCount + languageCount
	osDetailsCh := make(shared.Attributes, allAttsCount)

loop:
	for {
		select {
		case isAtts, opened := <-infoStatCh:
			shared.ProcessReceivedAttributes(isAtts, osDetailsCh, opened, &infoStatCh, &wg)
		case lAtts, opened := <-languageCh:
			shared.ProcessReceivedAttributes(lAtts, osDetailsCh, opened, &languageCh, &wg)
		case <-terminationCh:
			break loop
		}
	}
	ch <- osDetailsCh
}
