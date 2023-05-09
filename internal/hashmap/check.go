package hashmap

import (
	"fmt"

	"github.com/vimcki/go-dsa/internal/hashmap/openaddr"
)

type hashMap interface {
	Insert(key, value string)
	Get(key string) (string, error)
	Delete(key string)
}

func Test() {
	hm := openaddr.New()
	test("Open addressing", hm)
}

func test(name string, hm hashMap) {
	hm.Insert("one", "1")
	hm.Insert("two", "2")
	hm.Insert("three", "3")
	hm.Insert("four", "4")

	value, err := hm.Get("one")
	if err != nil || value != "1" {
		fmt.Println(name + " Get FAIL for key 'one'")
		return
	}

	value, err = hm.Get("two")
	if err != nil || value != "2" {
		fmt.Println(name + " Get FAIL for key 'two'")
		return
	}

	value, err = hm.Get("three")
	if err != nil || value != "3" {
		fmt.Println(name + " Get FAIL for key 'three'")
		return
	}

	value, err = hm.Get("four")
	if err != nil || value != "4" {
		fmt.Println(name + " Get FAIL for key 'four'")
		return
	}

	hm.Delete("one")
	hm.Delete("two")
	hm.Delete("three")

	_, err = hm.Get("one")
	if err == nil {
		fmt.Println(name + " Delete FAIL for key 'one'")
		return
	}

	_, err = hm.Get("two")
	if err == nil {
		fmt.Println(name + " Delete FAIL for key 'two'")
		return
	}

	_, err = hm.Get("three")
	if err == nil {
		fmt.Println(name + " Delete FAIL for key 'three'")
		return
	}

	fmt.Println(name + " was successful")
}
