package language

import (
	"fmt"
	"testing"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers"
	"github.com/stretchr/testify/assert"
)

func Test_Provide_ProvidesLanguageNameAndChannelIsClosedAfterDelivery(t *testing.T) {
	commandOutput := "LANGUAGE=en_US"

	expectedLanguage := Language{
		Name: "en_US",
	}

	sut := provider{
		cli: providers.CreateCommandLineExecutorMock(commandOutput, "", nil),
	}

	ch := sut.Provide()
	actualLanguage := <-ch
	_, open := <-ch // secondary receive

	assert.Equal(t, expectedLanguage, actualLanguage)
	assert.False(t, open, "channel must be closed")
}

func Test_Provide_ProvidesEmptyLanguageNameWhenLocaleNotSetAndChannelIsClosedAfterDelivery(t *testing.T) {
	commandOutput := "LANGUAGE="

	expectedLanguage := Language{}

	sut := provider{
		cli: providers.CreateCommandLineExecutorMock(commandOutput, "", nil),
	}

	ch := sut.Provide()
	actualLanguage := <-ch
	_, open := <-ch // secondary receive

	assert.Equal(t, expectedLanguage, actualLanguage)
	assert.False(t, open, "channel must be closed")
}

func Test_Provide_FailsAndProvidesEmptyObjectAndChannelIsClosedAfterDelivery(t *testing.T) {
	expectedLanguage := Language{}

	sut := provider{
		cli: providers.CreateCommandLineExecutorMock("", "", fmt.Errorf("kokoha error")),
	}

	ch := sut.Provide()
	actualLanguage := <-ch
	_, open := <-ch // secondary receive

	assert.Equal(t, expectedLanguage, actualLanguage)
	assert.False(t, open, "channel must be closed")
}

func Test_Provide_ErrorInStdErrItProvidesEmptyObjectAndChannelIsClosedAfterDelivery(t *testing.T) {
	expectedLanguage := Language{}

	sut := provider{
		cli: providers.CreateCommandLineExecutorMock("", "kokoha error", nil),
	}

	ch := sut.Provide()
	actualLanguage := <-ch
	_, open := <-ch // secondary receive

	assert.Equal(t, expectedLanguage, actualLanguage)
	assert.False(t, open, "channel must be closed")
}
