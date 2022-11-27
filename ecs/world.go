package ecs

type World interface {
	NewEntity() Entity
	AddSystem(system System) World
	Filter(mask ComponentMask) []Entity

	Init()
	Run()
}

type world struct {
	entities []*entity
	filters  map[ComponentMask]*filter
	systems  []System

	isInitialized bool
}

func (w *world) NewEntity() Entity {
	e := newEntity(w)

	for i, e2 := range w.entities {
		if e2 == nil {
			e.id = i
			w.entities[i] = &e
			return &e
		}
	}

	e.id = len(w.entities)
	w.entities = append(w.entities, &e)

	return &e
}

func (w *world) Filter(mask ComponentMask) []Entity {
	if f, ok := w.filters[mask]; ok {
		return f.get()
	}

	f := newFilter(w, mask)
	w.filters[mask] = &f

	// Add matching existing entities to filter
	for _, e := range w.entities {
		if e != nil {
			f.add(e)
		}
	}

	return f.get()
}

func (w *world) AddSystem(s System) World {
	w.systems = append(w.systems, s)

	return w
}

func (w *world) Init() {
	if w.isInitialized {
		panic("world is already initalized")
	}

	for _, s := range w.systems {
		if s, ok := s.(InitSystem); ok {
			s.Init()
		}
	}

	w.isInitialized = true
}

func (w *world) Run() {
	if !w.isInitialized {
		panic("world is not initalized")
	}

	for _, s := range w.systems {
		if s, ok := s.(RunSystem); ok {
			s.Run()
		}
	}
}

func (w *world) addEntityToFilters(mask ComponentMask, e *entity) {
	if mask != 0 {
		for _, f := range w.filters {
			f.add(e)
		}
	}
}

func (w *world) removeEntityFromFilters(mask ComponentMask, e *entity) {
	if e.mask != 0 {
		for _, f := range w.filters {
			if f.check(mask) {
				f.del(e)
			}
		}
	}

	if !e.IsValid() {
		w.entities[e.id] = nil
	}
}

func (w world) String() string {
	return "World"
}

func NewWorld() World {
	return &world{
		entities: make([]*entity, 0),
		systems:  make([]System, 0),
		filters:  make(map[ComponentMask]*filter),

		isInitialized: false,
	}
}
