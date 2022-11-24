package ecs

type system struct {
	w *world
}

func (s *system) Run() {

}

func (s *system) GetWorld() World {
	return s.w
}
