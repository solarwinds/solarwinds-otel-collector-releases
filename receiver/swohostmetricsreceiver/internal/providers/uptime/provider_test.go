package uptime

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Functional(t *testing.T) {
	t.Skip("This test should be run manually")

	sut := CreateUptimeProvider(
		CreateUptimeWrapper(),
	)

	result := <-sut.GetUptime()

	fmt.Printf("Result: %+v\n", result)
}

func Test_GetUptime_WhenSucceedsReturnsUptimeAndChannelIsClosedAfterDelivery(t *testing.T) {
	expectedUptime := (uint64)(1701)

	sut := CreateUptimeProvider(
		CreateSuceedingUptimeWrapper(expectedUptime),
	)

	ch := sut.GetUptime()
	actualUptime := <-ch
	_, open := <-ch // secondary receive

	assert.Equal(t, expectedUptime, actualUptime.Uptime)
	assert.Nil(t, actualUptime.Error)
	assert.False(t, open, "channel must be closed")
}

func Test_GetUptime_WhenFailsReturnsZeroUptimeWithErrorAndChannelIsClosedAfterDelivery(t *testing.T) {
	expectedError := fmt.Errorf("kokoha happened")

	sut := CreateUptimeProvider(
		CreateFailingUptimeWrapper(expectedError),
	)

	ch := sut.GetUptime()
	actualUptime := <-ch
	_, open := <-ch // secondary receive

	assert.Equal(t, expectedError, actualUptime.Error)
	assert.Zero(t, actualUptime.Uptime)
	assert.False(t, open, "channel must be closed")
}
