package hostdetails

import (
	"fmt"
	"testing"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/attributes/shared"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers/domain"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers/model"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers/timezone"
	"github.com/stretchr/testify/assert"
)

func Test_HostDetailsAttributesGenerator_Functional(t *testing.T) {
	t.Skip("This test should be run manually")

	sut := CreateHostDetailsAttributesGenerator(
		CreateDomainAttributesGenerator(
			domain.CreateDomainProvider(),
		),
		CreateModelAttributesGenerator(
			model.CreateModelProvider(),
		),
		CreateTimeZoneAttributesGenerator(
			timezone.CreateTimeZoneProvider(),
		),
	)

	result := <-sut.Generate()

	fmt.Printf("Result %+v\n", result)
}

func Test_Generate_AllGeneratorsGenerates_GeneratesFullAttributesCollection(t *testing.T) {
	domainAttributes := shared.Attributes{
		"domain.mock.empty": "",
		"domain.mock.full":  "any",
	}
	modelAttributes := shared.Attributes{
		"model.mock.empty": "",
		"model.mock.full":  "any",
	}
	timezoneAttributes := shared.Attributes{
		"timezone.mock.empty": "",
		"timezone.mock.full":  "any",
	}

	expectedAttributes := shared.Attributes{
		"domain.mock.empty":   "",
		"domain.mock.full":    "any",
		"model.mock.empty":    "",
		"model.mock.full":     "any",
		"timezone.mock.empty": "",
		"timezone.mock.full":  "any",
	}

	sut := CreateHostDetailsAttributesGenerator(
		shared.CreateAttributesGeneratorMock(domainAttributes),
		shared.CreateAttributesGeneratorMock(modelAttributes),
		shared.CreateAttributesGeneratorMock(timezoneAttributes),
	)

	actualAttributes := <-sut.Generate()

	assert.Equal(t, expectedAttributes, actualAttributes)
}

func Test_Generate_SomeGeneratorsFails_GeneratesOnlyPartialCollection(t *testing.T) {
	domainAttributes := shared.Attributes{
		"domain.mock.empty": "",
		"domain.mock.full":  "any",
	}
	timezoneAttributes := shared.Attributes{
		"timezone.mock.empty": "",
		"timezone.mock.full":  "any",
	}

	expectedAttributes := shared.Attributes{
		"domain.mock.empty":   "",
		"domain.mock.full":    "any",
		"timezone.mock.empty": "",
		"timezone.mock.full":  "any",
	}

	sut := CreateHostDetailsAttributesGenerator(
		shared.CreateAttributesGeneratorMock(domainAttributes),
		shared.CreateAttributesGeneratorMock(nil), // model attributes
		shared.CreateAttributesGeneratorMock(timezoneAttributes),
	)

	actualAttributes := <-sut.Generate()

	assert.Equal(t, expectedAttributes, actualAttributes)
}

func Test_Generate_AllGeneratorsFails_GeneratesEmptyCollection(t *testing.T) {
	sut := CreateHostDetailsAttributesGenerator(
		shared.CreateAttributesGeneratorMock(nil),
		shared.CreateAttributesGeneratorMock(nil),
		shared.CreateAttributesGeneratorMock(nil),
	)

	actualAttributes := <-sut.Generate()

	assert.Equal(t, shared.Attributes{}, actualAttributes)
}
