package stack

import (
	"fmt"

	"github.com/vimcki/go-dsa/internal/stack/array"
)

type lifo interface {
	Push(int)
	Pop() (int, error)
}

func Test() {
	s := array.New()
	test("Array stack", s)
}

func test(name string, s lifo) {
	s.Push(10)

	v, err := s.Pop()
	if err != nil {
		fmt.Printf("test %s failed 1 on err != nil\n", name)
		return
	}
	if v != 10 {
		fmt.Printf("test %s failed on v == 10\n", name)
		return
	}

	_, err = s.Pop()
	if err == nil {
		fmt.Printf("test %s failed on err == nil\n", name)
		return
	}

	s.Push(10)
	s.Push(30)
	s.Push(15)
	s.Push(21)
	s.Push(1)

	v, err = s.Pop()
	if err != nil {
		fmt.Printf("test %s failed 2 on err != nil\n", name)
		return
	}
	if v != 1 {
		fmt.Printf("test %s failed on v == 1\n", name)
		return
	}

	v, err = s.Pop()
	if err != nil {
		fmt.Printf("test %s failed 3 on err != nil\n", name)
		return
	}
	if v != 21 {
		fmt.Printf("test %s failed on v == 21\n", name)
		return
	}

	v, err = s.Pop()
	if err != nil {
		fmt.Printf("test %s failed 4 on err != nil\n", name)
		return
	}
	if v != 15 {
		fmt.Printf("test %s failed on v == 15\n", name)
		return
	}

	s.Push(20)

	v, err = s.Pop()
	if err != nil {
		fmt.Printf("test %s failed 5 on err != nil\n", name)
		return
	}
	if v != 20 {
		fmt.Printf("test %s failed on v == 20\n", name)
		return
	}

	v, err = s.Pop()
	if err != nil {
		fmt.Printf("test %s failed 6 on err != nil\n", name)
		return
	}
	if v != 30 {
		fmt.Printf("test %s failed on v == 30\n", name)
		return
	}

	v, err = s.Pop()
	if err != nil {
		fmt.Printf("test %s failed 7 on err != nil\n", name)
		return
	}
	if v != 10 {
		fmt.Printf("test %s failed on v == 10\n", name)
		return
	}

	_, err = s.Pop()
	if err == nil {
		fmt.Printf("test %s failed on err == nil\n", name)
		return
	}

	fmt.Printf("%s was successful\n", name)
}
