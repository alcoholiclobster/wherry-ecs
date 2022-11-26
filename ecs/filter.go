package ecs

type filter struct {
	w        *world
	entities []Entity
	mask     ComponentMask
}

func (f *filter) GetEntities() []Entity {
	return f.entities
}
