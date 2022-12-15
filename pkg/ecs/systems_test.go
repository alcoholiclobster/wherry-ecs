package ecs_test

import (
	"testing"

	"github.com/alcoholiclobster/wherry-ecs/pkg/ecs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockSystem struct {
	mock.Mock
}

func (s *MockSystem) Init(world ecs.World, args ...any) {
	s.MethodCalled("Init", args...)
}

func (s *MockSystem) Run(world ecs.World, args ...any) {
	s.MethodCalled("Run", args...)
}

func NewMockSystem() *MockSystem {
	return &MockSystem{
		mock.Mock{},
	}
}

func TestSystemsLoop(t *testing.T) {
	assert := assert.New(t)
	systems := ecs.NewSystems(ecs.NewWorld())

	assert.Panics(func() { systems.Run() }, "should not allow Run before Init")
	assert.NotPanics(func() { systems.Init() }, "should initialize")
	assert.Panics(func() { systems.Init() }, "should not allow calling Init twice")
	assert.NotPanics(func() { systems.Run() }, "should run after initialization")
}

func TestSystemsAdd(t *testing.T) {
	world := ecs.NewWorld()
	systems := ecs.NewSystems(world)

	mockSystem := NewMockSystem()
	mockSystem.On("Run", mock.Anything, mock.Anything).Return()
	mockSystem.On("Init", mock.Anything).Return()
	mockSystem.On(mock.Anything)

	systems.
		Add(mockSystem).
		Init(123)

	mockSystem.AssertCalled(t, "Init", 123)

	systems.Run("Test Value 1", 1234)

	mockSystem.AssertCalled(t, "Run", "Test Value 1", 1234)
}
