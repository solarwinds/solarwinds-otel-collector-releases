package cpu

import (
	"fmt"
	"testing"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/wmi"

	"github.com/stretchr/testify/assert"
)

func Test_Provide_ProvidesCompleteDataAndChannelIsClosedAfterDelivery(t *testing.T) {
	wmiOutput := []Win32_Processor{
		{
			Name:                      "Proc 1",
			Manufacturer:              "Manufacturer 1",
			CurrentClockSpeed:         10,
			NumberOfCores:             4,
			NumberOfLogicalProcessors: 5,
			Stepping:                  "1",
			Caption:                   "Some Caption With No Extra Data",
		},
		{
			Name:                      "Proc 2",
			Manufacturer:              "Manufacturer 2",
			CurrentClockSpeed:         3,
			NumberOfCores:             2,
			NumberOfLogicalProcessors: 6,
			Stepping:                  "",
			Caption:                   "Some Caption With Model X and Stepping 50",
		},
	}
	expectedProcessors := []Processor{
		{
			Name:         "Proc 1",
			Manufacturer: "Manufacturer 1",
			Speed:        10,
			Cores:        4,
			Threads:      5,
			Model:        "",
			Stepping:     "1",
		},
		{
			Name:         "Proc 2",
			Manufacturer: "Manufacturer 2",
			Speed:        3,
			Cores:        2,
			Threads:      6,
			Model:        "X",
			Stepping:     "50",
		},
	}
	expectedModel := Container{
		Processors: expectedProcessors,
		Error:      nil,
	}

	sut := provider{
		wmi: wmi.CreateWmiExecutorMock([]interface{}{&wmiOutput}, nil),
	}

	ch := sut.Provide()
	actualModel := <-ch
	_, open := <-ch // secondary receive

	assert.Equal(t, expectedModel, actualModel)
	assert.False(t, open, "channel must be closed")
}

func Test_Provide_FailsAndProvidesEmptyObjectAndChannelIsClosedAfterDelivery(t *testing.T) {
	err := fmt.Errorf("processor error")
	expectedModel := Container{
		Processors: nil,
		Error:      err,
	}

	sut := provider{
		wmi: wmi.CreateWmiExecutorMock(nil, map[interface{}]error{
			&[]Win32_Processor{}: err,
		}),
	}

	ch := sut.Provide()
	actualModel := <-ch
	_, open := <-ch // secondary receive

	assert.Equal(t, expectedModel, actualModel)
	assert.False(t, open, "channel must be closed")
}
