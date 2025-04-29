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
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
)

const (
	swoEntityEventAsLog = "otel.entity.event_as_log"
	swoEntityEventType  = "otel.entity.event.type"
	swoEntityType       = "otel.entity.type"
	swoEntityIds        = "otel.entity.id"
	swoEntityAttributes = "otel.entity.attributes"
)

// BuildEventLog prepares a clean LogRecordSlice, where log records representing events should be appended.
// In given plog.Logs creates a resource log with one scope log and set attributes needed by SWO ingestion.
func BuildEventLog(logs *plog.Logs) *plog.LogRecordSlice {
	resourceLog := logs.ResourceLogs().AppendEmpty()
	scopeLog := resourceLog.ScopeLogs().AppendEmpty()
	scopeLog.Scope().Attributes().PutBool(swoEntityEventAsLog, true)
	lrs := scopeLog.LogRecords()

	return &lrs
}

// setIdAttributes sets the entity id attributes in the log record as needed by SWO.
// Attributes are used to infer the entity in the system.
//
// Returns false if any of the attributes are missing in the resourceAttrs.
// If any ID attribute is missing the entity would not be inferred.
// Returns true if all attributes were set.
func setIdAttributes(lr *plog.LogRecord, entityIds []string, resourceAttrs pcommon.Map) bool {
	attrs := lr.Attributes()
	logIds := attrs.PutEmptyMap(swoEntityIds)
	for _, id := range entityIds {
		if !copyAttribute(&logIds, id, &resourceAttrs) {
			zap.L().Warn("failed to put entity id", zap.String("key", id))
			return false
		}
	}
	return true
}

// copyAttribute copies the value of attribute identified as key, from source to dest pcommon.Map.
// It returns true if the attribute was found and copied, false otherwise.
func copyAttribute(dest *pcommon.Map, key string, src *pcommon.Map) bool {
	attrVal, ok := src.Get(key)

	if !ok {
		zap.L().Warn("attribute not found", zap.String("key", key))
		return false
	}

	switch typeAttr := attrVal.Type(); typeAttr {
	case pcommon.ValueTypeInt:
		dest.PutInt(key, attrVal.Int())
	case pcommon.ValueTypeDouble:
		dest.PutDouble(key, attrVal.Double())
	case pcommon.ValueTypeBool:
		dest.PutBool(key, attrVal.Bool())
	case pcommon.ValueTypeBytes:
		dest.PutEmptyBytes(attrVal.Str())
	default:
		dest.PutStr(key, attrVal.Str())
	}

	return true
}
