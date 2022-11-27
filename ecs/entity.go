package ecs

import "fmt"

type entity struct {
	id          int
	mask        ComponentMask
	world       *world
	components  map[ComponentMask]Component
	isDestroyed bool
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
		panic("entity is destroyed")
	}

	e.components[component.GetMask()] = component
	e.setMask(e.mask | component.GetMask())

	return e
}

func (e entity) Has(mask ComponentMask) bool {
	if e.isDestroyed {
		panic("entity is destroyed")
	}

	_, ok := e.components[mask]
	return ok
}

func (e *entity) Get(mask ComponentMask) *Component {
	if e.isDestroyed {
		panic("entity is destroyed")
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
		panic("entity is destroyed")
	}

	delete(e.components, mask)
	e.setMask(e.mask &^ mask)

	// Last component was deleted
	if e.mask == 0 {
		e.Destroy()
	}

	return e
}

func (e *entity) Destroy() {
	if e.isDestroyed {
		return
	}

	e.world.removeEntityFromFilters(e.mask, e, true)
	e.mask = 0
	e.components = nil
	e.isDestroyed = true
}

func (e entity) String() string {
	if !e.IsValid() {
		return "Entity(destroyed)"
	}
	return fmt.Sprintf("Entity(%d)", e.id)
}

func (e *entity) setMask(mask ComponentMask) {
	e.world.removeEntityFromFilters(e.mask, e, false)
	e.mask = mask
	e.world.addEntityToFilters(e.mask, e)
}

func newEntity(world *world) entity {
	return entity{
		id:          0,
		mask:        0,
		world:       world,
		components:  map[ComponentMask]Component{},
		isDestroyed: false,
	}
}
