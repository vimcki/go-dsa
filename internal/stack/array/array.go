package array

import "fmt"

type Array struct {
	items [100]int
	top   int
}

func New() *Array {
	return &Array{
		items: [100]int{},
	}
}

func (q *Array) Push(item int) {
	q.items[q.top] = item
	q.top += 1
}

func (q *Array) Pop() (int, error) {
	if q.top == 0 {
		return 0, fmt.Errorf("Stack is empty")
	}
	q.top -= 1
	return q.items[q.top], nil
}
