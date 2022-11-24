package ecs_test

import (
	"testing"

	"github.com/alcoholiclobster/go-ecs/ecs"
	"github.com/stretchr/testify/assert"
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

func TestEntity(t *testing.T) {
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

func TestFilters(t *testing.T) {
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
	ecs.System

	filter ecs.Filter
}

func (s *MessageSystem) Run() {

}

func NewMessageSystem(world ecs.World) *MessageSystem {
	return &MessageSystem{
		filter: world.GetFilter(MessageComponentMask),
	}
}

func TestSystems(t *testing.T) {
	// assert := assert.New(t)

	world := ecs.NewWorld()

	world.AddSystem(NewMessageSystem(world))
	world.Run()
}
