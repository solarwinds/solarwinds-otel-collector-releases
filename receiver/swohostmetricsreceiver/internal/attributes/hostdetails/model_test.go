package hostdetails

import (
	"fmt"
	"testing"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/attributes/shared"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers/model"
	"github.com/stretchr/testify/assert"
)

func Test_ModelAttributesGenerator_Functional(t *testing.T) {
	t.Skip("This test should be run only manually")

	sut := CreateModelAttributesGenerator(
		model.CreateModelProvider(),
	)

	result := <-sut.Generate()

	fmt.Printf("Result: %+v\n", result)
}

func Test_Generate_OnFullModel_AllAttributesAreGenerated(t *testing.T) {
	modelName := "USS Enterprise"
	modelManufacturer := "Utopia Planitia Shipyards"
	modelSerialNumber := "NCC-1701E"

	providedModel := model.Model{
		SerialNumber: modelSerialNumber,
		Manufacturer: modelManufacturer,
		Name:         modelName,
	}

	expectedAttributes := shared.Attributes{
		"hostdetails.model.manufacturer": modelManufacturer,
		"hostdetails.model.name":         modelName,
		"hostdetails.model.serialnumber": modelSerialNumber,
	}

	sut := CreateModelAttributesGenerator(
		shared.CreateProviderMock(
			providedModel,
		),
	)

	actualAttributes := <-sut.Generate()

	assert.Equal(t, expectedAttributes, actualAttributes)
}

func Test_Generate_OnPartialModel_SomeAttributesAreGenerated(t *testing.T) {
	modelName := "USS Enterprise"
	modelManufacturer := "Utopia Planitia Shipyards"
	modelSerialNumber := ""

	providedModel := model.Model{
		SerialNumber: modelSerialNumber,
		Manufacturer: modelManufacturer,
		Name:         modelName,
	}

	expectedAttributes := shared.Attributes{
		"hostdetails.model.manufacturer": modelManufacturer,
		"hostdetails.model.name":         modelName,
	}

	sut := CreateModelAttributesGenerator(
		shared.CreateProviderMock(
			providedModel,
		),
	)

	actualAttributes := <-sut.Generate()

	assert.Equal(
		t,
		expectedAttributes,
		actualAttributes,
		"expected attributes are not the same as actual",
	)
}

func Test_Generate_OnNoModel_NilAttributesAreGenerated(t *testing.T) {
	sut := CreateModelAttributesGenerator(
		shared.CreateEmptyProviderMock[model.Model](),
	)

	actualAttributes := <-sut.Generate()

	assert.Equal(t, shared.Attributes(nil), actualAttributes)
}
