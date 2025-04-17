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
