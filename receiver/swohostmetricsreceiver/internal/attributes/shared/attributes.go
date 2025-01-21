package shared

// Abstraction over attributes.
type Attributes map[string]string

// Attributes channel represent channel for sending attributes
// when they are generated.
type AttributesChannel chan Attributes
