package ecs

type entityId int

type entity struct {
	id    entityId
	mask  ComponentMask
	world *world

	isDirty     bool
	isDestroyed bool

	components map[ComponentMask]Component
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
	e.setMask(e.mask | component.GetMask())
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
	e.setMask(e.mask &^ mask)
	e.isDirty = true

	return e
}

func (e *entity) setMask(mask ComponentMask) {
	e.world.removeMaskEntity(e.mask, e.id)
	e.mask = mask
	e.world.addMaskEntity(e.mask, e.id)
}

func (e *entity) Destroy() {
	if e.isDestroyed {
		panic("access deleted entity")
	}

	e.isDestroyed = true
	e.world.removeMaskEntity(e.mask, e.id)
	e.mask = 0
}
