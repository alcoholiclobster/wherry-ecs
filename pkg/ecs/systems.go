package ecs

type RunSystem interface {
	Run(World)
}

type InitSystem interface {
	Init(World)
}

type Systems interface {
	Add(system any) Systems

	Init(args ...any)
	Run(args ...any)
}

type systems struct {
	isInitialized bool

	runSystems  []RunSystem
	initSystems []InitSystem

	world World
}

// Adds a RunSystem or InitSystem to systems
func (s *systems) Add(system any) Systems {
	if runSystem, ok := system.(RunSystem); ok {
		s.runSystems = append(s.runSystems, runSystem)
	}

	if initSystem, ok := system.(InitSystem); ok {
		s.initSystems = append(s.initSystems, initSystem)
	}

	return s
}

func (s *systems) Init(args ...any) {
	if s.isInitialized {
		panic("systems are already initialized")
	}

	for _, system := range s.initSystems {
		system.Init(s.world)
	}

	s.isInitialized = true
}

func (s *systems) Run(args ...any) {
	if !s.isInitialized {
		panic("systems are not initialized")
	}

	for _, system := range s.runSystems {
		system.Run(s.world)
	}
}

func NewSystems(world World) Systems {
	return &systems{
		runSystems:  []RunSystem{},
		initSystems: []InitSystem{},
		world:       world,
	}
}
