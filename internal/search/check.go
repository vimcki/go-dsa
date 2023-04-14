package search

import (
	"fmt"

	"github.com/vimcki/go-dsa/internal/search/binary"
	"github.com/vimcki/go-dsa/internal/search/linear"
)

type testCase struct {
	name   string
	list   []int
	target int
	want   bool
}

type searcher interface {
	Search([]int, int) bool
}

func Test() {
	ls := linear.New()
	test("Linear Search", ls)
	bs := binary.New()
	test("Binary Search", bs)
}

func test(name string, s searcher, tests ...string) {
	testCases := []testCase{
		{
			name:   "s1",
			list:   []int{},
			target: 3,
			want:   false,
		},
		{
			name:   "s2",
			list:   []int{1, 2, 3, 4, 5},
			target: 3,
			want:   true,
		},
		{
			name:   "s3",
			list:   []int{1, 5, 12, 69, 323, 664, 832},
			target: 3,
			want:   false,
		},
		{
			name:   "s4",
			list:   []int{1, 5, 12, 69, 323, 664, 832},
			target: 832,
			want:   true,
		},
		{
			name:   "s5",
			list:   []int{1, 5, 12, 69, 323, 664, 832, 1234, 5432, 8756, 9213, 11111},
			target: 1,
			want:   true,
		},
		{
			name:   "s6",
			list:   []int{1, 5, 12, 69, 323, 664, 832, 1234, 5432, 8756, 9213, 11111},
			target: 11111,
			want:   true,
		},
	}

	failed := false
	for _, t := range testCases {
		if len(tests) != 0 {
			if !contains(tests, t.name) {
				continue
			}
		}
		got := s.Search(t.list, t.target)
		if got != t.want {
			fmt.Printf(
				"%s test failed, got: %v, want: %v for list: %v and target: %v\n",
				name,
				got,
				t.want,
				t.list,
				t.target,
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
