package firewall

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Provide_ProvidesEmptyFirewallProfileCollectionWithNoErrors(t *testing.T) {
	sut := CreateFirewallProvider()
	ch := sut.Provide()
	actualModel := <-ch
	_, open := <-ch

	assert.Nil(t, actualModel.FirewallProfiles, "unsupported provider must return no data")
	assert.Nil(t, actualModel.Error, "unsupported provider must return no error")
	assert.False(t, open, "channel must be closed afterward")
}
