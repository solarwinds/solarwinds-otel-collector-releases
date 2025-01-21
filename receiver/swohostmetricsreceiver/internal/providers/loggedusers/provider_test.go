//go:build !integration

package loggedusers

import (
	"fmt"
	"testing"

	"github.com/solarwinds-cloud/uams-plugin-lib/pkg/logger"
)

func Test_Functional(t *testing.T) {
	t.Skip("This test should be run manually")

	_ = logger.Setup(logger.WithLogToStdout(true))

	sut := CreateProvider()
	result := <-sut.Provide()
	fmt.Printf("Result: %+v\n", result)
}
