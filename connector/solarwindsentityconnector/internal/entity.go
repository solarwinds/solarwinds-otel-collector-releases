package internal

type Entities struct {
	entities map[string]Entity
}

type Entity struct {
	entityType string
	ids        []Attribute
	attributes []Attribute
}

type Attribute struct {
	resourceAttribute string
	property          string
}

func (e *Entities) GetEntity(entityType string) Entity {
	return e.entities[entityType]
}

func NewEntities(entities map[string]Entity) *Entities {
	return &Entities{
		entities,
	}
}

func NewEntity(entityType string, ids []Attribute, attributes []Attribute) Entity {
	return Entity{
		entityType: entityType,
		ids:        ids,
		attributes: attributes,
	}
}

func NewAttribute(resourceAttribute string, property string) Attribute {
	return Attribute{
		resourceAttribute: resourceAttribute,
		property:          property,
	}
}

func (e Entity) Type() string {
	return e.entityType
}

func (e Entity) IDs() []Attribute {
	return e.ids
}

func (e Entity) Attributes() []Attribute {
	return e.attributes
}

func (a Attribute) ResourceAttribute() string {
	return a.resourceAttribute
}

func (a Attribute) Property() string {
	return a.property
}
