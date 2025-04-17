package internal

import (
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"testing"
)

func TestPutStringAttribute(t *testing.T) {
	resourceAttrs := pcommon.NewMap()
	resourceAttrs.PutStr("string", "stringValue")

	destination := pcommon.NewMap()
	result := putAttribute(&destination, "string", &resourceAttrs)
	assert.True(t, result, "Attribute should be set successfully")
	inDest, exists := destination.Get("string")
	assert.True(t, exists, "Attribute should exist in destination")
	assert.Equal(t, "stringValue", inDest.Str(), "Attribute value should match")
}

func TestPutBoolAttribute(t *testing.T) {
	resourceAttrs := pcommon.NewMap()
	resourceAttrs.PutBool("bool", true)

	destination := pcommon.NewMap()
	result := putAttribute(&destination, "bool", &resourceAttrs)
	assert.True(t, result, "Attribute should be set successfully")
	inDest, exists := destination.Get("bool")
	assert.True(t, exists, "Attribute should exist in destination")
	assert.Equal(t, true, inDest.Bool(), "Attribute value should match")
}

func TestPutIntAttribute(t *testing.T) {
	resourceAttrs := pcommon.NewMap()
	resourceAttrs.PutInt("int", 123)

	destination := pcommon.NewMap()
	result := putAttribute(&destination, "int", &resourceAttrs)
	assert.True(t, result, "Attribute should be set successfully")
	inDest, exists := destination.Get("int")
	assert.True(t, exists, "Attribute should exist in destination")
	assert.Equal(t, int64(123), inDest.Int(), "Attribute value should match")
}

func TestPutDoubleAttribute(t *testing.T) {
	resourceAttrs := pcommon.NewMap()
	resourceAttrs.PutDouble("double", 123.456)

	destination := pcommon.NewMap()
	result := putAttribute(&destination, "double", &resourceAttrs)
	assert.True(t, result, "Attribute should be set successfully")
	inDest, exists := destination.Get("double")
	assert.True(t, exists, "Attribute should exist in destination")
	assert.Equal(t, 123.456, inDest.Double(), "Attribute value should match")
}

func TestPutBytesAttribute(t *testing.T) {
	resourceAttrs := pcommon.NewMap()
	resourceAttrs.PutStr("bytes", "byteValue")

	destination := pcommon.NewMap()
	result := putAttribute(&destination, "bytes", &resourceAttrs)
	assert.True(t, result, "Attribute should be set successfully")
	inDest, exists := destination.Get("bytes")
	assert.True(t, exists, "Attribute should exist in destination")
	assert.Equal(t, "byteValue", inDest.Str(), "Attribute value should match")
}

func TestPutNotExistingAttribute(t *testing.T) {
	resourceAttrs := pcommon.NewMap()

	destination := pcommon.NewMap()
	result := putAttribute(&destination, "notExisting", &resourceAttrs)
	assert.False(t, result, "Attribute should not be set successfully")
	assert.Equal(t, 0, destination.Len(), "Attribute should not exist in destination")
}
