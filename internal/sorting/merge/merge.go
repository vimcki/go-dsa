package merge

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
	left := make([]int, mid-low)
	right := make([]int, hi-mid)
	for i := 0; i < mid-low; i++ {
		left[i] = list[low+i]
	}
	for i := 0; i < hi-mid; i++ {
		right[i] = list[mid+i]
	}

	i := 0
	j := 0
	for i+j < hi-low {
		if i == len(left) {
			for ; j < len(right); j++ {
				list[low+i+j] = right[j]
			}
			return
		}
		if j == len(right) {
			for ; i < len(left); i++ {
				list[low+i+j] = left[i]
			}
			return
		}
		leftVal := left[i]
		rightVal := right[j]
		if leftVal > rightVal {
			list[low+i+j] = right[j]
			j++
		} else {
			list[low+i+j] = left[i]
			i++
		}
	}
}
