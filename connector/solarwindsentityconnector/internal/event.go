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
	"github.com/solarwinds/solarwinds-otel-collector-releases/connector/solarwindsentityconnector/config"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
	"time"
)

type EventBuilder struct {
	Entities      map[string]config.Entity
	Relationships []config.Relationship
	SourcePrefix  string
	DestPrefix    string
	Result        *plog.LogRecordSlice
	Logger        *zap.Logger
}

func NewEventBuilder(entities map[string]config.Entity, relationships []config.Relationship, sourcePrefix string, destPrefix string, events *plog.Logs, logger *zap.Logger) *EventBuilder {
	return &EventBuilder{
		Entities:      entities,
		Relationships: relationships,
		SourcePrefix:  sourcePrefix,
		DestPrefix:    destPrefix,
		Result:        BuildEventLog(events),
		Logger:        logger,
	}
}

func (e *EventBuilder) AppendEntityUpdateEvent(entity config.Entity, resourceAttrs pcommon.Map) {

	event, err := e.createEntityEvent(resourceAttrs, entity)
	if err != nil {
		zap.L().Debug("failed to create update event", zap.Error(err))
		return
	}
	event.SetObservedTimestamp(pcommon.NewTimestampFromTime(time.Now()))
	event.Attributes().PutStr(entityEventType, entityUpdateEventType)
	eventLog := e.Result.AppendEmpty()
	event.CopyTo(eventLog)
}

func (e *EventBuilder) createEntityEvent(resourceAttrs pcommon.Map, entity config.Entity) (plog.LogRecord, error) {
	lr := plog.NewLogRecord()
	attrs := lr.Attributes()
	attrs.PutStr(entityType, entity.Type)

	if err := setIdAttributes(attrs, entity.IDs, resourceAttrs, entityIds); err != nil {
		return plog.LogRecord{}, err
	}

	setAttributes(attrs, entity.Attributes, resourceAttrs, entityAttributes)

	lr.SetObservedTimestamp(pcommon.NewTimestampFromTime(time.Now()))

	return lr, nil
}

func (e *EventBuilder) AppendRelationshipUpdateEvent(relationship config.Relationship, resourceAttrs pcommon.Map) {
	relationshipLog, err := e.createRelationship(relationship, resourceAttrs)
	if err != nil {
		zap.L().Debug("failed to create relationship event", zap.Error(err))
		return
	}
	relationshipLog.SetObservedTimestamp(pcommon.NewTimestampFromTime(time.Now()))
	relationshipLog.Attributes().PutStr(entityEventType, relationshipUpdateEventType)
	eventLog := e.Result.AppendEmpty()
	relationshipLog.CopyTo(eventLog)

}

func (e *EventBuilder) createRelationship(relationship config.Relationship, resourceAttrs pcommon.Map) (plog.LogRecord, error) {
	source, ok := e.Entities[relationship.Source]
	if !ok {
		zap.L().Debug("source entity not found", zap.String("entity", relationship.Source))
		return plog.NewLogRecord(), nil
	}

	dest, ok := e.Entities[relationship.Destination]
	if !ok {
		zap.L().Debug("destination entity not found", zap.String("entity", relationship.Destination))
		return plog.NewLogRecord(), nil
	}

	lr := plog.NewLogRecord()
	attrs := lr.Attributes()
	if err := setIdAttributes2(attrs, source.IDs, resourceAttrs, relationshipSrcEntityIds, e.SourcePrefix, e.Logger); err != nil {
		return plog.NewLogRecord(), nil
	}

	if err := setIdAttributes2(attrs, dest.IDs, resourceAttrs, relationshipDestEntityIds, e.DestPrefix, e.Logger); err != nil {
		return plog.NewLogRecord(), nil
	}

	setAttributes(attrs, relationship.Attributes, resourceAttrs, relationshipAttributes)
	attrs.PutStr(relationshipType, relationship.Type)
	attrs.PutStr(srcEntityType, source.Type)
	attrs.PutStr(destEntityType, dest.Type)

	return lr, nil
}
