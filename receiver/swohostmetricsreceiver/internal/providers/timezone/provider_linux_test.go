package timezone

import (
	"fmt"
	"testing"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers"
	"github.com/stretchr/testify/assert"
)

func Test_Provide_ProvidesCompleteDataAndChannelIsClosedAfterDelivery(t *testing.T) {
	commandOutput := `                Time zone: Etc/UTC (UTC, +0000)`
	expectedTimeZone := TimeZone{
		Bias:    -1,
		Caption: "Etc/UTC (UTC, +0000)",
	}

	sut := provider{
		cli: providers.CreateCommandLineExecutorMock(commandOutput, "", nil),
	}

	ch := sut.Provide()
	actualTimeZone := <-ch
	_, open := <-ch // secondary receive

	assert.Equal(t, expectedTimeZone, actualTimeZone)
	assert.False(t, open, "channel must be closed")
}

func Test_Provide_FailsAndProvidesInvalidObjectAndChannelIsClosedAfterDelivery(t *testing.T) {
	expectedTimeZone := TimeZone{
		Bias: -1,
	}

	sut := provider{
		cli: providers.CreateCommandLineExecutorMock("", "", fmt.Errorf("something went wrong")),
	}

	ch := sut.Provide()
	actualDomain := <-ch
	_, open := <-ch // secondary receive

	assert.Equal(t, expectedTimeZone, actualDomain)
	assert.False(t, open, "channel must be closed")
}

func Test_Provide_ErrorIsInStdErrReturnsEmptyObjectAndChannelIsClosedAfterDelivery(t *testing.T) {
	expectedTimeZone := TimeZone{
		Bias: -1,
	}

	sut := provider{
		cli: providers.CreateCommandLineExecutorMock("some output", "some error in stderr", nil),
	}
	ch := sut.Provide()
	actualTimeZone := <-ch
	_, open := <-ch // secondary receive

	assert.Equal(t, expectedTimeZone, actualTimeZone)
	assert.False(t, open, "channel must be closed")
}
