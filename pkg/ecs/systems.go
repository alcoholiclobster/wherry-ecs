package ecs

type System interface {
	Init(args ...any)
	Run(args ...any)
}

type Systems interface {
	Add(system System) Systems

	Init(args ...any)
	Run(args ...any)
}

type systems struct {
	list   []System
	isInit bool
}

func (s *systems) Add(system System) Systems {
	s.list = append(s.list, system)

	return s
}

func (s *systems) Init(args ...any) {
	if s.isInit {
		panic("world is already initialized")
	}

	for _, s := range s.list {
		s.Init(args...)
	}

	s.isInit = true
}

func (s *systems) Run(args ...any) {
	if !s.isInit {
		panic("world is not initialized")
	}

	for _, s := range s.list {
		s.Run(args...)
	}
}

func NewSystems() Systems {
	return &systems{
		list: []System{},
	}
}
