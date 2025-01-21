package domain

import (
	"fmt"
	"testing"
)

func Test_Functional(t *testing.T) {
	t.Skip("This test should be run manually")

	sut := CreateDomainProvider()
	result := <-sut.Provide()
	fmt.Printf("Result: %+v\n", result)
}
