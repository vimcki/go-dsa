package linkedlist

import (
	"errors"
	"fmt"

	"github.com/vimcki/go-dsa/internal/linked_list/basic"
)

type linkedList interface {
	Length() int
	InsertAt(item string, index int)
	Remove(item string)
	RemoveAt(index int)
	Append(item string)
	Prepend(item string)
	Get(index int) (string, error)
}

func Test() {
	ll := basic.New()
	test("Base linked list", ll)
}

func test(name string, ll linkedList) {
	err := check("Empty list", ll, []string{})
	if err != nil {
		fmt.Println(err)
		return
	}

	ll.Append("A")
	ll.Append("B")
	ll.Append("C")
	err = check("Append", ll, []string{"A", "B", "C"})
	if err != nil {
		fmt.Println(err)
		return
	}

	ll.Prepend("D")
	ll.Prepend("E")
	ll.Prepend("F")
	err = check("Prepend", ll, []string{"F", "E", "D", "A", "B", "C"})
	if err != nil {
		fmt.Println(err)
		return
	}

	ll.InsertAt("G", 2)
	err = check("InsertAt", ll, []string{"F", "E", "G", "D", "A", "B", "C"})
	if err != nil {
		fmt.Println(err)
		return
	}

	ll.Remove("D")
	ll.Remove("F")
	ll.Remove("C")
	ll.Remove("Z")
	err = check("Remove", ll, []string{"E", "G", "A", "B"})
	if err != nil {
		fmt.Println(err)
		return
	}

	ll.Append("Z")
	ll.Append("X")
	ll.Append("C")
	err = check("Append 2", ll, []string{"E", "G", "A", "B", "Z", "X", "C"})
	if err != nil {
		fmt.Println(err)
		return
	}

	ll.RemoveAt(ll.Length() - 1)
	ll.RemoveAt(0)
	ll.RemoveAt(2)
	err = check("RemoveAt", ll, []string{"G", "A", "Z", "X"})
	if err != nil {
		fmt.Println(err)
		return
	}

	ll.InsertAt("C", 1)
	check("InsertAt: middle element", ll, []string{"G", "A", "Z", "X"})

	fmt.Printf("%s was successful\n", name)
}

func check(name string, ll linkedList, expected []string) error {
	if ll.Length() != len(expected) {
		return errors.New(name + " length mismatch")
	}

	for i, val := range expected {
		if item, err := ll.Get(i); err != nil || item != val {
			if err != nil {
				return err
			}
			return fmt.Errorf("%s value mismatch want %v, got  %v at %v", name, val, item, i)
		}
	}
	return nil
}
