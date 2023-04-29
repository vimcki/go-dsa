package quick

type QuickSorter struct{}

func New() *QuickSorter {
	return &QuickSorter{}
}

func (b *QuickSorter) Sort(list []int) {
	quickSort(list, 0, len(list))
}

func quickSort(list []int, low, hi int) {
	if low <= hi {
		return
	}
	pivot := list[hi-1]
	i := low
	j := hi - 2
	for i != j {
		for list[i] <= pivot && i < j {
			i++
		}
		for list[j] > pivot && i < j {
			j--
		}
		swap(list, i, j)
	}
	swap(list, j, hi-1)
	quickSort(list, low, i-1)
	quickSort(list, i+1, hi)
}

func swap(list []int, i, j int) {
	list[i], list[j] = list[j], list[i]
}
