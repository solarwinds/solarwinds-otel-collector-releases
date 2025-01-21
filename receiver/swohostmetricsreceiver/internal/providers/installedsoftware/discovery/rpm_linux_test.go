package discovery

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/cli"
)

func Test_Rpm_Functional(t *testing.T) {
	t.Skip("This test should be ran manually only")

	sut := NewRpmDiscovery()
	result := sut.Discover()
	assert.True(t, result)
}

func Test_RpmDiscover_OnCommandSuccessReturnsTrue(t *testing.T) {
	cliMock := cli.CreateNewCliExecutorMock("<whatever_command_output>", "", nil)
	sut := createRpmDiscovery(cliMock)

	actual := sut.Discover()
	assert.True(t, actual, "discovery must be positive")
}

func Test_RpmDiscover_OnCommandFailureReturnsFalse(t *testing.T) {
	cliMock := cli.CreateNewCliExecutorMock("", "", fmt.Errorf("discovery command failure"))
	sut := createRpmDiscovery(cliMock)

	actual := sut.Discover()
	assert.False(t, actual, "discovery must be negatove")
}
