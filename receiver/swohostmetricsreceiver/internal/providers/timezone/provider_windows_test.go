package timezone

import (
	"fmt"
	"testing"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/wmi"

	"github.com/stretchr/testify/assert"
)

func Test_Provider_Functional(t *testing.T) {
	t.Skip("This test should be run manually")

	sut := CreateTimeZoneProvider()
	result := <-sut.Provide()
	fmt.Printf("Result: %+v\n", result)
}

func Test_Provide_ProvidesCompleteDataAndChannelIsClosedAfterDelivery(t *testing.T) {
	expectedTimeZone := TimeZone{
		Bias:         30,
		Caption:      "(UTC+11:00) Some Time Zone",
		StandardName: "Central Antarctic Standard Time",
	}

	timeZoneMock := []Win32_TimeZone{
		{
			Bias:         30,
			Caption:      "(UTC+11:00) Some Time Zone",
			StandardName: "Central Antarctic Standard Time",
		},
	}

	sut := provider{
		wmi.CreateWmiExecutorMock([]interface{}{&timeZoneMock}, nil),
	}

	ch := sut.Provide()
	actualTimeZone := <-ch
	_, open := <-ch // secondary receive

	assert.Equal(t, expectedTimeZone, actualTimeZone)
	assert.False(t, open, "channel must be closed")
}

func Test_Provide_FailsAndProvidesEmptyObjectAndChannelIsClosedAfterDelivery(t *testing.T) {
	expectedTimeZone := TimeZone{
		Bias:         0,
		Caption:      "",
		StandardName: "",
	}

	errors := map[interface{}]error{
		&[]Win32_TimeZone{}: fmt.Errorf("some timezone error"),
	}
	sut := provider{
		wmi: wmi.CreateWmiExecutorMock(nil, errors),
	}

	ch := sut.Provide()
	actualDomain := <-ch
	_, open := <-ch // secondary receive

	assert.Equal(t, expectedTimeZone, actualDomain)
	assert.False(t, open, "channel must be closed")
}
