package hostdetails

import (
	"strconv"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/attributes/shared"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers/timezone"
)

// Following keys will be used in metric as attributes.
const (
	bias          = "hostdetails.timezone.bias"
	caption       = "hostdetails.timezone.caption"
	standardName  = "hostdetails.timezone.standardname"
	timeZoneCount = 3
)

type timeZoneAttributesGenerator struct {
	TimeZoneProvider providers.Provider[timezone.TimeZone]
}

var _ shared.AttributesGenerator = (*timeZoneAttributesGenerator)(nil)

func CreateTimeZoneAttributesGenerator(
	tp providers.Provider[timezone.TimeZone],
) shared.AttributesGenerator {
	return &timeZoneAttributesGenerator{
		TimeZoneProvider: tp,
	}
}

// Generate implements shared.AttributesGenerator.
func (g *timeZoneAttributesGenerator) Generate() shared.AttributesChannel {
	ch := make(shared.AttributesChannel)
	go g.generateInternal(ch)
	return ch
}

func (g *timeZoneAttributesGenerator) generateInternal(ch shared.AttributesChannel) {
	defer close(ch)

	// activate provider and store its channel
	pCh := g.TimeZoneProvider.Provide()

	// receive data and wait until provider's channel is done
	for t := range pCh {
		atts := processTimeZoneAttributes(t)
		// when attribute are ready send it
		ch <- atts
	}
}

func processTimeZoneAttributes(tz timezone.TimeZone) shared.Attributes {
	m := make(shared.Attributes, timeZoneCount)
	if len(tz.Caption) > 0 {
		m[caption] = tz.Caption
	}
	if len(tz.StandardName) > 0 {
		m[standardName] = tz.StandardName
	}
	if tz.Bias%60 == 0 {
		m[bias] = strconv.Itoa(tz.Bias)
	}
	return m
}
