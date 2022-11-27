package ecs

type ComponentMask uint64

type World interface {
	NewEntity() Entity
	AddSystem(system System) World
	Filter(mask ComponentMask) []Entity

	Init()
	Run()
}

type Entity interface {
	GetMask() ComponentMask
	GetId() int

	Add(component Component) Entity
	Get(mask ComponentMask) *Component
	Has(mask ComponentMask) bool
	Del(mask ComponentMask) Entity

	IsValid() bool
	Destroy()
}

type Component interface {
	GetMask() ComponentMask
}

type System interface{}

type InitSystem interface {
	Init()
}

type RunSystem interface {
	Run()
}
