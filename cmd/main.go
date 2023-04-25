package main

import (
	findMaxSubarray "github.com/vimcki/go-dsa/internal/find_max_subarray"
	priorityqueue "github.com/vimcki/go-dsa/internal/priority_queue"
	"github.com/vimcki/go-dsa/internal/search"
	"github.com/vimcki/go-dsa/internal/sorting"
)

func main() {
	search.Test()
	sorting.Test()
	findMaxSubarray.Test()
	priorityqueue.Test()
}
