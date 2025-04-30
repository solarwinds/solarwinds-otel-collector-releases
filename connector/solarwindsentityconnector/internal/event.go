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
	swoRelationshipAttributes   = "otel.entity_relationship.attributes"
)

type Entity struct {
	Type       string   `mapstructure:"entity"`
	IDs        []string `mapstructure:"id"`
	Attributes []string `mapstructure:"attributes"`
}

type Relationship struct {
	Type        string   `mapstructure:"type"`
	Source      string   `mapstructure:"source_entity"`
	Destination string   `mapstructure:"destination_entity"`
	Attributes  []string `mapstructure:"attributes"`
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
	setAttributes(attrs, entity.Attributes, resourceAttrs, swoEntityIds)
	lr.SetObservedTimestamp(pcommon.NewTimestampFromTime(time.Now()))

	eventLog := lrs.AppendEmpty()
	lr.CopyTo(eventLog)
}

func AppendRelationshipUpdateEvent(lrs *plog.LogRecordSlice, relationship Relationship, resourceAttrs pcommon.Map, entities map[string]Entity) {
	lr := plog.NewLogRecord()
	attrs := lr.Attributes()

	src := entities[relationship.Source]
	dest := entities[relationship.Destination]
	err := setSourceEntityProperties(attrs, src, resourceAttrs)
	if err != nil {
		return
	}
	err = setDestinationEntityProperties(attrs, dest, resourceAttrs)
	if err != nil {
		return
	}

	// set destination attributes
	setEventType(attrs, relationshipUpdateEventType)
	setAttributes(attrs, []string{"placeholderAttr1"}, resourceAttrs, swoRelationshipAttributes)
	attrs.PutStr(swoRelationshipType, "placeholder")
	lr.SetObservedTimestamp(pcommon.NewTimestampFromTime(time.Now()))

	eventLog := lrs.AppendEmpty()
	lr.CopyTo(eventLog)
}
