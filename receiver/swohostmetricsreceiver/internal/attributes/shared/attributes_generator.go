package shared

// Represents general prescription of attributes generator.
type AttributesGenerator interface {
	// Generates attributes through channel. When it is done, channel is closed.
	Generate() AttributesChannel
}
