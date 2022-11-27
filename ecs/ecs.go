package ecs

type ComponentMask uint64

type World interface {
	CreateEntity() Entity
	Filter(mask ComponentMask) []Entity
	AddSystem(system System) World

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

type InitSystem interface {
	Init()
}

type RunSystem interface {
	Run()
}

type System interface{}
