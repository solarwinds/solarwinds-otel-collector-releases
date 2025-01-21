package loggedusers

import (
	"fmt"
	"testing"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/registry"

	"github.com/stretchr/testify/assert"
)

func Test_Provide_RetrievingLastLoggedUserProvidesDataAndChannelIsClosedAfterDelivery(t *testing.T) {
	registryValues := map[string]string{
		`LastLoggedOnUser`:        `TestDomain/TestUser`,
		`LastLoggedOnDisplayName`: `User Test`,
	}

	expectedResult := Data{
		Users: []User{{
			Name:        `TestDomain/TestUser`,
			DisplayName: `User Test`,
		}},
	}

	sut := provider{
		getRegistryValues: func(_ registry.RootKey, _, _ string, _ []string) (map[string]string, error) {
			return registryValues, nil
		},
	}

	ch := sut.Provide()
	actualResult := <-ch
	_, open := <-ch // secondary receive

	assert.Equal(t, expectedResult, actualResult)
	assert.False(t, open, `channel must be closed`)
}

func Test_Provide_RetrievingLastLoggedUserFailsProvidesErrorAndChannelIsClosedAfterDelivery(t *testing.T) {
	registryError := fmt.Errorf(`cardinal mistake`)
	expectedResult := Data{Error: registryError}

	sut := provider{
		getRegistryValues: func(_ registry.RootKey, _, _ string, _ []string) (map[string]string, error) {
			return nil, registryError
		},
	}

	ch := sut.Provide()
	actualResult := <-ch
	_, open := <-ch // secondary receive

	assert.Equal(t, expectedResult, actualResult)
	assert.False(t, open, `channel must be closed`)
}
