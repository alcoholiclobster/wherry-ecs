package ecs

type ComponentMask uint64

type Component interface {
	GetMask() ComponentMask
}

var lastComponentMaskIndex int = 0

// Create new ComponentMask and ensure that it's unique
func NewMask() ComponentMask {
	mask := ComponentMask(1 << lastComponentMaskIndex)
	lastComponentMaskIndex++

	return mask
}
