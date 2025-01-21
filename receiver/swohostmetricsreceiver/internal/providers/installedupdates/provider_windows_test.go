package installedupdates

import (
	"fmt"
	"testing"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/wmi"

	"github.com/stretchr/testify/require"
)

func Test_ProvidesInstalledUpdates(t *testing.T) {
	expectedUpdates := []InstalledUpdate{
		{
			Caption:     "https://support.microsoft.com/help/111111",
			Description: "Security Update",
			HotFixID:    "KB111111",
			InstalledBy: "NT AUTHORITY\\SYSTEM",
			InstalledOn: "2021-3-21",
		},
		{
			Caption:     "",
			Description: "Update",
			HotFixID:    "KB222222",
			InstalledBy: "NT AUTHORITY\\SYSTEM",
			InstalledOn: "2022-3-21",
		},
		{
			Caption:     "http://support.microsoft.com/?kbid=333333",
			Description: "Update",
			HotFixID:    "KB333333",
			InstalledBy: "",
			InstalledOn: "2023-3-21",
		},
	}
	updatesMock := []Win32_QuickFixEngineering{
		{
			Caption:     "https://support.microsoft.com/help/111111",
			Description: "Security Update",
			HotFixID:    "KB111111",
			InstalledBy: "NT AUTHORITY\\SYSTEM",
			InstalledOn: "3/21/2021",
		},
		{
			Caption:     "",
			Description: "Update",
			HotFixID:    "KB222222",
			InstalledBy: "NT AUTHORITY\\SYSTEM",
			InstalledOn: "3/21/2022",
		},
		{
			Caption:     "http://support.microsoft.com/?kbid=333333",
			Description: "Update",
			HotFixID:    "KB333333",
			InstalledBy: "",
			InstalledOn: "3/21/2023",
		},
	}

	sut := &windowsProvider{
		wmi: wmi.CreateWmiExecutorMock([]interface{}{&updatesMock}, nil),
	}

	installedUpdates, err := sut.GetUpdates()
	require.NoError(t, err)
	require.Equal(t, expectedUpdates, installedUpdates)
}

func Test_ProvidesEmptyUpdatesOnFailure(t *testing.T) {
	expectedUpdates := []InstalledUpdate{}

	sut := &windowsProvider{
		wmi: wmi.CreateWmiExecutorMock(nil, map[interface{}]error{
			&[]Win32_QuickFixEngineering{}: fmt.Errorf("no updates for you"),
		}),
	}

	installedUpdates, err := sut.GetUpdates()
	require.Error(t, err)
	require.Equal(t, expectedUpdates, installedUpdates)
}
