package ecs

import (
	"fmt"
	"reflect"
)

type EntityController struct {
	// Значение id следующей, но еще не существующей, сущности.
	nextEntityID int
	// Карта где ключ это компонент, а значение список id сущностей с данным компонентом
	components map[reflect.Type][]int
	// Карта зарегестрированных сущностей
	entities map[int]Entity
}

func NewEntityController() *EntityController {
	return &EntityController{
		nextEntityID: 0,
		components:   make(map[reflect.Type][]int),
		entities:     make(map[int]Entity),
	}
}

func (controller *EntityController) Create(components []Component) (Entity, error) {
	entity := NewSimpleEntity(controller.nextEntityID)

	controller.entities[entity.id] = entity

	err := controller.AddComponents(entity.id, components)
	if err != nil {
		return nil, err
	}
	controller.nextEntityID++

	return entity, nil
}

func (controller *EntityController) IsEntityExists(entityId int) bool {
	_, ok := controller.entities[entityId]

	return ok
}

func (controller *EntityController) GetEntityById(entityId int) (Entity, error) {
	if entity, ok := controller.entities[entityId]; ok == true {
		return entity, nil
	}

	return nil, fmt.Errorf("entity with id [%d] not registered", entityId)
}

func (controller *EntityController) GetEntitiesByComponent(component Component) []Entity {
	var entities []Entity

	entityIds, ok := controller.components[component.TypeOf()]

	if ok == false {
		return []Entity{}
	}

	for _, entityId := range entityIds {
		entity, _ := controller.GetEntityById(entityId)
		entities = append(entities, entity)
	}

	return entities
}

func (controller *EntityController) GetEntitiesByComponents(components ...Component) []Entity {
	var foundEntities []Entity

	if len(components) == 0 {
		return foundEntities
	}

	component := components[0]

	entities := controller.GetEntitiesByComponent(component)

	for _, entity := range entities {
		if entity.Components().HasAll(components...) {
			foundEntities = append(foundEntities, entity)
		}
	}

	return foundEntities
}

func (controller *EntityController) DeleteEntityById(entityId int) {
	for key, entityIds := range controller.components {
		controller.components[key] = removeItemFromEntityIds(entityIds, entityId)
	}

	delete(controller.entities, entityId)
}

func (controller *EntityController) AddComponent(entityId int, component Component) error {
	entity, err := controller.GetEntityById(entityId)

	// TODO wrap error
	if err != nil {
		return err
	}

	componentType := component.TypeOf()

	controller.components[componentType] = append(controller.components[componentType], entityId)

	_ = entity.Components().add(component)

	return nil
}

func (controller *EntityController) AddComponents(entityId int, components []Component) error {
	entity, err := controller.GetEntityById(entityId)

	if err != nil {
		return err
	}

	for _, component := range components {
		componentType := component.TypeOf()

		controller.components[componentType] = append(controller.components[componentType], entityId)
		_ = entity.Components().add(component)
	}

	return nil
}

func (controller *EntityController) RemoveComponentFromEntity(entityId int, component Component) (Entity, error) {
	entity, err := controller.GetEntityById(entityId)
	if err != nil {
		return nil, err
	}

	_ = entity.Components().remove(component)

	entityIds := controller.components[component.TypeOf()]

	controller.components[component.TypeOf()] = removeItemFromEntityIds(entityIds, entityId)

	return entity, nil
}

func (controller *EntityController) UpdateComponentForEntity(entityId int, component Component) error {
	_, err := controller.RemoveComponentFromEntity(entityId, component)

	if err != nil {
		return err
	}

	err = controller.AddComponent(entityId, component)
	if err != nil {
		return err
	}

	return nil
}

func removeItemFromEntityIds(entityIds []int, target int) []int {
	counter := 0
	for _, entityId := range entityIds {
		if entityId == target {
			continue
		}

		entityIds[counter] = entityId
		counter++
	}

	return entityIds[:counter]
}
