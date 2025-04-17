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

type Entities struct {
	entities map[string]Entity
}

type Entity struct {
	entityType string
	ids        []string
	attributes []string
}

func (e *Entities) GetEntity(entityType string) Entity {
	return e.entities[entityType]
}

func NewEntities(entities map[string]Entity) *Entities {
	return &Entities{
		entities,
	}
}

func NewEntity(entityType string, ids []string, attributes []string) Entity {
	return Entity{
		entityType: entityType,
		ids:        ids,
		attributes: attributes,
	}
}

func (e Entity) Type() string {
	return e.entityType
}

func (e Entity) IDs() []string {
	return e.ids
}

func (e Entity) Attributes() []string {
	return e.attributes
}
