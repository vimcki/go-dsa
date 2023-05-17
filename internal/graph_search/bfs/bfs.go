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
	visited := make([]bool, len(b.Nodes))

	visited[from] = true

	visitedFrom := make([]int, len(b.Nodes))

	for i := range visitedFrom {
		visitedFrom[i] = -1
	}
	nextNodes := []int{from}
	var current int
	for len(nextNodes) > 0 {
		current, nextNodes = nextNodes[0], nextNodes[1:]
		if current == to {
			break
		}
		for _, edge := range b.Nodes[current] {
			if visited[edge.To] {
				continue
			}
			nextNodes = append(nextNodes, edge.To)
			visited[edge.To] = true
			visitedFrom[edge.To] = current
		}

		visited[current] = true
	}

	current = to
	path := []int{current}
	for current != from {
		current = visitedFrom[current]
		path = append(path, current)
	}
	path = reverse(path)

	return path, nil
}

func reverse(path []int) []int {
	reversed := make([]int, len(path))
	for i := range path {
		reversed[i] = path[len(path)-1-i]
	}
	return reversed
}
