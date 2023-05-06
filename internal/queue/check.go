package queue

import (
	"fmt"

	"github.com/vimcki/go-dsa/internal/queue/array"
)

type fifo interface {
	Enqueue(int)
	Dequeue() (int, error)
}

func Test() {
	q := array.New()
	test("Array queue", q)
}

func test(name string, q fifo) {
	q.Enqueue(10)

	v, err := q.Dequeue()
	if err != nil {
		fmt.Printf("test %s failed on err != nil\n", name)
		return
	}
	if v != 10 {
		fmt.Printf("test %s failed on v == 10\n", name)
		return
	}

	_, err = q.Dequeue()
	if err == nil {
		fmt.Printf("test %s failed on err == nil\n", name)
		return
	}

	q.Enqueue(10)
	q.Enqueue(30)
	q.Enqueue(15)
	q.Enqueue(21)
	q.Enqueue(1)

	v, err = q.Dequeue()
	if err != nil {
		fmt.Printf("test %s failed on err != nil\n", name)
		return
	}
	if v != 10 {
		fmt.Printf("test %s failed on v == 10\n", name)
		return
	}

	q.Enqueue(11)
	q.Enqueue(12)
	q.Enqueue(13)

	v, err = q.Dequeue()
	if err != nil {
		fmt.Printf("test %s failed on err != nil\n", name)
		return
	}
	if v != 30 {
		fmt.Printf("test %s failed on v == 30\n", name)
		return
	}

	v, err = q.Dequeue()
	if err != nil {
		fmt.Printf("test %s failed on err != nil\n", name)
		return
	}
	if v != 15 {
		fmt.Printf("test %s failed on v == 15\n", name)
		return
	}

	q.Enqueue(20)

	v, err = q.Dequeue()
	if err != nil {
		fmt.Printf("test %s failed on err != nil\n", name)
		return
	}
	if v != 21 {
		fmt.Printf("test %s failed on v == 21\n", name)
		return
	}

	v, err = q.Dequeue()
	if err != nil {
		fmt.Printf("test %s failed on err != nil\n", name)
		return
	}
	if v != 1 {
		fmt.Printf("test %s failed on v == 1\n", name)
		return
	}

	v, err = q.Dequeue()
	if err != nil {
		fmt.Printf("test %s failed on err != nil\n", name)
		return
	}
	if v != 11 {
		fmt.Printf("test %s failed on v == 11\n", name)
		return
	}

	v, err = q.Dequeue()
	if err != nil {
		fmt.Printf("test %s failed on err != nil\n", name)
		return
	}
	if v != 12 {
		fmt.Printf("test %s failed on v == 12\n", name)
		return
	}

	v, err = q.Dequeue()
	if err != nil {
		fmt.Printf("test %s failed on err != nil\n", name)
		return
	}
	if v != 13 {
		fmt.Printf("test %s failed on v == 13\n", name)
		return
	}

	v, err = q.Dequeue()
	if err != nil {
		fmt.Printf("test %s failed on err != nil\n", name)
		return
	}
	if v != 20 {
		fmt.Printf("test %s failed on v == 20\n", name)
		return
	}

	fmt.Printf("%s was successful\n", name)
}
