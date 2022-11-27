package ecs

type world struct {
	entities []*entity
	filters  map[ComponentMask]*filter
	systems  []System
}

func (w *world) CreateEntity() Entity {
	e := entity{
		id:          0,
		mask:        0,
		world:       w,
		isDestroyed: false,
		components:  map[ComponentMask]Component{},
	}

	for i, e2 := range w.entities {
		if e2 == nil {
			e.id = int(i)
			return &e
		}
	}

	e.id = int(len(w.entities))
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
		f.add(e)
	}

	return f.get()
}

func (w *world) AddSystem(s System) World {
	w.systems = append(w.systems, s)

	return w
}

func (w *world) Init() {
	for _, s := range w.systems {
		if s, ok := s.(InitSystem); ok {
			s.Init()
			w.cleanup()
		}
	}

	if len(w.systems) == 0 {
		w.cleanup()
	}
}

func (w *world) Run() {
	for _, s := range w.systems {
		if s, ok := s.(RunSystem); ok {
			s.Run()
			w.cleanup()
		}
	}

	if len(w.systems) == 0 {
		w.cleanup()
	}
}

func (w *world) cleanup() {
	for i, e := range w.entities {
		if e.isDestroyed || e.mask == 0 {
			e.isDestroyed = true
			w.entities[i] = nil
		}
	}
}

func (w *world) addEntityToFilters(mask ComponentMask, entity Entity) {
	if mask == 0 {
		return
	}

	for _, f := range w.filters {
		f.add(entity)
	}
}

func (w *world) removeEntityFromFilters(mask ComponentMask, entity Entity) {
	for _, f := range w.filters {
		if f.check(mask) {
			f.del(entity)
		}
	}
}

func NewWorld() World {
	return &world{
		make([]*entity, 0),
		make(map[ComponentMask]*filter),
		make([]System, 0),
	}
}
