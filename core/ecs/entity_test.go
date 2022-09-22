package ecs

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleEntity(t *testing.T) {
	entity := NewSimpleEntity(10)

	assert.Equal(t, 10, entity.Id())
	assert.Equal(t, &ComponentCollection{components: map[reflect.Type]Component{}}, entity.Components())
}
