package ecs

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComponentCollection_AddWhenComponentNotExists(t *testing.T) {
	componentCollection := createComponentCollection()
	barComponent := barComponent{}

	componentCollection.add(barComponent)
}

func TestComponentCollection_AddWhenComponentExists(t *testing.T) {
	componentCollection := createComponentCollection()
	barComponent := barComponent{}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	componentCollection.add(barComponent)
	componentCollection.add(barComponent)
}

func TestComponentCollection_HasWhenComponentExists(t *testing.T) {
	componentCollection := createComponentCollection()
	assert.True(t, componentCollection.Has(fooComponent{}))
}

func TestComponentCollection_HasWhenComponentNotExists(t *testing.T) {
	componentCollection := createComponentCollection()
	assert.False(t, componentCollection.Has(barComponent{}))
}

func TestComponentCollection_HasAll(t *testing.T) {
	componentCollection := &ComponentCollection{
		components: map[reflect.Type]Component{
			fooComponent{}.TypeOf(): fooComponent{},
			barComponent{}.TypeOf(): barComponent{},
		},
	}

	assert.True(t, componentCollection.HasAll(fooComponent{}))
	assert.True(t, componentCollection.HasAll(fooComponent{}, barComponent{}))
	assert.False(t, componentCollection.HasAll(fooComponent{}, barComponent{}, buzComponent{}))
}

func TestComponentCollection_GetWhenComponentExists(t *testing.T) {
	componentCollection := createComponentCollection()
	assert.Equal(t, fooComponent{}, componentCollection.GetOrPanic(fooComponent{}))
}

func TestComponentCollection_GetWhenComponentNotExists(t *testing.T) {
	componentCollection := createComponentCollection()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	componentCollection.GetOrPanic(barComponent{})
}

func TestComponentCollection_Remove(t *testing.T) {
	componentCollection := createComponentCollection()

	componentCollection.remove(fooComponent{})

	expected := &ComponentCollection{
		components: map[reflect.Type]Component{},
	}

	assert.Equal(t, expected, componentCollection)

	componentCollection.remove(fooComponent{})

	assert.Equal(t, expected, componentCollection)
}

func createComponentCollection() *ComponentCollection {
	fooComponent := fooComponent{}

	return &ComponentCollection{
		components: map[reflect.Type]Component{
			fooComponent.TypeOf(): fooComponent,
		},
	}
}

type fooComponent struct {
}

func (c fooComponent) TypeOf() reflect.Type {
	return reflect.TypeOf(c)
}

type barComponent struct {
}

func (b barComponent) TypeOf() reflect.Type {
	return reflect.TypeOf(b)
}

type buzComponent struct {
}

func (b buzComponent) TypeOf() reflect.Type {
	return reflect.TypeOf(b)
}
