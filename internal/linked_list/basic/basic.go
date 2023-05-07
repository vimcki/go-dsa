package basic

import "fmt"

type node struct {
	item string
	next *node
	prev *node
}

type LinkedList struct {
	length int
	head   *node
	tail   *node
}

func New() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) Length() int {
	return ll.length
}

func (ll *LinkedList) InsertAt(item string, index int) {
	if index < 0 || index > ll.length {
		return
	}
	if index == 0 {
		ll.Prepend(item)
		return
	}
	if index == ll.length {
		ll.Append(item)
		return
	}

	current := ll.head
	for i := 0; i != index-1; i++ {
		current = current.next
	}

	newNode := &node{
		item: item,
		prev: current,
		next: current.next,
	}
	current.next.prev = newNode
	current.next = newNode

	ll.length += 1
}

func (ll *LinkedList) Remove(item string) {
	if ll.length == 0 {
		return
	}

	current := ll.head
	for i := 0; ; i++ {
		if i == ll.length {
			return
		}
		if current.item == item {
			break
		}
		current = current.next
	}

	if current == ll.head {
		ll.head = current.next
		current.next.prev = nil
		ll.length -= 1

		return
	}

	if current == ll.tail {
		ll.tail = current.prev
		current.prev.next = nil
		ll.length -= 1

		return
	}

	current.prev.next = current.next
	current.next.prev = current.prev
	ll.length -= 1
}

func (ll *LinkedList) RemoveAt(index int) {
	if index < 0 || index >= ll.length {
		return
	}
	if ll.length == 0 {
		return
	}

	current := ll.head
	for i := 0; i != index; i++ {
		current = current.next
	}

	if current == ll.head {
		ll.head = current.next
		current.next.prev = nil
		ll.length -= 1

		return
	}

	if current == ll.tail {
		ll.tail = current.prev
		current.prev.next = nil
		ll.length -= 1

		return
	}

	current.prev.next = current.next
	current.next.prev = current.prev
	ll.length -= 1
}

func (ll *LinkedList) Append(item string) {
	if ll.length == 0 {
		node := &node{
			item: item,
		}
		ll.tail = node
		ll.head = node

		ll.length += 1

		return
	}

	tail := ll.tail
	tail.next = &node{
		prev: tail,
		item: item,
	}
	ll.tail = tail.next

	ll.length += 1
}

func (ll *LinkedList) Prepend(item string) {
	if ll.length == 0 {
		node := &node{
			item: item,
		}
		ll.tail = node
		ll.head = node

		ll.length += 1
	}

	head := ll.head
	head.prev = &node{
		next: head,
		item: item,
	}

	ll.head = head.prev

	ll.length += 1
}

func (ll *LinkedList) Get(index int) (string, error) {
	if index < 0 || index >= ll.length {
		return "", nil
	}
	current := ll.head
	for i := 0; i != index; i++ {
		current = current.next
	}
	return current.item, nil
}

func (ll *LinkedList) print() {
	rep := fmt.Sprint(ll.length) + ":"
	current := ll.head
	for current != nil {
		rep += current.item + " "
		current = current.next
	}
	fmt.Println(rep)
}
