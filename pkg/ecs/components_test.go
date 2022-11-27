package ecs_test

import (
	"github.com/alcoholiclobster/wherry-ecs/pkg/ecs"
)

var (
	FirstComponentMask  = ecs.NewMask()
	SecondComponentMask = ecs.NewMask()
)

type FirstComponent struct {
	Num int
}

func (c *FirstComponent) GetMask() ecs.ComponentMask {
	return FirstComponentMask
}

type SecondComponent struct {
	Text string
}

func (c *SecondComponent) GetMask() ecs.ComponentMask {
	return SecondComponentMask
}
