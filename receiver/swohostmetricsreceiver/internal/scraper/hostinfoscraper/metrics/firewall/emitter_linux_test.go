//go:build !integration

package firewall

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Linux is not supported but metric must not fail.
func Test_Emit_EmptyMetricSliceIsProvided(t *testing.T) {
	sut := NewEmitter()
	err := sut.Init()
	assert.Nil(t, err, "emitter initialization must not fail on Linux")
	er := sut.Emit()
	assert.Nil(t, er.Error, "emitter's emit must not fail on Linux")
	assert.Zero(t, er.Data.Len(), "emit must provided zero metric slices")
}
