package hostdetails

import (
	"fmt"
	"testing"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/attributes/shared"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers/timezone"
	"github.com/stretchr/testify/assert"
)

func Test_TimezoneAttributesGenerator_Functional(t *testing.T) {
	t.Skip("This test should be run manually")

	sut := CreateTimeZoneAttributesGenerator(timezone.CreateTimeZoneProvider())
	result := <-sut.Generate()

	fmt.Printf("Result %+v\n", result)
}

func Test_Generate_TimeZoneIsProvided_AttributesAreGenerated(t *testing.T) {
	timeZoneMock := timezone.TimeZone{
		Bias:         120,
		Caption:      "(UTC+02:00) City",
		StandardName: "Not Europe Standard Time",
	}
	expectedAttributes := shared.Attributes{
		"hostdetails.timezone.bias":         "120",
		"hostdetails.timezone.caption":      "(UTC+02:00) City",
		"hostdetails.timezone.standardname": "Not Europe Standard Time",
	}
	sut := CreateTimeZoneAttributesGenerator(
		shared.CreateProviderMock(timeZoneMock), // send valid data
	)

	actualAttributes := <-sut.Generate()

	assert.Equal(t, expectedAttributes, actualAttributes, "expected attributes are not the same as actual")
}

func Test_Generate_TimeZoneDataUnavailable_AttributesAreNotGenerated(t *testing.T) {
	sut := CreateTimeZoneAttributesGenerator(shared.CreateEmptyProviderMock[timezone.TimeZone]()) // no data, closing the channel only

	actualAttributes := <-sut.Generate()

	assert.Equal(t, shared.Attributes(nil), actualAttributes, "expected atrributes are not nil")
}
