package solarwindsentityconnector

import "github.com/solarwinds/solarwinds-otel-collector-releases/connector/solarwindsentityconnector/internal"

type Schema struct {
	Entities []Entity `mapstructure:"entities"`
}

type Entity struct {
	Type       string   `mapstructure:"entity"`
	IDs        []string `mapstructure:"id"`
	Attributes []string `mapstructure:"attributes"`
}

func (s *Schema) NewEntities() *internal.Entities {
	entities := make(map[string]internal.Entity)
	for _, entity := range s.Entities {
		entities[entity.Type] = internal.NewEntity(
			entity.Type,
			entity.IDs,
			entity.Attributes)
	}

	return internal.NewEntities(entities)
}
