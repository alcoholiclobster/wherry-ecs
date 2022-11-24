package ecs

type world struct {
	entities []*entity
	filters  map[ComponentMask]*filter
	systems  []System
}

func (w *world) CreateEntity() Entity {
	newEntity := entity{
		id:          0,
		mask:        0,
		components:  map[ComponentMask]Component{},
		isDirty:     false,
		isDestroyed: false,
	}

	for i, e := range w.entities {
		if e == nil {
			newEntity.id = i
			return &newEntity
		}
	}

	newEntity.id = len(w.entities)
	w.entities = append(w.entities, &newEntity)

	return &newEntity
}

func (w *world) GetFilter(mask ComponentMask) Filter {
	if f, ok := w.filters[mask]; ok {
		return f
	}

	f := filter{w, mask}

	return &f
}

func (w *world) AddSystem(s System) World {
	w.systems = append(w.systems, s)

	return w
}

func (w *world) Init() World {
	for _, s := range w.systems {
		if s, ok := s.(InitSystem); ok {
			s.Init()
		}
	}

	return w
}

func (w *world) Run() {
	for i, e := range w.entities {
		if e.isDestroyed {
			w.entities[i] = nil
		}
	}
	for _, s := range w.systems {
		if s, ok := s.(RunSystem); ok {
			s.Run()
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
