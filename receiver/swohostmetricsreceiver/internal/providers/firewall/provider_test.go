package firewall

import (
	"fmt"
	"testing"
)

func Test_Functional(t *testing.T) {
	t.Skip("This test should be run manually")

	sut := CreateFirewallProvider()
	ch := sut.Provide()
	result := <-ch
	fmt.Printf("Result: %+v", result)
}
