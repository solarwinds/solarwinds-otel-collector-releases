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

package scope

import (
	"github.com/solarwinds/solarwinds-otel-collector-releases/receiver/swohostmetricsreceiver/internal/scraper/framework/metric"
	"go.uber.org/zap"
)

// ScopeEmitterCreateFunc is a functor for creation of
// scope emitter instance.
// string determines scope name for which emitter is created.
type EmitterCreateFunc func(string, map[string]metric.Emitter) Emitter

// Descriptor for scope emitter. It is used for declarative description of
// scope emitter.
type Descriptor struct {
	// ScopeName is name of scope for described scope emitter.
	ScopeName string
	// Map of metrics descriptors for this scope. Map keys represent names
	// metrics.
	MetricDescriptors map[string]metric.Descriptor
	// Overrideable creator for custom scope creator. In case there is none
	// it is supposed to be replaced by generic one.
	Create EmitterCreateFunc
}

func TraverseThroughScopeDescriptors(
	scopeDescriptors map[string]Descriptor,
	enabledMetrics *metric.Enabled,
) map[string]Emitter {
	ses := make(map[string]Emitter, 0)

	for sName, sDescriptor := range scopeDescriptors {
		// Traverse metric descriptors for given scope descriptor.
		mes := metric.TraverseThroughMetricDescriptors(
			sDescriptor.MetricDescriptors,
			enabledMetrics,
		)

		// Given scope was not configured.
		if len(mes) == 0 {
			continue
		}

		// Choose allocator - custom or default.
		create := chooseEmitterAllocator(&sDescriptor)

		// Creates scope emitter with proper setup for given metric emitters.
		se := create(sDescriptor.ScopeName, mes)

		zap.L().Sugar().Debugf(
			"creation of scope emitter for scope '%s' was finished successfully",
			sName,
		)
		ses[se.Name()] = se
	}

	return ses
}

func chooseEmitterAllocator(
	descriptor *Descriptor,
) EmitterCreateFunc {
	var createEmitter EmitterCreateFunc

	if descriptor.Create != nil {
		zap.L().Sugar().Debugf(
			"custom scope allocator will be used for scope '%s'",
			descriptor.ScopeName,
		)
		createEmitter = descriptor.Create
	} else {
		zap.L().Sugar().Debugf(
			"default scope allocator will be used for scope '%s'",
			descriptor.ScopeName,
		)
		createEmitter = CreateDefaultScopeEmitter
	}
	return createEmitter
}
