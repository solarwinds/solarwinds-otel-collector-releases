// Copyright 2025 SolarWinds Worldwide, LLC. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal

import (
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"testing"
)

func TestSetIdAttributes(t *testing.T) {
	resourceAttrs := pcommon.NewMap()
	resourceAttrs.PutStr("id1", "idvalue1")
	resourceAttrs.PutStr("id2", "idvalue2")

	destination := plog.NewLogRecord()
	result := setIdAttributes(&destination, []string{"id1"}, resourceAttrs)
	assert.True(t, result)
	ids, exists := destination.Attributes().Get(swoEntityIds)
	assert.True(t, exists)
	assert.Equal(t, 1, ids.Map().Len())
	id, exists := ids.Map().Get("id1")
	assert.True(t, exists)
	assert.Equal(t, "idvalue1", id.Str())
}

func TestPutStringAttribute(t *testing.T) {
	resourceAttrs := pcommon.NewMap()
	resourceAttrs.PutStr("string", "stringValue")

	destination := pcommon.NewMap()
	result := copyAttribute(&destination, "string", &resourceAttrs)
	assert.True(t, result, "Attribute should be set successfully")
	inDest, exists := destination.Get("string")
	assert.True(t, exists, "Attribute should exist in destination")
	assert.Equal(t, "stringValue", inDest.Str(), "Attribute value should match")
}

func TestPutBoolAttribute(t *testing.T) {
	resourceAttrs := pcommon.NewMap()
	resourceAttrs.PutBool("bool", true)

	destination := pcommon.NewMap()
	result := copyAttribute(&destination, "bool", &resourceAttrs)
	assert.True(t, result, "Attribute should be set successfully")
	inDest, exists := destination.Get("bool")
	assert.True(t, exists, "Attribute should exist in destination")
	assert.Equal(t, true, inDest.Bool(), "Attribute value should match")
}

func TestPutIntAttribute(t *testing.T) {
	resourceAttrs := pcommon.NewMap()
	resourceAttrs.PutInt("int", 123)

	destination := pcommon.NewMap()
	result := copyAttribute(&destination, "int", &resourceAttrs)
	assert.True(t, result, "Attribute should be set successfully")
	inDest, exists := destination.Get("int")
	assert.True(t, exists, "Attribute should exist in destination")
	assert.Equal(t, int64(123), inDest.Int(), "Attribute value should match")
}

func TestPutDoubleAttribute(t *testing.T) {
	resourceAttrs := pcommon.NewMap()
	resourceAttrs.PutDouble("double", 123.456)

	destination := pcommon.NewMap()
	result := copyAttribute(&destination, "double", &resourceAttrs)
	assert.True(t, result, "Attribute should be set successfully")
	inDest, exists := destination.Get("double")
	assert.True(t, exists, "Attribute should exist in destination")
	assert.Equal(t, 123.456, inDest.Double(), "Attribute value should match")
}

func TestPutBytesAttribute(t *testing.T) {
	resourceAttrs := pcommon.NewMap()
	resourceAttrs.PutStr("bytes", "byteValue")

	destination := pcommon.NewMap()
	result := copyAttribute(&destination, "bytes", &resourceAttrs)
	assert.True(t, result, "Attribute should be set successfully")
	inDest, exists := destination.Get("bytes")
	assert.True(t, exists, "Attribute should exist in destination")
	assert.Equal(t, "byteValue", inDest.Str(), "Attribute value should match")
}

func TestPutNotExistingAttribute(t *testing.T) {
	resourceAttrs := pcommon.NewMap()

	destination := pcommon.NewMap()
	result := copyAttribute(&destination, "notExisting", &resourceAttrs)
	assert.False(t, result, "Attribute should not be set successfully")
	assert.Equal(t, 0, destination.Len(), "Attribute should not exist in destination")
}
