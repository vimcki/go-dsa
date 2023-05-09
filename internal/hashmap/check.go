package hashmap

import "github.com/vimcki/go-dsa/internal/hashmap/openaddr"

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
}
