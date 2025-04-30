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

package solarwindsentityconnector

import "github.com/solarwinds/solarwinds-otel-collector-releases/connector/solarwindsentityconnector/internal"

type Schema struct {
	Entities []internal.Entity `mapstructure:"entities"`
}

func (s *Schema) NewEntities() map[string]internal.Entity {
	entities := make(map[string]internal.Entity, len(s.Entities))
	for _, entity := range s.Entities {
		entities[entity.Type] = internal.Entity{
			Type:       entity.Type,
			IDs:        entity.IDs,
			Attributes: entity.Attributes}
	}

	return entities
}
