package main

import (
	"github.com/vimcki/go-dsa/internal/binarytree"
	findMaxSubarray "github.com/vimcki/go-dsa/internal/find_max_subarray"
	graphsearch "github.com/vimcki/go-dsa/internal/graph_search"
	"github.com/vimcki/go-dsa/internal/hashmap"
	linkedlist "github.com/vimcki/go-dsa/internal/linked_list"
	priorityqueue "github.com/vimcki/go-dsa/internal/priority_queue"
	"github.com/vimcki/go-dsa/internal/queue"
	"github.com/vimcki/go-dsa/internal/search"
	"github.com/vimcki/go-dsa/internal/sorting"
	"github.com/vimcki/go-dsa/internal/stack"
)

func main() {
	search.Test()
	sorting.Test()
	findMaxSubarray.Test()
	priorityqueue.Test()
	stack.Test()
	queue.Test()
	linkedlist.Test()
	hashmap.Test()
	binarytree.Test()
	graphsearch.Test()
}
