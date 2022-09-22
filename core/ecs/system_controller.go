package ecs

import (
	"reflect"
	"sort"
)

type SystemController struct {
	// Карта зарегестрированных систем.
	systems map[reflect.Type]System
	// Список зарегестрированных систем, отсортированных по весу.
	sortedSystems map[int][]System
	// Лист значений приоритетов
	priorityKeys []int
}

func NewSystemController() *SystemController {
	return &SystemController{
		systems:       map[reflect.Type]System{},
		sortedSystems: map[int][]System{},
		priorityKeys:  []int{},
	}
}

// Add registers a system to the controller. A priority can be provided, and systems will be processed in
// numeric order, low to high. If multiple systems are registered as the same priority, they will be randomly run within
// that priority group.
func (registry *SystemController) Add(system System, priority int) {
	systemType := reflect.TypeOf(system)

	if _, ok := registry.systems[systemType]; !ok {
		// A system of this type has not been added yet, so add it to the systems list
		registry.systems[systemType] = system

		// Now, append the system to a special list that will be used for sorting by priority
		if !intInSlice(priority, registry.priorityKeys) {
			registry.priorityKeys = append(registry.priorityKeys, priority)
		}
		registry.sortedSystems[priority] = append(registry.sortedSystems[priority], system)
		sort.Ints(registry.priorityKeys)
	}
}

// Has checks the controller to see if it has a given system associated with it
func (registry *SystemController) Has(systemType reflect.Type) bool {
	if _, ok := registry.systems[systemType]; ok {
		return true
	}

	return false
}

// Process kicks off system processing for all systems attached to the controller. Systems will be processed in the
// order they are found, or if they have a priority, in priority order. If there is a mix of systems with priority and
// without, systems with priority will be processed first (in order).
func (registry *SystemController) Process() {
	for _, key := range registry.priorityKeys {
		for _, system := range registry.sortedSystems[key] {
			system.Process()
		}
	}
}

// ProcessWithExcludedSystem kicks off system processing for all systems attached to the controller. Systems will be processed in the
// order they are found, or if they have a priority, in priority order. If there is a mix of systems with priority and
// without, systems with priority will be processed first (in order).
func (registry *SystemController) ProcessWithExcludedSystem(excludedSystems []reflect.Type) {
	for _, key := range registry.priorityKeys {
		for _, system := range registry.sortedSystems[key] {
			systemType := reflect.TypeOf(system)

			// Check if the current system type was marked as excluded on this call. If it was, not process it.
			if !typeInSlice(systemType, excludedSystems) {
				system.Process()
			}
		}
	}
}

// ProcessSystem allows for on demand processing of individual systems, rather than processing all at once via ProcessWithExcludedSystem
func (registry *SystemController) ProcessSystem(systemType reflect.Type) {
	if registry.Has(systemType) {
		system := registry.systems[systemType]
		system.Process()
	}
}
