package internal

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"time"
)

const (
	entityUpdate = "entity_state"

	// Attributes for OTel entity events identification
	otelEntityEventAsLog = "otel.entity.event_as_log"
	otelEntityEventType  = "otel.entity.event.type"
	swoEntityType        = "otel.entity.type"
	swoEntityId          = "otel.entity.id"
	swoEntityAttributes  = "otel.entity.attributes"
)

// BuildEventLog prepares a clean LogRecordSlice, where log records representing events should be appended.
// In give plog.Logs creates a resource log with one scope log and set all attributes needed for SWO ingestion.
func BuildEventLog(logs plog.Logs) plog.LogRecordSlice {
	resourceLog := logs.ResourceLogs().AppendEmpty()
	resourceLog.Resource().Attributes().PutStr(otelEntityEventAsLog, "true")
	scopeLog := resourceLog.ScopeLogs().AppendEmpty()
	lrs := scopeLog.LogRecords()

	return lrs
}

func AppendEvent(lrs *plog.LogRecordSlice, entity Entity, telemetryAttributes pcommon.Map) {
	lr := lrs.AppendEmpty()
	lr.SetObservedTimestamp(pcommon.NewTimestampFromTime(time.Now()))
	attrs := lr.Attributes()

	attrs.PutStr(otelEntityEventType, entityUpdate)
	attrs.PutStr(swoEntityType, entity.entityType)

	// Identification attributes
	ids := attrs.PutEmptyMap(swoEntityId)
	for _, id := range entity.ids {
		ids.PutStr(id.ResourceAttribute(), getValue(id.ResourceAttribute(), telemetryAttributes))
	}

	// Entity attributes
	ea := attrs.PutEmptyMap(swoEntityAttributes)
	for _, a := range entity.attributes {
		ea.PutStr(a.ResourceAttribute(), getValue(a.ResourceAttribute(), telemetryAttributes))
	}

}

func getValue(key string, attributes pcommon.Map) string {
	value, exists := attributes.Get(key)
	if exists {
		return value.AsString()
	}

	return "TEST_JT"
}
