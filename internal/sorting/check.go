package sorting

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/vimcki/go-dsa/internal/sorting/bubble"
	"github.com/vimcki/go-dsa/internal/sorting/count"
	"github.com/vimcki/go-dsa/internal/sorting/heap"
	"github.com/vimcki/go-dsa/internal/sorting/merge"
	"github.com/vimcki/go-dsa/internal/sorting/quick"
)

type testCase struct {
	name string
	list []int
}

type sorter interface {
	Sort(list []int)
}

func Test() {
	b := bubble.New()
	test("Bubble Sort", b)
	m := merge.New()
	test("Merge sort", m)
	h := heap.New(heap.Max)
	test("Heap sort", h)
	q := quick.New()
	test("Quick sort", q)
	c := count.New()
	test("Count sort", c)
}

func test(name string, s sorter, tests ...string) {
	testCases := []testCase{
		{
			name: "s1",
			list: []int{1, 2, 3, 4, 5},
		},
		{
			name: "s2",
			list: []int{5, 5, 4, 3, 2, 1},
		},
		{
			name: "s3",
			list: []int{},
		},
		{
			name: "s4",
			list: []int{
				31, 1, 81, 100, 34, 95, 12, 28, 84, 51, 71, 57, 31, 42, 22,
				63, 67, 64, 70, 24, 89, 69, 12, 60, 62, 57, 75, 15, 79, 10,
				33, 7, 97, 92, 28, 56, 35, 66, 51, 46, 12, 96, 63, 72, 16, 96,
			},
		},
	}
	failed := false
	for _, t := range testCases {
		if len(tests) != 0 {
			if !contains(tests, t.name) {
				continue
			}
		}
		want := make([]int, len(t.list))
		got := make([]int, len(t.list))
		copy(want, t.list)
		copy(got, t.list)
		sort.Ints(want)
		s.Sort(got)
		if !reflect.DeepEqual(want, got) {
			fmt.Printf(
				"%s %s test failed, got: %v, want: %v for list: %v\n",
				name,
				t.name,
				got,
				want,
				t.list,
			)
			failed = true
		}
	}
	if !failed {
		fmt.Printf("%s was successful\n", name)
	}
}

func contains(list []string, target string) bool {
	for _, value := range list {
		if target == value {
			return true
		}
	}
	return false
}
