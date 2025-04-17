package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	expectedEntities = map[string]Entity{
		"testEntityType":  NewEntity("testEntityType", []string{"id1", "id2"}, []string{"attr1", "attr2"}),
		"testEntityType2": NewEntity("testEntityType2", []string{"id3"}, []string{"attr3"}),
	}
)

func TestGetEntity(t *testing.T) {
	// arrange
	entities := NewEntities(expectedEntities)

	// act
	actualEntity := entities.GetEntity("testEntityType")

	// assert
	assert.NotNil(t, actualEntity)
	assert.Equal(t, expectedEntities["testEntityType"], actualEntity)
}
