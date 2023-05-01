package quick

type QuickSorter struct{}

func New() *QuickSorter {
	return &QuickSorter{}
}

func (b *QuickSorter) Sort(list []int) {
	quickSort(list, 0, len(list)-1)
}

func quickSort(list []int, low, hi int) {
	if low >= hi {
		return
	}
	pivot := list[hi]
	i := low
	j := hi - 1

	for i <= j {
		for i <= j && list[i] < pivot {
			i++
		}
		for i <= j && list[j] > pivot {
			j--
		}
		if i <= j {
			swap(list, i, j)
			i++
			j--
		}
	}

	swap(list, i, hi)
	quickSort(list, low, i-1)
	quickSort(list, i+1, hi)
}

func swap(list []int, i, j int) {
	list[i], list[j] = list[j], list[i]
}
