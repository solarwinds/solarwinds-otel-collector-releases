//go:build !integration

package firewall

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Emitter_Functional(t *testing.T) {
	t.Skip("This test should be run manually")

	sut := NewEmitter()
	if err := sut.Init(); err != nil {
		assert.Fail(t, "initialization must not fail")
	}
	er := sut.Emit()
	if er.Error != nil {
		assert.Fail(t, "metric emitation must not fail")
	}
	fmt.Printf("Result: %+v", er.Data)
}
