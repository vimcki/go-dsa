package array

import "errors"

type Array struct {
	items [100]int
	head  int
	tail  int
}

func New() *Array {
	return &Array{
		items: [100]int{},
	}
}

func (q *Array) Enqueue(item int) {
	q.items[q.tail] = item
	if q.tail == len(q.items)-1 {
		q.tail = 0
		return
	}
	q.tail++
}

func (q *Array) Dequeue() (int, error) {
	if q.head == q.tail {
		return 0, errors.New("queue is empty")
	}

	item := q.items[q.head]
	if q.head == len(q.items)-1 {
		q.head = 0
		return item, nil
	}
	q.head++
	return item, nil
}
