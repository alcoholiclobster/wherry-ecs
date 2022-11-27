package ecs_test

import (
	"testing"

	"github.com/alcoholiclobster/go-ecs/ecs"
	"github.com/stretchr/testify/assert"
)

func TestMask(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(ecs.ComponentMask(1<<0), FirstComponentMask)
	assert.Equal(ecs.ComponentMask(1<<1), SecondComponentMask)
}

func TestFilter(t *testing.T) {
	assert := assert.New(t)
	world := ecs.NewWorld()

	assert.NotNil(world, "should create world")

	assert.Len(world.Filter(FirstComponentMask|SecondComponentMask), 0)

	world.NewEntity().
		Add(&FirstComponent{Num: 1}).
		Add(&SecondComponent{Text: "a"})

	world.NewEntity().
		Add(&FirstComponent{Num: 2}).
		Add(&SecondComponent{Text: "b"})

	world.NewEntity().
		Add(&FirstComponent{Num: 3})

	e1 := world.NewEntity().
		Add(&SecondComponent{Text: "c"})

	e2 := world.NewEntity().
		Add(&SecondComponent{Text: "d"})

	assert.Len(world.Filter(FirstComponentMask|SecondComponentMask), 2)
	assert.Len(world.Filter(FirstComponentMask), 3)
	assert.Len(world.Filter(SecondComponentMask), 4)

	ecs.Get(world.Filter(FirstComponentMask | SecondComponentMask)[0], &FirstComponent{}).Num = 123
	assert.Equal(
		123,
		ecs.Get(world.Filter(FirstComponentMask)[0], &FirstComponent{}).Num,
	)

	e1.Del(SecondComponentMask)
	assert.Len(world.Filter(SecondComponentMask), 3)

	e2.Destroy()
	assert.Len(world.Filter(SecondComponentMask), 2)
}

func TestAddSystem(t *testing.T) {
	world := ecs.NewWorld()
	messageSystem := NewMockSystem(world)
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
	assert.Panics(func() { world.Run() }, "should not allow Run before Init")

	world.Init()
	assert.Panics(func() { world.Init() }, "should not allow calling Init twice")

	entity := world.NewEntity()
	entity.Add(&FirstComponent{Num: 5})

	assert.NotEqual(
		world.NewEntity().Add(&SecondComponent{}).GetId(),
		entity.GetId(),
	)

	filter := world.Filter(FirstComponentMask)

	filterEntity := filter[0]
	assert.Equal(entity, filterEntity)

	world.Run()

	entity.Del(FirstComponentMask)
	assert.Len(world.Filter(FirstComponentMask), 0)
	assert.False(entity.IsValid(), "entity should destroy")

	world.Run()

	assert.False(entity.IsValid())
	assert.NotPanics(func() {
		world.NewEntity().Add(&FirstComponent{}).GetId()
	})
	assert.Len(world.Filter(FirstComponentMask), 1)
}
