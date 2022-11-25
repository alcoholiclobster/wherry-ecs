package ecs

type world struct {
	entities     []*entity
	maskEntities map[ComponentMask]map[entityId]bool
	filters      map[ComponentMask]*filter
	systems      []System
}

func (w *world) CreateEntity() Entity {
	e := entity{
		id:    0,
		mask:  0,
		world: w,

		isDirty:     false,
		isDestroyed: false,

		components: map[ComponentMask]Component{},
	}

	for i, e2 := range w.entities {
		if e2 == nil {
			e.id = entityId(i)
			return &e
		}
	}

	e.id = entityId(len(w.entities))
	w.entities = append(w.entities, &e)

	return &e
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

func (w *world) addMaskEntity(mask ComponentMask, id entityId) {
	if mask == 0 {
		return
	}

	me, ok := w.maskEntities[mask]
	if !ok {
		me = make(map[entityId]bool)
		w.maskEntities[mask] = me
	}

	me[id] = true
}

func (w *world) removeMaskEntity(mask ComponentMask, id entityId) {
	me, ok := w.maskEntities[mask]
	if !ok {
		return
	}

	delete(me, id)
}

func NewWorld() World {
	return &world{
		make([]*entity, 0),
		make(map[ComponentMask]map[entityId]bool),
		make(map[ComponentMask]*filter),
		make([]System, 0),
	}
}
