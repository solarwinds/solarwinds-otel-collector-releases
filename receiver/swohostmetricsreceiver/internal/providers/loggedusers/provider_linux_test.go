package loggedusers

import (
	"fmt"
	"testing"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers"

	"github.com/stretchr/testify/assert"
)

func Test_Provide_RetrievingLastLoggedUsersProvidesDataAndChannelIsClosedAfterDelivery(t *testing.T) {
	// trailing whitespaces have to be tested
	commandOutput := `ubuntu   pts/0        Mon Dec  4 17:28   still logged in   
ubuntu   pts/1        Mon Dec  4 17:28   still logged in
reboot   system boot  Mon Dec 15 09:05   still running
reboot   system boot  Fri Dec  4 10:15 - 12:25  (02:10)`
	expectedResult := Data{
		Users: []User{{
			Name: "ubuntu",
			TTY:  "pts/0",
		}, {
			Name: "ubuntu",
			TTY:  "pts/1",
		}},
	}
	sut := provider{
		cli: providers.CreateCommandLineExecutorMock(commandOutput, "", nil),
	}

	ch := sut.Provide()
	actualResult := <-ch
	_, open := <-ch // secondary receive

	assert.Equal(t, expectedResult, actualResult)
	assert.False(t, open, `channel must be closed`)
}

func Test_Provide_IgnoresStillRunningSessionAndChannelIsClosedAfterDelivery(t *testing.T) {
	commandOutput := `reboot   system boot  Mon Dec  4 09:05   still running`
	expectedResult := Data{}
	sut := provider{
		cli: providers.CreateCommandLineExecutorMock(commandOutput, "", nil),
	}

	ch := sut.Provide()
	actualResult := <-ch
	_, open := <-ch // secondary receive

	assert.Equal(t, expectedResult, actualResult)
	assert.False(t, open, `channel must be closed`)
}

func Test_Provide_ReturnsErrorIndicationOnStderrErrorAndChannelIsClosedAfterDelivery(t *testing.T) {
	expectedResult := Data{
		Error: fmt.Errorf("loggedusers provider stderr"),
	}
	sut := provider{
		cli: providers.CreateCommandLineExecutorMock("", "loggedusers provider stderr", nil),
	}

	ch := sut.Provide()
	actualResult := <-ch
	_, open := <-ch // secondary receive

	assert.Equal(t, expectedResult, actualResult)
	assert.False(t, open, `channel must be closed`)
}

func Test_Provide_ReturnsErrorIndicationOnErrorAndChannelIsClosedAfterDelivery(t *testing.T) {
	expectedResult := Data{
		Error: fmt.Errorf("loggedusers provider error"),
	}
	sut := provider{
		cli: providers.CreateCommandLineExecutorMock("", "", fmt.Errorf("loggedusers provider error")),
	}

	ch := sut.Provide()
	actualResult := <-ch
	_, open := <-ch // secondary receive

	assert.Equal(t, expectedResult, actualResult)
	assert.False(t, open, `channel must be closed`)
}

func Test_Provide_ReturnsEmptyObjectWhenNoUserSessionIsOnAndChannelIsClosedAfterDelivery(t *testing.T) {
	expectedResult := Data{}
	sut := provider{
		cli: providers.CreateCommandLineExecutorMock("", "", nil),
	}

	ch := sut.Provide()
	actualResult := <-ch
	_, open := <-ch // secondary receive

	assert.Equal(t, expectedResult, actualResult)
	assert.False(t, open, `channel must be closed`)
}
