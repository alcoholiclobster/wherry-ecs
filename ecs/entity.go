package ecs

type entity struct {
	id          int
	mask        ComponentMask
	components  map[ComponentMask]Component
	isDirty     bool
	isDestroyed bool
}

func (e entity) GetMask() ComponentMask {
	return e.mask
}

func (e entity) IsValid() bool {
	return !e.isDestroyed
}

func (e *entity) Add(component Component) Entity {
	if e.isDestroyed {
		panic("access deleted entity")
	}

	e.components[component.GetMask()] = component
	e.mask |= component.GetMask()
	e.isDirty = true

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
	e.mask &^= mask
	e.isDirty = true

	return e
}

func (e *entity) Destroy() {
	if e.isDestroyed {
		panic("access deleted entity")
	}

	for mask := range e.components {
		e.Del(mask)
	}

	e.isDestroyed = true
}
