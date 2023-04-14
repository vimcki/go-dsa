package merge

import "fmt"

type MergeSorter struct{}

func New() *MergeSorter {
	return &MergeSorter{}
}

func (b *MergeSorter) Sort(list []int) {
	sort(list, 0, len(list))
}

func sort(list []int, low, hi int) {
	if hi-low <= 1 {
		return
	}
	mid := low + (hi-low)/2
	sort(list, low, mid)
	sort(list, mid, hi)
	merge(list, low, mid, hi)
}

func merge(list []int, low, mid, hi int) {
	fmt.Println(list, low, mid, hi)
	left := make([]int, mid-low)
	right := make([]int, hi-mid)
	for i := 0; i < mid-low; i++ {
		left[i] = list[low+i]
	}
	for i := 0; i < hi-mid; i++ {
		right[i] = list[mid+i]
	}

	fmt.Println(left, right)
	i := 0
	j := 0
	for i+j < hi-low {
		fmt.Println(i, j)
		if i == len(left) {
			fmt.Println("finishing right")
			for ; j < len(right); j++ {
				list[low+i+j] = right[j]
				return
			}
		}
		if j == len(right) {
			fmt.Println("finishing left")
			for ; i < len(left); i++ {
				list[low+i+j] = left[i]
				return
			}
		}
		leftVal := left[i]
		rightVal := right[j]
		if leftVal > rightVal {
			fmt.Println("indisde j", list, low, i, j, right[j])
			list[low+i+j] = right[j]
			j++
		} else {
			fmt.Println("indisde i", list, low, i, j, left[i])
			list[low+i+j] = left[i]
			i++
		}
	}
}
