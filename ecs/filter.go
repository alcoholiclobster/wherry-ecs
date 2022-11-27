package ecs

const resize_amount = 500

type filter struct {
	w    *world
	mask ComponentMask

	sparse []int
	dense  []int
}

func (f *filter) get() []Entity {
	entities := make([]Entity, len(f.dense))
	for i, item := range f.dense {
		entities[i] = f.w.entities[item]
	}

	return entities
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

// Delete entity from the filter if it exists
func (f *filter) del(e Entity) {
	if len(f.dense) == 0 {
		return
	}

	element := e.GetId()

	last := f.dense[len(f.dense)-1]
	f.dense[len(f.dense)-1] = f.dense[f.sparse[element]]
	f.dense[f.sparse[element]] = last

	tmp := f.sparse[element]
	f.sparse[element] = f.sparse[last]
	f.sparse[last] = tmp

	f.dense = f.dense[:len(f.dense)-1]
}

// check if given mask is matching filter
func (f *filter) check(m ComponentMask) bool {
	return m&f.mask == f.mask
}

func newFilter(world *world, mask ComponentMask) filter {
	return filter{
		w:      world,
		mask:   mask,
		sparse: make([]int, 1),
		dense:  make([]int, 0),
	}
}
