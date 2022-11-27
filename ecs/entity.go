package ecs

type entity struct {
	id    int
	mask  ComponentMask
	world *world

	isDestroyed bool

	components map[ComponentMask]Component
}

func (e entity) GetMask() ComponentMask {
	return e.mask
}

func (e entity) GetId() int {
	return e.id
}

func (e entity) IsValid() bool {
	return !e.isDestroyed
}

func (e *entity) Add(component Component) Entity {
	if e.isDestroyed {
		panic("access deleted entity")
	}

	e.components[component.GetMask()] = component
	e.setMask(e.mask | component.GetMask())

	return e
}

func (e entity) Has(mask ComponentMask) bool {
	if e.isDestroyed {
		panic("access deleted entity")
	}

	_, ok := e.components[mask]
	return ok
}

func (e *entity) Get(mask ComponentMask) *Component {
	if e.isDestroyed {
		panic("access deleted entity")
	}

	component, ok := e.components[mask]

	if ok {
		return &component
	} else {
		return nil
	}
}

func (e *entity) Del(mask ComponentMask) Entity {
	if e.isDestroyed {
		panic("access deleted entity")
	}

	delete(e.components, mask)
	e.setMask(e.mask &^ mask)

	return e
}

func (e *entity) setMask(mask ComponentMask) {
	e.world.removeEntityFromFilters(e.mask, e)
	e.mask = mask
	e.world.addEntityToFilters(e.mask, e)
}

func (e *entity) Destroy() {
	if e.isDestroyed {
		panic("access deleted entity")
	}

	e.isDestroyed = true
	e.world.removeEntityFromFilters(e.mask, e)
	e.mask = 0
}
