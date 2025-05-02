package internal

const (
	// Log properties
	entityEventAsLog = "otel.entity.event_as_log"
	entityEventType  = "otel.entity.event.type"

	// Event type values
	entityUpdateEventType       = "entity_state"
	relationshipUpdateEventType = "entity_relationship_state"

	// Entity properties
	entityType       = "otel.entity.type"
	entityIds        = "otel.entity.id"
	entityAttributes = "otel.entity.attributes"

	// Relationship properties
	relationshipSrcEntityIds  = "otel.entity_relationship.source_entity.id"
	relationshipDestEntityIds = "otel.entity_relationship.destination_entity.id"
	relationshipAttributes    = "otel.entity_relationship.attributes"
	relationshipType          = "otel.entity_relationship.type"
	srcEntityType             = "otel.entity_relationship.source_entity.type"
	destEntityType            = "otel.entity_relationship.destination_entity.type"
)
