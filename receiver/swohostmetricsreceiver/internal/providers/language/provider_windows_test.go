package language

import (
	"fmt"
	"testing"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/registry"

	"github.com/stretchr/testify/assert"
)

func Test_Provide_ProvidesFullLanguageAndChannelIsClosedAfterDelivery(t *testing.T) {
	expectedLanguage := Language{
		LCID:        999,
		Name:        "test-short",
		DisplayName: "Test Display Name",
	}

	sut := provider{
		lcidProvider:    &mockLCIDProvider{lcid: 999},
		displayNamesMap: map[string]string{"test-short": "Test Display Name"},
		getRegistryValues: func(_ registry.RootKey, _, _ string, _ []string) (map[string]string, error) {
			return map[string]string{"LocaleName": "test-short"}, nil
		},
	}

	ch := sut.Provide()
	actualLanguage := <-ch
	_, open := <-ch // secondary receive

	assert.Equal(t, expectedLanguage, actualLanguage)
	assert.False(t, open, "channel must be closed")
}

func Test_Provide_ShortNameGettingFailsAndProvidesEmptyObjectAndChannelIsClosedAfterDelivery(t *testing.T) {
	expectedLanguage := Language{}

	sut := provider{
		lcidProvider:    &mockLCIDProvider{lcid: 999},
		displayNamesMap: map[string]string{},
		getRegistryValues: func(_ registry.RootKey, _, _ string, _ []string) (map[string]string, error) {
			return nil, fmt.Errorf("no short name for you")
		},
	}

	ch := sut.Provide()
	actualLanguage := <-ch
	_, open := <-ch // secondary receive

	assert.Equal(t, expectedLanguage, actualLanguage)
	assert.False(t, open, "channel must be closed")
}

// stubs.
type mockLCIDProvider struct {
	lcid uint32
}

func (m *mockLCIDProvider) GetUserDefaultLCID() uint32 {
	return m.lcid
}
