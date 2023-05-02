package position

import (
	"unsafe"
)

type stableSorter interface {
	SortPosition(list []int, pos uint)
}

type PositionSorter struct {
	ss stableSorter
}

func New(s stableSorter) *PositionSorter {
	return &PositionSorter{ss: s}
}

func (p *PositionSorter) Sort(list []int) {
	size := unsafe.Sizeof(list[0])
	for i := 0; i < int(size)*8; i++ {
		p.ss.SortPosition(list, uint(i))
	}
}
