package providers

// Provider is a general interface for most of the data providers.
type Provider[T any] interface {
	// Provide returns a channel through which the data are provided.
	Provide() <-chan T
}
