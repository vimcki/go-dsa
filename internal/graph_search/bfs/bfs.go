package bfs

import (
	"container/heap"
	"errors"
	"math"
)

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

type PriorityQueue []*Item

type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (b *BFS) Djikstra(from, to int) ([]int, error) {
	if from < 0 || from >= len(b.Nodes) || to < 0 || to >= len(b.Nodes) {
		return nil, errors.New("invalid node index")
	}

	dist := make([]int, len(b.Nodes))
	prev := make([]int, len(b.Nodes))
	for i := range dist {
		dist[i] = math.MaxInt32
		prev[i] = -1
	}
	dist[from] = 0

	pq := make(PriorityQueue, 1)
	pq[0] = &Item{value: from, priority: 0, index: 0}
	heap.Init(&pq)

	for len(pq) > 0 {
		u := heap.Pop(&pq).(*Item).value
		if u == to {
			path := []int{}
			for u != -1 {
				path = append([]int{u}, path...)
				u = prev[u]
			}
			return path, nil
		}
		for _, edge := range b.Nodes[u] {
			alt := dist[u] + edge.Weight
			if alt < dist[edge.To] {
				dist[edge.To] = alt
				prev[edge.To] = u
				heap.Push(&pq, &Item{value: edge.To, priority: alt})
			}
		}
	}

	return nil, errors.New("no path found")
}
