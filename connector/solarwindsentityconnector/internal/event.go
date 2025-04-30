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
	"time"
)

const (
	entityUpdateEventType       = "entity_state"
	relationshipUpdateEventType = "entity_relationship_state"
	swoEntityType               = "otel.entity.type"
	swoEntityIds                = "otel.entity.id"
	swoSourceEntityIds          = "otel.entity_relationship.source_entity.id"
	swoDestinationEntityIds     = "otel.entity_relationship.destination_entity.id"
	swoRelationshipType         = "otel.entity_relationship.type"
	swoSourceEntityType         = "otel.entity_relationship.source_entity.type"
	swoDestinationEntityType    = "otel.entity_relationship.destination_entity.type"
)

type Entity struct {
	Type       string   `mapstructure:"entity"`
	IDs        []string `mapstructure:"id"`
	Attributes []string `mapstructure:"attributes"`
}

func AppendEntityUpdateEvent(lrs *plog.LogRecordSlice, entity Entity, resourceAttrs pcommon.Map) {
	lr := plog.NewLogRecord()
	attrs := lr.Attributes()

	err := setIdAttributes(attrs, entity.IDs, resourceAttrs, swoEntityIds)
	if err != nil {
		zap.L().Debug("failed to create update event", zap.Error(err))
		return
	}

	setEventType(attrs, entityUpdateEventType)
	setEntityType(attrs, entity.Type, swoEntityType)
	setEntityAttributes(attrs, entity.Attributes, resourceAttrs)
	lr.SetObservedTimestamp(pcommon.NewTimestampFromTime(time.Now()))

	eventLog := lrs.AppendEmpty()
	lr.CopyTo(eventLog)
}

func AppendRelationshipUpdateEvent(lrs *plog.LogRecordSlice, src Entity, dest Entity, resourceAttrs pcommon.Map) {
	lr := plog.NewLogRecord()
	attrs := lr.Attributes()

	// set id attributes of both source and destination entities
	err := setIdAttributes(attrs, src.IDs, resourceAttrs, swoSourceEntityIds)
	if err != nil {
		zap.L().Debug("failed to create update event", zap.Error(err))
		return
	}

	err = setIdAttributes(attrs, dest.IDs, resourceAttrs, swoDestinationEntityIds)
	if err != nil {
		zap.L().Debug("failed to create update event", zap.Error(err))
		return
	}

	// set destination attributes
	setEventType(attrs, relationshipUpdateEventType)
	setEntityType(attrs, src.Type, swoSourceEntityType)
	setEntityType(attrs, dest.Type, swoDestinationEntityType)
	lr.SetObservedTimestamp(pcommon.NewTimestampFromTime(time.Now()))

	eventLog := lrs.AppendEmpty()
	lr.CopyTo(eventLog)
}
