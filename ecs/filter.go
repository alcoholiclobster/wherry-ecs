package ecs

const resize_amount = 500

type filter struct {
	world *world
	mask  ComponentMask

	sparse []int
	dense  []int
}

// Check if given mask is matching filter
func (f *filter) check(m ComponentMask) bool {
	return m&f.mask == f.mask
}

// Add entity to the filter if entity mask is matching
func (f *filter) add(e Entity) {
	if !f.check(e.GetMask()) {
		return
	}

	entityId := e.GetId()

	pos := len(f.dense)
	f.dense = append(f.dense, entityId)

	if len(f.sparse) < entityId+1 {
		f.sparse = append(f.sparse, make([]int, resize_amount)...)
	}
	f.sparse[entityId] = pos
}

func (f *filter) get() []Entity {
	entities := make([]Entity, len(f.dense))
	for i, item := range f.dense {
		entities[i] = f.world.entities[item]
	}

	return entities
}

// Delete entity from the filter
func (f *filter) del(e Entity) {
	if len(f.dense) == 0 {
		return
	}

	entityId := e.GetId()

	dlen := len(f.dense) - 1
	last := f.dense[dlen]
	// swap items
	f.dense[dlen], f.dense[f.sparse[entityId]] = f.dense[f.sparse[entityId]], f.dense[dlen]
	f.sparse[entityId], f.sparse[last] = f.sparse[last], f.sparse[entityId]

	f.dense = f.dense[:dlen]
}

func newFilter(world *world, mask ComponentMask) filter {
	return filter{
		world: world,
		mask:  mask,

		sparse: make([]int, 1),
		dense:  make([]int, 0),
	}
}
