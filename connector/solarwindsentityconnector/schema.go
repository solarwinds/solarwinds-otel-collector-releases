package solarwindsentityconnector

import "github.com/solarwinds/solarwinds-otel-collector-releases/connector/solarwindsentityconnector/internal"

type Schema struct {
	Entities []Entity `mapstructure:"entities"`
}

type Entity struct {
	Type       string      `mapstructure:"entity"`
	ID         []Attribute `mapstructure:"id"`
	Attributes []Attribute `mapstructure:"attributes"`
}

type Attribute struct {
	ResourceAttribute string `mapstructure:"resource_attribute"`
	Property          string `mapstructure:"property"`
}

func (s *Schema) NewEntities() *internal.Entities {
	entities := make(map[string]internal.Entity)
	for _, entity := range s.Entities {
		entities[entity.Type] = internal.NewEntity(
			entity.Type,
			convert(entity.ID),
			convert(entity.Attributes))
	}

	return internal.NewEntities(entities)
}

func convert(schemaAttributes []Attribute) []internal.Attribute {
	attrs := make([]internal.Attribute, len(schemaAttributes))
	for _, attr := range schemaAttributes {
		attrs = append(attrs, internal.NewAttribute(attr.ResourceAttribute, attr.Property))
	}

	return attrs
}
