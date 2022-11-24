package ecs

type filter struct {
	w    *world
	mask ComponentMask
}

func (f *filter) GetEntities() []Entity {
	entities := make([]Entity, len(f.w.entities))
	index := 0

	for _, e := range f.w.entities {
		if e != nil && e.IsValid() && e.GetMask()&f.mask == f.mask {
			entities[index] = e
			index++
		}
	}

	return entities[:index]
}
