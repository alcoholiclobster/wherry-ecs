package ecs

type System interface{}

type InitSystem interface {
	Init()
}

type RunSystem interface {
	Run()
}
