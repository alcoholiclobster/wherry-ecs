package sparseset

type SparseSet interface {
	Add(element int)
	Remove(element int)
	GetElements() []int
}

type sparseSet struct {
	sparse []int
	dense  []int

	minResizeCount int
}

func (s *sparseSet) Add(element int) {
	if element < 0 {
		panic("element should be a non-negative integer")
	}

	pos := len(s.dense)
	s.dense = append(s.dense, element)

	if len(s.sparse) < element+1 {
		count := element + 1 - len(s.sparse)
		if count < s.minResizeCount {
			count = s.minResizeCount
		}

		s.sparse = append(s.sparse, make([]int, count)...)
		print(len(s.sparse))
	}
	s.sparse[element] = pos
}

func (s *sparseSet) Remove(element int) {
	if element < 0 {
		panic("element should be a non-negative integer")
	}

	if len(s.dense) == 0 {
		return
	}

	dlen := len(s.dense) - 1
	last := s.dense[dlen]
	// swap items
	s.dense[dlen], s.dense[s.sparse[element]] = s.dense[s.sparse[element]], s.dense[dlen]
	s.sparse[element], s.sparse[last] = s.sparse[last], s.sparse[element]

	s.dense = s.dense[:dlen]
}

func (s *sparseSet) GetElements() []int {
	return s.dense
}

func NewSparseSet(minResizeCount int) SparseSet {
	return &sparseSet{
		sparse: []int{},
		dense:  []int{},

		minResizeCount: minResizeCount,
	}
}
