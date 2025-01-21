package domain

import (
	"fmt"
	"testing"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/wmi"

	"github.com/stretchr/testify/assert"
)

func Test_Provider_Functional(t *testing.T) {
	t.Skip("This test should be run manually")

	sut := CreateDomainProvider()
	result := <-sut.Provide()
	fmt.Printf("Result: %+v\n", result)
}

func Test_Provide_ProvidesCompleteDataAndChannelIsClosedAfterDelivery(t *testing.T) {
	wmiOutput := []Win32_ComputerSystem{
		{
			Name:        "SWI-D5ZRKQ2",
			Domain:      "swdev.local",
			DNSHostName: "SWI-D5ZRKQ2",
			DomainRole:  1,
		},
	}

	expectedDomain := Domain{
		Domain:     "swdev.local",
		FQDN:       "SWI-D5ZRKQ2.swdev.local",
		DomainRole: 1,
		Workgroup:  "",
	}

	sut := provider{
		wmi: wmi.CreateWmiExecutorMock([]interface{}{&wmiOutput}, nil),
	}

	ch := sut.Provide()
	actualDomain := <-ch
	_, open := <-ch // secondary receive

	assert.Equal(t, expectedDomain, actualDomain)
	assert.False(t, open, "channel must be closed")
}

func Test_Provide_FailsAndProvidesEmptyObjectAndChannelIsClosedAfterDelivery(t *testing.T) {
	expectedDomain := Domain{
		Domain:     "",
		FQDN:       "",
		DomainRole: 0,
		Workgroup:  "",
	}

	sut := provider{
		wmi: wmi.CreateWmiExecutorMock(nil, map[interface{}]error{
			&[]Win32_ComputerSystem{}: fmt.Errorf("kokoha error"),
		}),
	}

	ch := sut.Provide()
	actualDomain := <-ch
	_, open := <-ch // secondary receive

	assert.Equal(t, expectedDomain, actualDomain)
	assert.False(t, open, "channel must be closed")
}
