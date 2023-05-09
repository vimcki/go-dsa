package binarytree

import (
	"fmt"

	"github.com/vimcki/go-dsa/internal/binarytree/basic"
)

type binaryTree interface {
	Search(value int) bool
	Insert(value int)
	Delete(value int)
}

func Test() {
	bt := basic.New()
	test("Base binary tree", bt)
}

func test(name string, bt binaryTree) {
	bt.Insert(50)
	bt.Insert(30)
	bt.Insert(20)
	bt.Insert(40)
	bt.Insert(70)
	bt.Insert(60)
	bt.Insert(80)

	if !bt.Search(20) || !bt.Search(30) || !bt.Search(50) || !bt.Search(70) || !bt.Search(80) {
		fmt.Println(name + " FAIL")
		return
	}

	bt.Delete(20)
	bt.Delete(30)
	bt.Delete(50)

	if bt.Search(20) || bt.Search(30) || bt.Search(50) {
		fmt.Println(name + "Delete FAIL")
	}

	fmt.Println(name + " was successful")
}
