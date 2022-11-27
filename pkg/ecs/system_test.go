package ecs_test

import (
	"github.com/alcoholiclobster/wherry-ecs/pkg/ecs"
	"github.com/stretchr/testify/mock"
)

type MockSystem struct {
	world ecs.World
	mock.Mock
}

func (s *MockSystem) Init() {
	s.Called()
}

func (s *MockSystem) Run() {
	s.Called()
}

func NewMockSystem(world ecs.World) *MockSystem {
	return &MockSystem{
		world,
		mock.Mock{},
	}
}
