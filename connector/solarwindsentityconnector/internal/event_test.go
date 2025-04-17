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

func TestNewEvents(t *testing.T) {
	logs := plog.NewLogs()
	events := NewEvents(logs)

	assert.NotNil(t, events)
	assert.NotNil(t, logs)
	assert.Equal(t, 1, logs.ResourceLogs().Len())
	assert.Equal(t, 1, logs.ResourceLogs().At(0).ScopeLogs().Len())
	assert.Equal(t, 0, logs.LogRecordCount())
	assert.Equal(t, *events.logRecords, logs.ResourceLogs().At(0).ScopeLogs().At(0).LogRecords())
	assertOtelEventAsLogIsPresent(t, logs)
}

func TestAppendEntityUpdateEventWhenAttributesArePresent(t *testing.T) {
	// arrange
	logs := plog.NewLogs()
	events := NewEvents(logs)
	testEntity := NewEntity("testEntityType", []string{"id1", "id2"}, []string{"attr1", "attr2"})
	resourceAttrs := pcommon.NewMap()
	resourceAttrs.PutStr("id1", "idvalue1")
	resourceAttrs.PutStr("id2", "idvalue2")
	resourceAttrs.PutStr("attr1", "attrvalue1")
	resourceAttrs.PutStr("attr2", "attrvalue2")

	// act
	events.AppendEntityUpdateEvent(testEntity, resourceAttrs)

	// assert
	assert.Equal(t, 1, logs.LogRecordCount())
	actualLogRecord := logs.ResourceLogs().At(0).ScopeLogs().At(0).LogRecords().At(0)
	assertEntityType(t, actualLogRecord.Attributes(), testEntity.entityType)
	assertEventType(t, actualLogRecord.Attributes(), entityUpdateEventType)

	ids := getMap(actualLogRecord.Attributes(), swoEntityIds)
	assert.Equal(t, 2, ids.Len())
	assertAttributeIsPresent(t, ids, "id1", "idvalue1")
	assertAttributeIsPresent(t, ids, "id2", "idvalue2")

	attrs := getMap(actualLogRecord.Attributes(), swoEntityAttributes)
	assert.Equal(t, 2, attrs.Len())
	assertAttributeIsPresent(t, attrs, "attr1", "attrvalue1")
	assertAttributeIsPresent(t, attrs, "attr2", "attrvalue2")
	assertOtelEventAsLogIsPresent(t, logs)
}

func TestAppendEntityUpdateEventWhenIDAttributeIsMissing(t *testing.T) {
	// arrange
	logs := plog.NewLogs()
	events := NewEvents(logs)
	testEntity := NewEntity("testEntityType", []string{"id1", "id2"}, []string{})
	resourceAttrs := pcommon.NewMap()
	resourceAttrs.PutStr("id1", "idvalue1")

	// act
	events.AppendEntityUpdateEvent(testEntity, resourceAttrs)

	// assert
	assert.Equal(t, *events.logRecords, logs.ResourceLogs().At(0).ScopeLogs().At(0).LogRecords())
	assertOtelEventAsLogIsPresent(t, logs)
}

func TestAppendEntityUpdateEventWhenAttributeIsMissing(t *testing.T) {
	// arrange
	logs := plog.NewLogs()
	events := NewEvents(logs)
	testEntity := NewEntity("testEntityType", []string{"id1"}, []string{"attr1", "attr2"})
	resourceAttrs := pcommon.NewMap()
	resourceAttrs.PutStr("id1", "idvalue1")
	resourceAttrs.PutStr("attr1", "attrvalue1")

	// act
	events.AppendEntityUpdateEvent(testEntity, resourceAttrs)

	// assert
	assert.Equal(t, 1, logs.LogRecordCount())
	actualLogRecord := logs.ResourceLogs().At(0).ScopeLogs().At(0).LogRecords().At(0)
	assertEntityType(t, actualLogRecord.Attributes(), testEntity.entityType)
	assertEventType(t, actualLogRecord.Attributes(), entityUpdateEventType)
	assertOtelEventAsLogIsPresent(t, logs)

	ids := getMap(actualLogRecord.Attributes(), swoEntityIds)
	assert.Equal(t, 1, ids.Len())
	assertAttributeIsPresent(t, ids, "id1", "idvalue1")

	attrs := getMap(actualLogRecord.Attributes(), swoEntityAttributes)
	assert.Equal(t, 1, attrs.Len())
	assertAttributeIsPresent(t, attrs, "attr1", "attrvalue1")
}

func assertEventType(t *testing.T, attrs pcommon.Map, expected string) {
	if val, ok := attrs.Get(swoEntityEventType); ok {
		assert.Equal(t, true, ok)
		assert.Equal(t, expected, val.Str())
	}
}

func assertEntityType(t *testing.T, attrs pcommon.Map, expected string) {
	if val, ok := attrs.Get(swoEntityType); ok {
		assert.Equal(t, true, ok)
		assert.Equal(t, expected, val.Str())
	}
}

func assertAttributeIsPresent(t *testing.T, attrs pcommon.Map, key string, expected string) {
	if val, ok := attrs.Get(key); ok {
		assert.Equal(t, true, ok)
		assert.Equal(t, expected, val.Str())
	}
}

func assertOtelEventAsLogIsPresent(t *testing.T, logs plog.Logs) {
	isEntityEvent, ok := logs.ResourceLogs().At(0).ScopeLogs().At(0).Scope().Attributes().Get(swoEntityEventAsLog)
	assert.Equal(t, true, ok)
	assert.Equal(t, true, isEntityEvent.Bool())
}

func getMap(attrs pcommon.Map, key string) pcommon.Map {
	if val, ok := attrs.Get(key); ok {
		return val.Map()
	}
	return pcommon.Map{}
}
