package osdetails

import (
	"fmt"
	"testing"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/attributes/shared"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers/infostat"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers/language"
	"github.com/stretchr/testify/assert"
)

func Test_OsDetailsAttributesGenerator_Functional(t *testing.T) {
	t.Skip("This test should be run manually")

	sut := CreateOsDetailsAttributesGenerator(
		CreateInfoStatAttributesGenerator(
			infostat.CreateInfoStatProvider(),
		),
		CreateLanguageAttributesGenerator(
			language.CreateLanguageProvider(),
		),
	)

	result := <-sut.Generate()

	fmt.Printf("Result %+v\n", result)
}

func Test_Generate_AllGeneratorsGenerates_GeneratesFullAttributesCollection(t *testing.T) {
	infoStatAttributes := shared.Attributes{
		"domain.mock.empty": "",
		"domain.mock.full":  "any",
	}

	languageAttributes := shared.Attributes{
		"language.mock.empty": "",
		"language.mock.full":  "any",
	}

	expectedAttributes := shared.Attributes{
		"domain.mock.empty":   "",
		"domain.mock.full":    "any",
		"language.mock.empty": "",
		"language.mock.full":  "any",
	}

	sut := CreateOsDetailsAttributesGenerator(
		shared.CreateAttributesGeneratorMock(infoStatAttributes),
		shared.CreateAttributesGeneratorMock(languageAttributes),
	)

	actualAttributes := <-sut.Generate()

	assert.Equal(t, expectedAttributes, actualAttributes)
}

func Test_Generate_SomeGeneratorsFails_GeneratesOnlyPartialCollection(t *testing.T) {
	infoStatAttributes := shared.Attributes{
		"infostat.mock.empty": "",
		"infostat.mock.full":  "any",
	}

	expectedAttributes := shared.Attributes{
		"infostat.mock.empty": "",
		"infostat.mock.full":  "any",
	}

	sut := CreateOsDetailsAttributesGenerator(
		shared.CreateAttributesGeneratorMock(infoStatAttributes),
		shared.CreateAttributesGeneratorMock(shared.Attributes(nil)), // language attributes
	)

	actualAttributes := <-sut.Generate()

	assert.Equal(t, expectedAttributes, actualAttributes)
}

func Test_Generate_AllGeneratorsFails_GeneratesEmptyCollection(t *testing.T) {
	sut := CreateOsDetailsAttributesGenerator(
		shared.CreateAttributesGeneratorMock(shared.Attributes(nil)),
		shared.CreateAttributesGeneratorMock(shared.Attributes(nil)), // language attributes
	)

	actualAttributes := <-sut.Generate()

	assert.Equal(t, shared.Attributes{}, actualAttributes)
}
