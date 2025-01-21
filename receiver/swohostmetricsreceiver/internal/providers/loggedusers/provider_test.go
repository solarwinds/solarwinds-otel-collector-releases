//go:build !integration

package loggedusers

import (
	"fmt"
	"os"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Test_Functional(t *testing.T) {
	t.Skip("This test should be run manually")

	// Mimics previous version of logger setup.
	zap.ReplaceGlobals(
		zap.New(
			zapcore.NewCore(
				zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
				zapcore.AddSync(os.Stdout),
				zap.NewAtomicLevelAt(zapcore.DebugLevel),
			),
		),
	)

	sut := CreateProvider()
	result := <-sut.Provide()
	fmt.Printf("Result: %+v\n", result)
}
