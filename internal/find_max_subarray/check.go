package findmaxsubarray

import (
	"fmt"
	"reflect"

	"github.com/vimcki/go-dsa/internal/find_max_subarray/linear"
)

type testCase struct {
	name           string
	prices         []int
	wantStartIndex int
	wantEndIndex   int
}

type finder interface {
	Find(prices []int) (int, int)
}

func Test() {
	l := linear.New()
	test("Linear Find Max Subarray", l)
}

func test(name string, f finder, tests ...string) {
	testCases := []testCase{
		{
			name:           "s1",
			prices:         []int{1, 2},
			wantStartIndex: 0,
			wantEndIndex:   1,
		},
		{
			name:           "s2",
			prices:         []int{5, 4, 3, 2, 1},
			wantStartIndex: -1,
			wantEndIndex:   -1,
		},
		{
			name: "s4",
			prices: []int{
				100, 113, 110, 85, 105, 102,
				86, 63, 81, 101, 94, 106,
				101, 79, 94, 90, 97, 12,
			},
			wantStartIndex: 7,
			wantEndIndex:   11,
		},
		{
			name: "s5",
			prices: []int{
				31, 1, 81, 100, 34, 95, 12, 28, 84, 51, 71, 57, 31, 42, 22,
				63, 67, 64, 70, 24, 89, 69, 12, 60, 62, 57, 75, 15, 79, 10,
				33, 7, 97, 92, 28, 56, 35, 66, 51, 46, 12, 96, 63, 72, 16, 96,
			},
			wantStartIndex: 1,
			wantEndIndex:   3,
		},
	}
	failed := false
	for _, t := range testCases {
		if len(tests) != 0 {
			if !contains(tests, t.name) {
				continue
			}
		}
		start, end := f.Find(t.prices)
		if !reflect.DeepEqual(t.wantStartIndex, start) {
			fmt.Printf(
				"%s %s test failed for start, got: %v, want: %v for list: %v\n",
				name,
				t.name,
				start,
				t.wantStartIndex,
				t.prices,
			)
			failed = true
		}
		if !reflect.DeepEqual(t.wantEndIndex, end) {
			fmt.Printf(
				"%s %s test failed for end, got: %v, want: %v for list: %v\n",
				name,
				t.name,
				end,
				t.wantEndIndex,
				t.prices,
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
