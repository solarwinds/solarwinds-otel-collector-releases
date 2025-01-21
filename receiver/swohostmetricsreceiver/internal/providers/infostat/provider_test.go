package infostat

import (
	"fmt"
	"testing"

	"github.com/shirou/gopsutil/v3/host"
	"github.com/stretchr/testify/assert"
)

func Test_Functional(t *testing.T) {
	t.Skip("This test should be run manually")

	sut := CreateInfoStatProvider()
	result := <-sut.Provide()
	fmt.Printf("Result: %+v\n", result)
}

func Test_Provide_ProvidesCompleteDataAndChannelIsClosedAfterDelivery(t *testing.T) {
	expectedInfoStat := InfoStat{
		Hostname:             "HOSTNAME",
		BootTime:             123456789,
		Os:                   "windows",
		Platform:             "Microsoft Windows 42 Enterprise",
		PlatformFamily:       "Workstation",
		PlatformVersion:      "1.1 Build 2.2",
		KernelVersion:        "3.3 Build 4.4",
		KernelArchitecture:   "x86_64",
		VirtualizationSystem: "virtualization system",
		VirtualizationRole:   "role",
		HostID:               "id-1",
	}
	sut := provider{
		internalExecutor: CreateInfoStatProviderMock(basicInfoStat(), nil),
	}

	ch := sut.Provide()
	actualInfoStat := <-ch
	_, open := <-ch // secondary receive

	assert.Equal(t, expectedInfoStat, actualInfoStat)
	assert.False(t, open, "channel must be closed")
}

func Test_Provide_FailsAndProvidesEmptyObjectAndChannelIsClosedAfterDelivery(t *testing.T) {
	expectedInfoStat := InfoStat{
		Hostname:             "",
		BootTime:             0,
		Os:                   "",
		Platform:             "",
		PlatformFamily:       "",
		PlatformVersion:      "",
		KernelVersion:        "",
		KernelArchitecture:   "",
		VirtualizationSystem: "",
		VirtualizationRole:   "",
		HostID:               "",
	}
	sut := provider{
		internalExecutor: CreateInfoStatProviderMock(&host.InfoStat{}, fmt.Errorf("something went wrong")),
	}

	ch := sut.Provide()
	actualInfoStat := <-ch
	_, open := <-ch // secondary receive

	assert.Equal(t, expectedInfoStat, actualInfoStat)
	assert.False(t, open, "channel must be closed")
}

func basicInfoStat() *host.InfoStat {
	return &host.InfoStat{
		Hostname:             "HOSTNAME",
		Uptime:               10,
		BootTime:             123456789,
		Procs:                187,
		OS:                   "windows",
		Platform:             "Microsoft Windows 42 Enterprise",
		PlatformFamily:       "Workstation",
		PlatformVersion:      "1.1 Build 2.2",
		KernelVersion:        "3.3 Build 4.4",
		KernelArch:           "x86_64",
		VirtualizationSystem: "virtualization system",
		VirtualizationRole:   "role",
		HostID:               "id-1",
	}
}
