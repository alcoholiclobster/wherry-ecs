package ecs

import "github.com/alcoholiclobster/wherry-ecs/internal/sparseset"

type filter struct {
	world    *world
	mask     ComponentMask
	entities sparseset.SparseSet
}

// Check if given mask is matching filter
func (f *filter) check(mask ComponentMask) bool {
	return mask&f.mask == f.mask
}

// Add entity to the filter if entity mask is matching
func (f *filter) add(e Entity) {
	if f.check(e.GetMask()) {
		f.entities.Add(e.GetId())
	}
}

func (f *filter) get() []Entity {
	elements := f.entities.GetElements()
	entities := make([]Entity, len(elements))

	for i, item := range elements {
		entities[i] = f.world.entities[item]
	}

	return entities
}

// Delete entity from the filter
func (f *filter) del(e Entity) {
	f.entities.Remove(e.GetId())
}

func newFilter(world *world, mask ComponentMask) filter {
	return filter{
		world:    world,
		mask:     mask,
		entities: sparseset.NewSparseSet(100),
	}
}
