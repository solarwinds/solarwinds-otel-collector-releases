package types

type AttributesProvider interface {
	// Provides map of attributes.
	GetAttributes() (map[string]any, error)
}
