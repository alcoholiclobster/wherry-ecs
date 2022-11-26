package ecs

type filter struct {
	w        *world
	entities []Entity
	mask     ComponentMask
}

func (f *filter) GetEntities() []Entity {
	return f.entities
	// entities := make([]Entity, len(f.w.entities))
	// index := 0

	// mask := ComponentMask(0)
	// for m := range f.w.maskEntities {
	// 	if m&f.mask == f.mask {
	// 		break
	// 	}
	// }

	// if mask == 0 {
	// 	return make([]Entity, 0)
	// }

	// if me, ok := f.w.maskEntities[mask]; ok {
	// 	for id := range me {
	// 		e := f.w.entities[id]
	// 		if e != nil && e.IsValid() {
	// 			entities[index] = e
	// 			index++
	// 		}
	// 	}
	// }

	// return entities[:index]
}
