package heap

import (
	"fmt"
	"math"
)

type Type = string

const (
	Min Type = "min"
	Max Type = "max"
)

type Heap struct {
	hType     Type
	edgeValue int
}

func New(hType Type) *Heap {
	var edgeValue int
	switch hType {
	case Min:
		edgeValue = math.MaxInt
	case Max:
		edgeValue = math.MinInt
	}
	return &Heap{hType: hType, edgeValue: edgeValue}
}

func (h *Heap) Sort(list []int) {
	h.buildHeap(list)
	heapSize := len(list)
	for i := len(list) - 1; i >= 0; i-- {
		top := list[0]
		list[0] = list[i]
		list[i] = top
		heapSize -= 1
		h.heapify(list, i, 0)
	}
}

func (h *Heap) buildHeap(list []int) {
	for i := len(list) / 2; i >= 0; i-- {
		h.heapify(list, len(list), i)
	}
}

func (h *Heap) heapify(list []int, heapSize, i int) {
	if i >= heapSize {
		return
	}

	left := left(i)
	right := right(i)

	if right >= heapSize && left >= heapSize {
		return
	}

	val := list[i]
	var leftVal int

	if left < heapSize {
		leftVal = list[left]
	} else {
		leftVal = h.edgeValue
	}

	var rightVal int
	if right < heapSize {
		rightVal = list[right]
	} else {
		rightVal = h.edgeValue
	}

	var dest int
	var destVal int
	if h.compare(rightVal, leftVal) {
		dest = left
		destVal = leftVal
	} else if right < heapSize {
		dest = right
		destVal = rightVal
	} else {
		return
	}

	if h.compare(val, destVal) {
		list[dest] = val
		list[i] = destVal
		h.heapify(list, heapSize, dest)
	}
}

func (h *Heap) compare(a, b int) bool {
	switch h.hType {
	case Min:
		return a > b
	case Max:
		return b > a
	default:
		fmt.Println("Unknown heap type")

	}
	return false
}

func parent(i int) int {
	return (i + 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}
