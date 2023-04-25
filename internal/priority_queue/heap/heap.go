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
	list      []int
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

func (h *Heap) Insert(v int) {
	h.list = append(h.list, v)
	i := len(h.list) - 1
	if i == 0 {
		return
	}
	h.heapifyUp(i)
}

func (h *Heap) heapifyUp(i int) {
	for {
		i = parent(i)
		h.heapify(i)
		if i == 0 {
			break
		}
	}
}

func (h *Heap) ExtractNext() int {
	heapSize := len(h.list)
	if heapSize == 0 {
		return -1
	}
	if heapSize == 1 {
		result := h.list[0]
		h.list = []int{}
		return result
	}
	result := h.list[0]
	h.list[0] = h.list[heapSize-1]
	h.list = h.list[:heapSize-1]
	h.heapify(0)
	return result
}

func (h *Heap) IncreaseKeyBy(key int, value int) {
	for i, v := range h.list {
		if v == key {
			h.list[i] = key + value
			h.heapifyUp(i)
			break
		}
	}
}

func (h *Heap) heapify(i int) {
	heapSize := len(h.list)
	if i >= heapSize {
		return
	}

	left := left(i)
	right := right(i)

	if right >= heapSize && left >= heapSize {
		return
	}

	val := h.list[i]
	var leftVal int

	if left < heapSize {
		leftVal = h.list[left]
	} else {
		leftVal = h.edgeValue
	}

	var rightVal int
	if right < heapSize {
		rightVal = h.list[right]
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
		h.list[dest] = val
		h.list[i] = destVal
		h.heapify(dest)
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
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}
