package ecs_test

import (
	"fmt"
	"testing"

	"github.com/alcoholiclobster/go-ecs/ecs"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert := assert.New(t)
	world := ecs.NewWorld()

	e := world.NewEntity()
	assert.NotNil(e, "should create entity")

	e.Add(&FirstComponent{Num: 5})
	assert.Equal(5, (*e.Get(FirstComponentMask)).(*FirstComponent).Num, "should add component")
	assert.Equal(FirstComponentMask, e.GetMask(), "should update entity mask")

	e.Add(&FirstComponent{Num: 15})
	assert.Equal(15, ecs.Get(e, &FirstComponent{}).Num, "should overwrite component")

	e.Add(&SecondComponent{Text: "Hello"})
	assert.Equal("Hello", (*e.Get(SecondComponentMask)).(*SecondComponent).Text, "should add another component")
	assert.Equal(FirstComponentMask|SecondComponentMask, e.GetMask(), "should make mask of two components")

	e.Del(FirstComponentMask)
	assert.False(e.Has(FirstComponentMask), "should del component")
	assert.Nil(e.Get(FirstComponentMask), "should del component")
	assert.True(e.Has(SecondComponentMask), "should not del another component")
	assert.Equal(SecondComponentMask, e.GetMask(), "should subtract removed component from mask")

	e.Del(SecondComponentMask)
	assert.Equal(ecs.ComponentMask(0), e.GetMask(), "should return empty mask")
}

func TestDel(t *testing.T) {
	assert := assert.New(t)
	world := ecs.NewWorld()

	e := world.NewEntity()
	e.Add(&FirstComponent{Num: 5})
	e.Add(&SecondComponent{Text: "Test"})

	e.Del(FirstComponentMask)
	assert.False(e.Has(FirstComponentMask), "should delete component from entity")

	e.Del(SecondComponentMask)
	assert.False(e.IsValid(), "should destroy entity when last component is deleted")
}

func TestDestroyed(t *testing.T) {
	assert := assert.New(t)
	world := ecs.NewWorld()

	e := world.NewEntity()
	e.Destroy()

	assert.NotPanics(func() { e.Destroy() }, "should allow multiple Destroy calls")
	assert.Panics(func() { e.Add(&FirstComponent{}) }, "should panic")
	assert.Panics(func() { e.Get(FirstComponentMask) }, "should panic")
	assert.Panics(func() { e.Del(FirstComponentMask) }, "should panic")
	assert.Panics(func() { e.Has(FirstComponentMask) }, "should panic")
	assert.NotEmpty(fmt.Sprint(e), "should return string")
}
