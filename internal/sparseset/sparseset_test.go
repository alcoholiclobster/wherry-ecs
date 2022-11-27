package sparseset_test

import (
	"testing"

	"github.com/alcoholiclobster/wherry-ecs/internal/sparseset"
	"github.com/stretchr/testify/assert"
)

func TestSetAdd(t *testing.T) {
	assert := assert.New(t)

	s := sparseset.NewSparseSet(100)
	s.Add(20)
	s.Add(123)
	assert.Len(s.GetElements(), 2)

	assert.Panics(func() {
		s.Add(-1)
	}, "should not add negative values")
	assert.Len(s.GetElements(), 2)
}

func TestSetRemove(t *testing.T) {
	assert := assert.New(t)

	s := sparseset.NewSparseSet(100)
	s.Add(1500)
	s.Remove(1500)
	s.Remove(123)
	assert.Len(s.GetElements(), 0)

	assert.Panics(func() {
		s.Remove(-1)
	}, "should not add negative values")
}
