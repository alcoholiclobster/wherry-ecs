package ecs

var lastMask int = 0

// Get component by mask and assert it to generic type
// Example: c := ecs.GetByMask[*MyComponent](entity, MyComponentMask)
func GetByMask[T Component](entity Entity, mask ComponentMask) T {
	return (*entity.Get(mask)).(T)
}

// Get component by component reference
// For example: c := ecs.GetRef(entity, &MyComponent{})
func GetRef[T Component](entity Entity, component T) T {
	return (*entity.Get(component.GetMask())).(T)
}

// Create new ComponentMask and ensure that it's unique
func NewMask() ComponentMask {
	mask := ComponentMask(1 << lastMask)
	lastMask++

	return mask
}
