package ecs_test

import (
	"testing"

	"github.com/alcoholiclobster/go-ecs/ecs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	ValueComponentMask   = ecs.ComponentMask(1 << 0)
	MessageComponentMask = ecs.ComponentMask(1 << 1)
)

type ValueComponent struct {
	Value int
}

func (c *ValueComponent) GetMask() ecs.ComponentMask {
	return ValueComponentMask
}

type MessageComponent struct {
	Message string
}

func (c *MessageComponent) GetMask() ecs.ComponentMask {
	return MessageComponentMask
}

func TestComponent(t *testing.T) {
	assert := assert.New(t)

	world := ecs.NewWorld()
	assert.NotNil(world, "should create world")

	entity := world.CreateEntity()
	assert.NotNil(entity, "should create entity")

	entity.Add(&ValueComponent{Value: 5})
	assert.Equal(5, (*entity.Get(ValueComponentMask)).(*ValueComponent).Value, "should add component")
	assert.Equal(ValueComponentMask, entity.GetMask(), "should update entity mask")

	entity.Add(&ValueComponent{Value: 15})
	assert.Equal(15, (*entity.Get(ValueComponentMask)).(*ValueComponent).Value, "should overwrite component")

	entity.Add(&MessageComponent{Message: "Hello"})
	assert.Equal("Hello", (*entity.Get(MessageComponentMask)).(*MessageComponent).Message, "should add another component")
	assert.Equal(ValueComponentMask|MessageComponentMask, entity.GetMask(), "should make mask of two components")

	entity.Del(ValueComponentMask)
	assert.False(entity.Has(ValueComponentMask), "should del component")
	assert.Nil(entity.Get(ValueComponentMask), "should del component")
	assert.True(entity.Has(MessageComponentMask), "should not del another component")
	assert.Equal(MessageComponentMask, entity.GetMask(), "should subtract removed component from mask")

	entity.Del(MessageComponentMask)
	assert.Equal(ecs.ComponentMask(0), entity.GetMask(), "should return empty mask")
}

func TestFilter(t *testing.T) {
	assert := assert.New(t)

	world := ecs.NewWorld()

	world.CreateEntity().
		Add(&ValueComponent{Value: 1}).
		Add(&MessageComponent{Message: "a"})

	world.CreateEntity().
		Add(&ValueComponent{Value: 2}).
		Add(&MessageComponent{Message: "b"})

	world.CreateEntity().
		Add(&ValueComponent{Value: 3})

	world.CreateEntity().
		Add(&MessageComponent{Message: "c"})

	world.CreateEntity().
		Add(&MessageComponent{Message: "d"})

	assert.Len(world.GetFilter(ValueComponentMask|MessageComponentMask).GetEntities(), 2)
	assert.Len(world.GetFilter(ValueComponentMask).GetEntities(), 3)
	assert.Len(world.GetFilter(MessageComponentMask).GetEntities(), 4)

	(*world.GetFilter(ValueComponentMask | MessageComponentMask).GetEntities()[0].Get(ValueComponentMask)).(*ValueComponent).Value = 123
	assert.Equal(123, (*world.GetFilter(ValueComponentMask).GetEntities()[0].Get(ValueComponentMask)).(*ValueComponent).Value)
}

type MessageSystem struct {
	world  ecs.World
	filter ecs.Filter

	mock.Mock
}

func (s *MessageSystem) Init() {
	s.Called()
}

func (s *MessageSystem) Run() {
	s.Called()
}

func NewMessageSystem(world ecs.World) *MessageSystem {
	return &MessageSystem{
		world,
		world.GetFilter(MessageComponentMask),

		mock.Mock{},
	}
}

func TestSystem(t *testing.T) {
	world := ecs.NewWorld()
	messageSystem := NewMessageSystem(world)
	messageSystem.On("Run").Return()
	messageSystem.On("Init").Return()

	world.
		AddSystem(messageSystem).
		Init()

	messageSystem.AssertCalled(t, "Init")

	world.Run()

	messageSystem.AssertCalled(t, "Run")
}

func TestEntity(t *testing.T) {
	assert := assert.New(t)
	world := ecs.NewWorld()

	entity := world.CreateEntity()
	entity.Add(&ValueComponent{Value: 5})

	filter := world.GetFilter(ValueComponentMask)

	filterEntity := filter.GetEntities()[0]
	assert.Equal(entity, filterEntity)

	world.Run()

	entity.Del(ValueComponentMask)
	assert.Len(filter.GetEntities(), 0)

	world.Run()

	assert.False(entity.IsValid())
}
