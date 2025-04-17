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

type eventsHandler interface {
	AppendUpdateEvent(entity Entity, telemetryAttributes pcommon.Map)
}

type Events struct {
	logRecords *plog.LogRecordSlice
}

func NewEvents(logs plog.Logs) *Events {
	return &Events{
		logRecords: buildEventLog(&logs),
	}
}

func (e *Events) AppendUpdateEvent(entity Entity, resourceAttrs pcommon.Map) {
	lr := plog.NewLogRecord()
	lr.SetObservedTimestamp(pcommon.NewTimestampFromTime(time.Now()))

	setEventType(&lr, entityUpdateEventType)
	setEntityType(&lr, entity.Type())
	setIdAttributes(&lr, entity.IDs(), resourceAttrs)
	setAttributes(&lr, entity.IDs(), resourceAttrs)

}

var _ eventsHandler = (*Events)(nil)
