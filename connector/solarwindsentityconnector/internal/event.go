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
	"time"
)

const (
	entityUpdateEventType = "entity_state"
)

type Entity struct {
	Type       string   `mapstructure:"entity"`
	IDs        []string `mapstructure:"id"`
	Attributes []string `mapstructure:"attributes"`
}

func AppendEntityUpdateEvent(lrs *plog.LogRecordSlice, entity Entity, resourceAttrs pcommon.Map) {
	lr := plog.NewLogRecord()

	if exists := setIdAttributes(&lr, entity.IDs, resourceAttrs); !exists {
		return
	}
	attributes := lr.Attributes()
	// event type
	attributes.PutStr(swoEntityEventType, entityUpdateEventType)
	// entity type
	attributes.PutStr(swoEntityType, entity.Type)
	// entity attributes
	logIds := attributes.PutEmptyMap(swoEntityAttributes)
	for _, id := range entity.Attributes {
		copyAttribute(&logIds, id, &resourceAttrs)
	}
	// timestamp
	lr.SetObservedTimestamp(pcommon.NewTimestampFromTime(time.Now()))

	eventLog := lrs.AppendEmpty()
	lr.CopyTo(eventLog)
}
