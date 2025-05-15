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
	"fmt"
	"time"

	"github.com/solarwinds/solarwinds-otel-collector-releases/connector/solarwindsentityconnector/config"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
)

type EventBuilder struct {
	Entities       map[string]config.Entity
	Relationships  []config.Relationship
	SourcePrefix   string
	DestPrefix     string
	ResultEventLog *plog.LogRecordSlice
	Logger         *zap.Logger
}

func NewEventBuilder(entities map[string]config.Entity, relationships []config.Relationship, sourcePrefix string, destPrefix string, events *plog.Logs, logger *zap.Logger) *EventBuilder {
	return &EventBuilder{
		Entities:       entities,
		Relationships:  relationships,
		SourcePrefix:   sourcePrefix,
		DestPrefix:     destPrefix,
		ResultEventLog: CreateResultEventLog(events),
		Logger:         logger,
	}
}

// CreateResultEventLog prepares a clean LogRecordSlice, where log records representing events should be appended.
// In given plog.Logs creates a resource log with one scope log and set attributes needed by SWO ingestion.
func CreateResultEventLog(logs *plog.Logs) *plog.LogRecordSlice {
	resourceLog := logs.ResourceLogs().AppendEmpty()
	scopeLog := resourceLog.ScopeLogs().AppendEmpty()
	scopeLog.Scope().Attributes().PutBool(entityEventAsLog, true)
	lrs := scopeLog.LogRecords()

	return &lrs
}

func (e *EventBuilder) AppendEntityUpdateEvent(entity config.Entity, resourceAttrs pcommon.Map) {

	event, err := e.createEntityEvent(resourceAttrs, entity)
	if err != nil {
		e.Logger.Debug("failed to create update event", zap.Error(err))
		return
	}

	event.SetObservedTimestamp(pcommon.NewTimestampFromTime(time.Now()))
	event.Attributes().PutStr(entityEventType, entityUpdateEventType)
	eventLog := e.ResultEventLog.AppendEmpty()
	event.CopyTo(eventLog)
}

func (e *EventBuilder) createEntityEvent(resourceAttrs pcommon.Map, entity config.Entity) (plog.LogRecord, error) {
	lr := plog.NewLogRecord()
	attrs := lr.Attributes()
	attrs.PutStr(entityType, entity.Type)

	if err := setIdAttributesDefault(attrs, entity.IDs, resourceAttrs, entityIds); err != nil {
		return plog.LogRecord{}, fmt.Errorf("failed to set id attributes: %w", err)
	}

	setAttributes(attrs, entity.Attributes, resourceAttrs, entityAttributes)

	lr.SetObservedTimestamp(pcommon.NewTimestampFromTime(time.Now()))

	return lr, nil
}

func (e *EventBuilder) AppendRelationshipUpdateEvent(relationship config.Relationship, resourceAttrs pcommon.Map) {
	relationshipLog, err := e.createRelationshipEvent(relationship, resourceAttrs)
	if err != nil {
		e.Logger.Debug("Failed to create relationship event", zap.Error(err))
		return
	}

	relationshipLog.SetObservedTimestamp(pcommon.NewTimestampFromTime(time.Now()))
	relationshipLog.Attributes().PutStr(entityEventType, relationshipUpdateEventType)
	eventLog := e.ResultEventLog.AppendEmpty()
	relationshipLog.CopyTo(eventLog)

}

func (e *EventBuilder) createRelationshipEvent(relationship config.Relationship, resourceAttrs pcommon.Map) (plog.LogRecord, error) {
	source, ok := e.Entities[relationship.Source]
	if !ok {
		return plog.NewLogRecord(), fmt.Errorf("bad source entity")
	}

	dest, ok := e.Entities[relationship.Destination]
	if !ok {
		return plog.NewLogRecord(), fmt.Errorf("bad destination entity")
	}

	lr := plog.NewLogRecord()
	attrs := lr.Attributes()

	if source.Type == dest.Type {
		// same type relationships
		hasPrefixSrc, err := setIdAttributesSameType(attrs, source.IDs, resourceAttrs, relationshipSrcEntityIds, e.SourcePrefix)
		if err != nil {
			return plog.NewLogRecord(), fmt.Errorf("bad source entity id attributes")
		}

		hasPrefixDst, err := setIdAttributesSameType(attrs, dest.IDs, resourceAttrs, relationshipDestEntityIds, e.DestPrefix)
		if err != nil {
			return plog.NewLogRecord(), fmt.Errorf("bad destination entity id attributes")
		}

		if !hasPrefixSrc || !hasPrefixDst {
			return plog.NewLogRecord(), fmt.Errorf("same type entity ids are not prefixed")
		}

	} else {

		if err := setIdAttributesDefault(attrs, source.IDs, resourceAttrs, relationshipSrcEntityIds); err != nil {
			return plog.NewLogRecord(), fmt.Errorf("bad source entity id attributes")
		}

		if err := setIdAttributesDefault(attrs, dest.IDs, resourceAttrs, relationshipDestEntityIds); err != nil {
			return plog.NewLogRecord(), fmt.Errorf("bad destination entity id attributes")
		}
	}

	setAttributes(attrs, relationship.Attributes, resourceAttrs, relationshipAttributes)
	attrs.PutStr(relationshipType, relationship.Type)
	attrs.PutStr(srcEntityType, source.Type)
	attrs.PutStr(destEntityType, dest.Type)

	return lr, nil
}
