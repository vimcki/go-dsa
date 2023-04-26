package priorityqueue

import (
	"fmt"

	"github.com/vimcki/go-dsa/internal/priority_queue/heap"
)

type queue interface {
	Insert(int)
	ExtractNext() int
	IncreaseKeyBy(key int, value int)
}

func Test() {
	h := heap.New(heap.Max)
	test("Heap priority queue", h)
}

func test(name string, q queue, tests ...string) {
	q.Insert(10)

	v := q.ExtractNext()
	if v != 10 {
		fmt.Printf("test %s failed on v == 10\n", name)
		return
	}

	q.Insert(10)
	q.Insert(30)
	q.Insert(15)
	q.Insert(21)
	q.Insert(1)

	v = q.ExtractNext()
	if v != 30 {
		fmt.Printf("test %s failed on v == 30\n", name)
		return
	}

	q.Insert(11)
	q.Insert(12)
	q.Insert(13)

	v = q.ExtractNext()
	if v != 21 {
		fmt.Printf("test %s failed on v == 21\n", name)
		return
	}

	q.IncreaseKeyBy(2, 20)

	v = q.ExtractNext()
	if v != 33 {
		fmt.Printf("test %s failed on v == 33\n", name)
		return
	}

	v = q.ExtractNext()
	if v != 15 {
		fmt.Printf("test %s failed on v == 15\n", name)
		return
	}
	v = q.ExtractNext()
	if v != 12 {
		fmt.Printf("test %s failed on v == 12\n", name)
		return
	}
	v = q.ExtractNext()
	if v != 11 {
		fmt.Printf("test %s failed on v == 11\n", name)
		return
	}
	v = q.ExtractNext()
	if v != 10 {
		fmt.Printf("test %s failed on v == 10\n", name)
		return
	}
	v = q.ExtractNext()
	if v != 1 {
		fmt.Printf("test %s failed on v == 1\n", name)
		return
	}

	fmt.Printf("%s was successful\n", name)
}
