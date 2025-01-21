package shared

import (
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers"
)

type generatingAttributeGenerator struct {
	Payload Attributes
}

var _ AttributesGenerator = (*generatingAttributeGenerator)(nil)

func CreateAttributesGeneratorMock(
	payload Attributes,
) AttributesGenerator {
	return &generatingAttributeGenerator{
		Payload: payload,
	}
}

// Generate implements shared.AttributesGenerator.
func (g *generatingAttributeGenerator) Generate() AttributesChannel {
	ch := make(AttributesChannel)
	go func() {
		ch <- g.Payload
		close(ch)
	}()
	return ch
}

type providerMock[T any] struct {
	providedObject T
}

func CreateProviderMock[T any](d T) providers.Provider[T] {
	return &providerMock[T]{
		providedObject: d,
	}
}

// Provide implements domain.DomainProvider.
func (m *providerMock[T]) Provide() <-chan T {
	ch := make(chan T)
	go func() {
		ch <- m.providedObject
		close(ch)
	}()

	return ch
}

type emptyProviderMock[T any] struct{}

func CreateEmptyProviderMock[T any]() providers.Provider[T] {
	return &emptyProviderMock[T]{}
}

// Provide implements domain.DomainProvider.
func (*emptyProviderMock[T]) Provide() <-chan T {
	ch := make(chan T)
	defer close(ch)
	return ch
}
