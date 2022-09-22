package ecs

import (
	"fmt"
	"reflect"
)

// Component is a metadata container used to house information about something an entities state.
// Example Components might look like:
// type PositionComponent struct {
//     X int
//     Y int
// }
// This position component represents where an entity is in the game world. If an entity does not have a position
// component, it can be assumed they are not present in the world.
// Another type of component might look like:
// type CanAttackComponent {}
// CanAttackComponent has no data attached, and acts merely as a flag. If an entity has this component, they can attack
// if an entity is missing this component, they cannot attack.
// Components are a flexible way of attaching metadata to an entity.
type Component interface {
	TypeOf() reflect.Type
}

type ComponentCollection struct {
	components map[reflect.Type]Component
}

func NewEmptyComponentCollection() *ComponentCollection {
	return &ComponentCollection{components: map[reflect.Type]Component{}}
}

func (collection *ComponentCollection) Has(component Component) bool {
	_, ok := collection.components[component.TypeOf()]

	return ok
}

func (collection *ComponentCollection) HasAll(components ...Component) bool {
	for _, component := range components {
		if !collection.Has(component) {
			return false
		}
	}

	return true
}

func (collection *ComponentCollection) GetOrPanic(target Component) Component {
	component, err := collection.Get(target)

	if err != nil {
		panic(err)
	}

	return component
}

func (collection *ComponentCollection) Get(component Component) (Component, error) {
	if component, ok := collection.components[component.TypeOf()]; ok == true {
		return component, nil
	}

	return nil, fmt.Errorf("value [%s] not exist in collection", component.TypeOf().Name())
}

func (collection *ComponentCollection) All() map[reflect.Type]Component {
	return collection.components
}

func (collection *ComponentCollection) add(component Component) error {
	if collection.Has(component) == true {
		return fmt.Errorf("component [%s] already exists in collection", component.TypeOf().Name())
	}

	collection.components[component.TypeOf()] = component

	return nil
}

func (collection *ComponentCollection) remove(component Component) error {
	if !collection.Has(component) {
		return fmt.Errorf("Try to remove component [%s] that not exists in collection.", component.TypeOf())
	}

	delete(collection.components, component.TypeOf())

	return nil
}
