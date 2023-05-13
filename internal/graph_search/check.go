package graphsearch

import "github.com/vimcki/go-dsa/internal/graph_search/bfs"

type searcher interface {
	Search(from, to int) ([]int, error)
}

func Test() {
	b := bfs.New([][]bfs.Edge{})
	test("bfs", b)
}

func test(name string, s searcher) {
}
