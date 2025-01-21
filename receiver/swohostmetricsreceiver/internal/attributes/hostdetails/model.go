package hostdetails

import (
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/attributes/shared"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers/model"
)

// Following keys will be used in metric as attributes.
const (
	modelManufacturer = "hostdetails.model.manufacturer"
	modelSerialnumber = "hostdetails.model.serialnumber"
	modelName         = "hostdetails.model.name"
	modelCount        = 3
)

type generator struct {
	ModelProvider providers.Provider[model.Model]
}

var _ shared.AttributesGenerator = (*generator)(nil)

func CreateModelAttributesGenerator(
	mp providers.Provider[model.Model],
) shared.AttributesGenerator {
	return &generator{
		ModelProvider: mp,
	}
}

// Generate implements shared.AttributesGenerator.
func (g *generator) Generate() shared.AttributesChannel {
	ch := make(shared.AttributesChannel)
	go g.generateInternal(ch)
	return ch
}

func (g *generator) generateInternal(
	ch shared.AttributesChannel,
) {
	defer close(ch)

	mCh := g.ModelProvider.Provide()
	for m := range mCh {
		atts := processModelAttributes(m)
		ch <- atts
	}
}

func processModelAttributes(model model.Model) shared.Attributes {
	m := make(shared.Attributes, modelCount)
	if model.Manufacturer != "" {
		m[modelManufacturer] = model.Manufacturer
	}
	if model.Name != "" {
		m[modelName] = model.Name
	}
	if model.SerialNumber != "" {
		m[modelSerialnumber] = model.SerialNumber
	}
	return m
}
