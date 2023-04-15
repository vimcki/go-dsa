package heap

import (
	"reflect"
	"testing"
)

func TestHeap_heapify(t *testing.T) {
	tests := []struct {
		name     string
		hType    Type
		list     []int
		i        int
		heapSize int
		wantList []int
	}{
		{
			name:     "1",
			hType:    Max,
			list:     []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1},
			i:        1,
			wantList: []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1},
		},
		{
			name:     "1 but with heapSize != len(list)",
			hType:    Max,
			list:     []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1, 23, 1, 4, 5},
			heapSize: 10,
			i:        1,
			wantList: []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1, 23, 1, 4, 5},
		},
		{
			name:     "2",
			hType:    Max,
			list:     []int{1, 17, 18, 8, 3, 2, 9, 2, 2, 2},
			i:        0,
			wantList: []int{18, 17, 9, 8, 3, 2, 1, 2, 2, 2},
		},
		{
			name:     "2 but Min",
			hType:    Min,
			list:     []int{-1, -17, -18, -8, -3, -2, -9, -2, -2, -2},
			i:        0,
			wantList: []int{-18, -17, -9, -8, -3, -2, -1, -2, -2, -2},
		},
		{
			name:     "nothing to do",
			hType:    Max,
			list:     []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1},
			i:        2,
			wantList: []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Heap{
				hType: tt.hType,
			}
			heapSize := len(tt.list)
			if tt.heapSize != 0 {
				heapSize = tt.heapSize
			}
			h.heapify(tt.list, heapSize, tt.i)

			if reflect.DeepEqual(tt.list, tt.wantList) {
				t.Errorf("want: %+v, got: %+v", tt.wantList, tt.list)
			}
		})
	}
}
