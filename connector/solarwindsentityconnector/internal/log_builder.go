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

// buildEventLog prepares a clean LogRecordSlice, where log records representing events should be appended.
// In given plog.Logs creates a resource log with one scope log and set attributes needed by SWO ingestion.
func buildEventLog(logs *plog.Logs) *plog.LogRecordSlice {
	resourceLog := logs.ResourceLogs().AppendEmpty()
	scopeLog := resourceLog.ScopeLogs().AppendEmpty()
	scopeLog.Scope().Attributes().PutBool(swoEntityEventAsLog, true)
	lrs := scopeLog.LogRecords()

	return &lrs
}

// setEntityType sets the entity type attribute in the log record needed by SWO ingestion.
func setEntityType(lr *plog.LogRecord, entityType string) {
	attrs := lr.Attributes()
	attrs.PutStr(swoEntityType, entityType)
}

// setEventType sets the event type attribute in the log record for SWO to recognized
// what kind of event it is.
func setEventType(lr *plog.LogRecord, eventType string) {
	attrs := lr.Attributes()
	attrs.PutStr(swoEntityEventType, eventType)
}

// setIdAttributes sets the entity id attributes in the log record needed by SWO,
// which are used to identify or infer the entity in the system.
func setIdAttributes(lr *plog.LogRecord, entityIds []string, resourceAttrs pcommon.Map) bool {
	attrs := lr.Attributes()
	logIds := attrs.PutEmptyMap(swoEntityIds)
	for _, id := range entityIds {
		// If identification attribute is not found, entity will not be inferred
		if !putAttribute(&logIds, id, &resourceAttrs) {
			zap.L().Warn("failed to put entity id", zap.String("key", id))
			return false
		}
	}

	return true
}

// setAttributes sets the entity attributes in the log record used for updating state of a SWO entity.
func setAttributes(lr *plog.LogRecord, entityAttrs []string, resourceAttrs pcommon.Map) {
	attrs := lr.Attributes()
	logIds := attrs.PutEmptyMap(swoEntityAttributes)
	for _, id := range entityAttrs {
		putAttribute(&logIds, id, &resourceAttrs)
	}
}

// putAttribute copies the value of attribute identified as key, from source to dest pcommon.Map.
// It returns true if the attribute was found and copied, false otherwise.
func putAttribute(dest *pcommon.Map, key string, src *pcommon.Map) bool {
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
