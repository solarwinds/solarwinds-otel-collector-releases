package discovery

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/cli"
)

func Test_Functional(t *testing.T) {
	t.Skip("This test should be ran manually only")

	sut := NewDpkgDiscovery()
	result := sut.Discover()
	assert.True(t, result)
}

func Test_Discover_OnCommandSuccessReturnsTrue(t *testing.T) {
	cliMock := cli.CreateNewCliExecutorMock("<whatever_command_output>", "", nil)
	sut := createDpkgDiscovery(cliMock)

	actual := sut.Discover()
	assert.True(t, actual, "discovery must be positive")
}

func Test_Discover_OnCommandFailureReturnsFalse(t *testing.T) {
	cliMock := cli.CreateNewCliExecutorMock("", "", fmt.Errorf("discovery command failure"))
	sut := createDpkgDiscovery(cliMock)

	actual := sut.Discover()
	assert.False(t, actual, "discovery must be negative")
}
