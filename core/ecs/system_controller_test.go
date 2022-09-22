package ecs

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func createSystemController() SystemController {
	return SystemController{
		systems:       map[reflect.Type]System{},
		sortedSystems: map[int][]System{},
		priorityKeys:  []int{},
	}
}

type TestSystem struct {
	SystemRun bool
}

func (ts *TestSystem) Process() {
	ts.SystemRun = true
}

type AnotherSystem struct {
	SystemRun bool
}

func (as *AnotherSystem) Process() {
	as.SystemRun = true
}

type OneMoreSystem struct {
	SystemRun bool
}

func (oms *OneMoreSystem) Process() {
	oms.SystemRun = true
}

func TestNewSystemController(t *testing.T) {
	controller := NewSystemController()

	expected := &SystemController{
		systems:       map[reflect.Type]System{},
		sortedSystems: map[int][]System{},
		priorityKeys:  []int{},
	}

	assert.Equal(t, expected, controller)
}

func TestAddSystem(t *testing.T) {
	controller := createSystemController()

	// Add two systems, with different priorities
	controller.Add(&TestSystem{SystemRun: false}, 1)
	controller.Add(&TestSystem{SystemRun: false}, 1)
	controller.Add(&AnotherSystem{SystemRun: false}, 2)
	controller.Add(&OneMoreSystem{SystemRun: false}, 1)

	assert.True(t, controller.Has(reflect.TypeOf(&TestSystem{})))
	assert.True(t, controller.Has(reflect.TypeOf(&AnotherSystem{})))

	// Make sure the ordering of systems by priority is correct
	assert.Equal(t, len(controller.sortedSystems[1]), 2)
	assert.Equal(t, len(controller.sortedSystems[2]), 1)
}

func TestProcess(t *testing.T) {
	controller := createSystemController()

	system1 := &TestSystem{SystemRun: false}
	system2 := &AnotherSystem{SystemRun: false}
	system3 := &OneMoreSystem{SystemRun: false}

	controller.Add(system1, 1)
	controller.Add(system2, 2)
	controller.Add(system3, 3)

	controller.Process()

	assert.True(t, system1.SystemRun)
	assert.True(t, system2.SystemRun)
	assert.True(t, system3.SystemRun)
}

func TestProcessSystems(t *testing.T) {
	controller := createSystemController()

	system1 := &TestSystem{SystemRun: false}
	system2 := &AnotherSystem{SystemRun: false}
	system3 := &OneMoreSystem{SystemRun: false}

	controller.Add(system1, 1)
	controller.Add(system2, 2)
	controller.Add(system3, 3)

	controller.ProcessWithExcludedSystem([]reflect.Type{})

	assert.True(t, system1.SystemRun)
	assert.True(t, system2.SystemRun)
	assert.True(t, system3.SystemRun)

	// Create a new controller, and add all systems in the priority
	controller2 := createSystemController()

	system1 = &TestSystem{SystemRun: false}
	system2 = &AnotherSystem{SystemRun: false}
	system3 = &OneMoreSystem{SystemRun: false}

	controller2.Add(system1, 1)
	controller2.Add(system2, 1)
	controller2.Add(system3, 1)

	controller2.ProcessWithExcludedSystem([]reflect.Type{})

	assert.True(t, system1.SystemRun)
	assert.True(t, system2.SystemRun)
	assert.True(t, system3.SystemRun)

	// Create a new controller, and add all systems, but exclude one from processing
	controller3 := createSystemController()

	system1 = &TestSystem{SystemRun: false}
	system2 = &AnotherSystem{SystemRun: false}
	system3 = &OneMoreSystem{SystemRun: false}

	controller3.Add(system1, 1)
	controller3.Add(system2, 1)
	controller3.Add(system3, 1)

	// Exclude OneMoreSystem systems from processing
	controller3.ProcessWithExcludedSystem([]reflect.Type{reflect.TypeOf(&OneMoreSystem{})})

	assert.True(t, system1.SystemRun)
	assert.True(t, system2.SystemRun)
	assert.False(t, system3.SystemRun)

}

func TestProcessSingleSystem(t *testing.T) {
	controller := createSystemController()

	system1 := &TestSystem{SystemRun: false}
	system2 := &AnotherSystem{SystemRun: false}
	system3 := &OneMoreSystem{SystemRun: false}

	controller.Add(system1, 1)
	controller.Add(system2, 1)

	controller.ProcessSystem(reflect.TypeOf(&TestSystem{}))

	assert.True(t, system1.SystemRun)
	assert.False(t, system2.SystemRun)

	controller.ProcessSystem(reflect.TypeOf(&OneMoreSystem{}))
	assert.False(t, system3.SystemRun)
}

func TestSystemController_Process(t *testing.T) {
	gomockController := gomock.NewController(t)
	defer gomockController.Finish()
}
