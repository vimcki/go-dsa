package graphsearch

import (
	"fmt"
	"reflect"

	"github.com/vimcki/go-dsa/internal/graph_search/bfs"
)

type searcher interface {
	Search(from, to int) ([]int, error)
}

func Test() {
	b := bfs.New([][]bfs.Edge{
		{
			{To: 1, Weight: 1},
			{To: 3, Weight: 1},
		},

		{
			{To: 2, Weight: 2},
			{To: 4, Weight: 3},
			{To: 5, Weight: 5},
		},
		{
			{To: 5, Weight: 1},
		},
		{
			{To: 8, Weight: 1},
		},
		{
			{To: 3, Weight: 7},
			{To: 8, Weight: 2},
			{To: 5, Weight: 5},
		},
		{
			{To: 6, Weight: 7},
			{To: 10, Weight: 2},
		},
		{
			{To: 7, Weight: 1},
		},
		{
			{To: 6, Weight: 1},
		},
		{
			{To: 9, Weight: 3},
		},
		{
			{To: 11, Weight: 1},
		},
		{
			{To: 9, Weight: 2},
		},
		{},
	})
	test("bfs from 0 to 5", b.Search, 0, 5, []int{0, 1, 5})
	test("bfs from 0 to 11", b.Search, 0, 11, []int{0, 3, 8, 9, 11})
	test("djikstra from 0 to 5", b.Djikstra, 0, 5, []int{0, 1, 2, 5})
	test("djikstra from 1 to 11", b.Djikstra, 1, 11, []int{1, 2, 5, 10, 9, 11})
}

func test(name string, s func(int, int) ([]int, error), from, to int, want []int) {
	fmt.Printf("Running test: %s\n", name)
	got, err := s(from, to)
	if err != nil {
		fmt.Printf("Search returned error: %v\n", err)
		return
	}
	if !reflect.DeepEqual(got, want) {
		fmt.Printf("Search = %v, want %v\n", got, want)
	} else {
		fmt.Println("Test passed.")
	}
}
