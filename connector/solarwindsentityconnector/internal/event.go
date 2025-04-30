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
	entityUpdateEventType = "entity_state"
)

type Entity struct {
	Type       string   `mapstructure:"entity"`
	IDs        []string `mapstructure:"id"`
	Attributes []string `mapstructure:"attributes"`
}

func AppendEntityUpdateEvent(lrs *plog.LogRecordSlice, entity Entity, resourceAttrs pcommon.Map) {
	lr := plog.NewLogRecord()
	attrs := lr.Attributes()

	err := setIdAttributes(attrs, entity.IDs, resourceAttrs)
	if err != nil {
		zap.L().Debug("failed to create update event", zap.Error(err))
		return
	}

	setEventType(attrs, entityUpdateEventType)
	setEntityType(attrs, entity.Type)
	setEntityAttributes(attrs, entity.Attributes, resourceAttrs)
	lr.SetObservedTimestamp(pcommon.NewTimestampFromTime(time.Now()))

	eventLog := lrs.AppendEmpty()
	lr.CopyTo(eventLog)
}
