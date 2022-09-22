package ecs

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createEntityController() EntityController {
	return EntityController{
		nextEntityID: 2,
		components: map[reflect.Type][]int{
			fooComponent{}.TypeOf(): {0, 1},
			barComponent{}.TypeOf(): {0, 1},
		},
		entities: map[int]Entity{
			0: &SimpleEntity{
				id: 0,
				components: &ComponentCollection{
					components: map[reflect.Type]Component{
						fooComponent{}.TypeOf(): fooComponent{},
						barComponent{}.TypeOf(): barComponent{},
					},
				},
			},
			1: &SimpleEntity{
				id: 1,
				components: &ComponentCollection{
					components: map[reflect.Type]Component{
						fooComponent{}.TypeOf(): fooComponent{},
						barComponent{}.TypeOf(): barComponent{},
					},
				},
			},
		},
	}
}

func createEntityControllerWithOneEntity() EntityController {
	return EntityController{
		nextEntityID: 1,
		components: map[reflect.Type][]int{
			fooComponent{}.TypeOf(): {0},
		},
		entities: map[int]Entity{
			0: &SimpleEntity{
				id: 0,
				components: &ComponentCollection{
					components: map[reflect.Type]Component{
						fooComponent{}.TypeOf(): fooComponent{},
					},
				},
			},
		},
	}
}

func TestNewEntityController(t *testing.T) {
	controller := NewEntityController()

	expected := &EntityController{
		nextEntityID: 0,
		components:   make(map[reflect.Type][]int),
		entities:     make(map[int]Entity),
	}

	assert.Equal(t, expected, controller)
}

func TestEntityController_Create(t *testing.T) {
	controller := EntityController{
		nextEntityID: 0,
		components:   map[reflect.Type][]int{},
		entities:     map[int]Entity{},
	}

	components := []Component{
		fooComponent{},
	}

	controller.Create(components)

	expected := EntityController{
		nextEntityID: 1,
		components: map[reflect.Type][]int{
			fooComponent{}.TypeOf(): {0},
		},
		entities: map[int]Entity{
			0: &SimpleEntity{
				id: 0,
				components: &ComponentCollection{
					components: map[reflect.Type]Component{
						fooComponent{}.TypeOf(): fooComponent{},
					},
				},
			},
		},
	}

	assert.Equal(t, expected, controller)
}

func TestEntityController_IsEntityExists(t *testing.T) {
	controller := createEntityController()

	assert.True(t, controller.IsEntityExists(0))
	assert.False(t, controller.IsEntityExists(999))
}

func TestEntityController_GetEntityById(t *testing.T) {
	controller := createEntityController()

	entity, err := controller.GetEntityById(0)

	expected := NewSimpleEntity(0)
	_ = expected.components.add(fooComponent{})
	_ = expected.components.add(barComponent{})

	assert.Equal(t, expected, entity)
	assert.Nil(t, err)
}

func TestEntityController_GetEntityByIdWhenEntityNotExists(t *testing.T) {
	controller := createEntityController()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Entity exists in entity controller.")
		}
	}()

	controller.GetEntityById(999)
}

func TestEntityController_GetEntitiesByComponent(t *testing.T) {
	controller := createEntityController()

	entities := controller.GetEntitiesByComponent(fooComponent{})

	expected := []Entity{
		&SimpleEntity{
			id: 0,
			components: &ComponentCollection{
				components: map[reflect.Type]Component{
					fooComponent{}.TypeOf(): fooComponent{},
					barComponent{}.TypeOf(): barComponent{},
				},
			},
		},
		&SimpleEntity{
			id: 1,
			components: &ComponentCollection{
				components: map[reflect.Type]Component{
					fooComponent{}.TypeOf(): fooComponent{},
					barComponent{}.TypeOf(): barComponent{},
				},
			},
		},
	}

	assert.Equal(t, expected, entities)
}

func TestEntityController_GetEntitiesByComponentWhenEntitiesNotExists(t *testing.T) {
	controller := createEntityController()

	entities := controller.GetEntitiesByComponent(buzComponent{})

	expected := []Entity{}

	assert.Equal(t, expected, entities)
}

func TestEntityController_GetEntitiesByComponentsBySeveralComponents(t *testing.T) {
	controller := createEntityController()

	entities := controller.GetEntitiesByComponents(fooComponent{}, barComponent{})

	expected := []Entity{
		&SimpleEntity{
			id: 0,
			components: &ComponentCollection{
				components: map[reflect.Type]Component{
					fooComponent{}.TypeOf(): fooComponent{},
					barComponent{}.TypeOf(): barComponent{},
				},
			},
		},
		&SimpleEntity{
			id: 1,
			components: &ComponentCollection{
				components: map[reflect.Type]Component{
					fooComponent{}.TypeOf(): fooComponent{},
					barComponent{}.TypeOf(): barComponent{},
				},
			},
		},
	}

	assert.Equal(t, expected, entities)
}

func TestEntityController_GetEntitiesByComponentsByOneComponent(t *testing.T) {
	controller := createEntityController()

	entities := controller.GetEntitiesByComponents(fooComponent{})

	expected := []Entity{
		&SimpleEntity{
			id: 0,
			components: &ComponentCollection{
				components: map[reflect.Type]Component{
					fooComponent{}.TypeOf(): fooComponent{},
					barComponent{}.TypeOf(): barComponent{},
				},
			},
		},
		&SimpleEntity{
			id: 1,
			components: &ComponentCollection{
				components: map[reflect.Type]Component{
					fooComponent{}.TypeOf(): fooComponent{},
					barComponent{}.TypeOf(): barComponent{},
				},
			},
		},
	}

	assert.Equal(t, expected, entities)
}

func TestEntityController_GetEntitiesWhenNotFound(t *testing.T) {
	controller := createEntityController()

	entities := controller.GetEntitiesByComponents(buzComponent{})

	var expected []Entity

	assert.Equal(t, expected, entities)
}

func TestEntityController_GetEntitiesWhenSendEmptySlice(t *testing.T) {
	controller := createEntityController()

	var args []Component

	entities := controller.GetEntitiesByComponents(args...)

	var expected []Entity

	assert.Equal(t, expected, entities)
}

func TestEntityController_DeleteEntityById(t *testing.T) {
	controller := createEntityController()

	controller.DeleteEntityById(0)

	expected := EntityController{
		nextEntityID: 2,
		components: map[reflect.Type][]int{
			fooComponent{}.TypeOf(): {1},
			barComponent{}.TypeOf(): {1},
		},
		entities: map[int]Entity{
			1: &SimpleEntity{
				id: 1,
				components: &ComponentCollection{
					components: map[reflect.Type]Component{
						fooComponent{}.TypeOf(): fooComponent{},
						barComponent{}.TypeOf(): barComponent{},
					},
				},
			},
		},
	}

	assert.Equal(t, expected, controller)
}

func TestEntityController_AddComponent(t *testing.T) {
	controller := createEntityControllerWithOneEntity()

	controller.AddComponent(0, barComponent{})

	expected := EntityController{
		nextEntityID: 1,
		components: map[reflect.Type][]int{
			fooComponent{}.TypeOf(): {0},
			barComponent{}.TypeOf(): {0},
		},
		entities: map[int]Entity{
			0: &SimpleEntity{
				id: 0,
				components: &ComponentCollection{
					components: map[reflect.Type]Component{
						fooComponent{}.TypeOf(): fooComponent{},
						barComponent{}.TypeOf(): barComponent{},
					},
				},
			},
		},
	}

	assert.Equal(t, expected, controller)
}

func TestEntityController_AddComponents(t *testing.T) {
	controller := createEntityControllerWithOneEntity()

	controller.AddComponents(0, []Component{barComponent{}, buzComponent{}})

	expected := EntityController{
		nextEntityID: 1,
		components: map[reflect.Type][]int{
			fooComponent{}.TypeOf(): {0},
			barComponent{}.TypeOf(): {0},
			buzComponent{}.TypeOf(): {0},
		},
		entities: map[int]Entity{
			0: &SimpleEntity{
				id: 0,
				components: &ComponentCollection{
					components: map[reflect.Type]Component{
						fooComponent{}.TypeOf(): fooComponent{},
						barComponent{}.TypeOf(): barComponent{},
						buzComponent{}.TypeOf(): buzComponent{},
					},
				},
			},
		},
	}

	assert.Equal(t, expected, controller)
}

func TestEntityController_RemoveComponentFromEntity(t *testing.T) {
	controller := createEntityController()

	controller.RemoveComponentFromEntity(0, fooComponent{})

	expected := EntityController{
		nextEntityID: 2,
		components: map[reflect.Type][]int{
			fooComponent{}.TypeOf(): {1},
			barComponent{}.TypeOf(): {0, 1},
		},
		entities: map[int]Entity{
			0: &SimpleEntity{
				id: 0,
				components: &ComponentCollection{
					components: map[reflect.Type]Component{
						barComponent{}.TypeOf(): barComponent{},
					},
				},
			},
			1: &SimpleEntity{
				id: 1,
				components: &ComponentCollection{
					components: map[reflect.Type]Component{
						fooComponent{}.TypeOf(): fooComponent{},
						barComponent{}.TypeOf(): barComponent{},
					},
				},
			},
		},
	}

	assert.Equal(t, expected, controller)
}

func TestEntityController_UpdateComponentForEntity(t *testing.T) {
	controller := EntityController{
		nextEntityID: 1,
		components: map[reflect.Type][]int{
			componentWithParameter{}.TypeOf(): {0},
		},
		entities: map[int]Entity{
			0: &SimpleEntity{
				id: 0,
				components: &ComponentCollection{
					components: map[reflect.Type]Component{
						componentWithParameter{}.TypeOf(): componentWithParameter{counter: 10},
					},
				},
			},
		},
	}

	controller.UpdateComponentForEntity(0, componentWithParameter{counter: 20})

	expected := EntityController{
		nextEntityID: 1,
		components: map[reflect.Type][]int{
			componentWithParameter{}.TypeOf(): {0},
		},
		entities: map[int]Entity{
			0: &SimpleEntity{
				id: 0,
				components: &ComponentCollection{
					components: map[reflect.Type]Component{
						componentWithParameter{}.TypeOf(): componentWithParameter{counter: 20},
					},
				},
			},
		},
	}

	assert.Equal(t, expected, controller)
}

func TestEntityController_RemoveItemFromEntityIds(t *testing.T) {
	entityIds := []int{0, 1, 2, 3, 4, 5}

	entityIds = removeItemFromEntityIds(entityIds, 3)
	entityIds = removeItemFromEntityIds(entityIds, 0)

	expected := []int{1, 2, 4, 5}

	assert.Equal(t, expected, entityIds)
}

type componentWithParameter struct {
	counter int
}

func (c componentWithParameter) TypeOf() reflect.Type {
	return reflect.TypeOf(c)
}
