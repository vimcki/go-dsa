package bfs

type Edge struct {
	To     int
	Weight int
}

type BFS struct {
	Nodes [][]Edge
}

func New(nodes [][]Edge) *BFS {
	return &BFS{Nodes: nodes}
}

func (b *BFS) Search(from, to int) ([]int, error) {
	return []int{}, nil
}
