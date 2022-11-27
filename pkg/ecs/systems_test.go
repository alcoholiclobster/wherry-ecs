package ecs_test

import (
	"testing"

	"github.com/alcoholiclobster/wherry-ecs/pkg/ecs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockSystem struct {
	world ecs.World
	mock.Mock
}

func (s *MockSystem) Init(args ...any) {
	s.MethodCalled("Init", args...)
}

func (s *MockSystem) Run(args ...any) {
	s.MethodCalled("Run", args...)
}

func NewMockSystem(world ecs.World) *MockSystem {
	return &MockSystem{
		world,
		mock.Mock{},
	}
}

func TestSystemsLoop(t *testing.T) {
	assert := assert.New(t)
	systems := ecs.NewSystems()

	assert.Panics(func() { systems.Run() }, "should not allow Run before Init")
	assert.NotPanics(func() { systems.Init() }, "should initialize")
	assert.Panics(func() { systems.Init() }, "should not allow calling Init twice")
	assert.NotPanics(func() { systems.Run() }, "should run after initialization")
}

func TestSystemsAdd(t *testing.T) {
	world := ecs.NewWorld()
	systems := ecs.NewSystems()

	messageSystem := NewMockSystem(world)
	messageSystem.On("Run", mock.Anything, mock.Anything).Return()
	messageSystem.On("Init", mock.Anything).Return()
	messageSystem.On(mock.Anything)

	systems.
		Add(messageSystem).
		Init(123)

	messageSystem.AssertCalled(t, "Init", 123)

	systems.Run("Test Value 1", 1234)

	messageSystem.AssertCalled(t, "Run", "Test Value 1", 1234)
}
