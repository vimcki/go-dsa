package count

type CountSorter struct{}

func New() *CountSorter {
	return &CountSorter{}
}

func (b *CountSorter) Sort(list []int) {
	max := 0
	for _, item := range list {
		if item > max {
			max = item
		}
	}

	counts := make([]int, max+1)
	for _, item := range list {
		counts[item] += 1
	}

	sums := make([]int, max+1)
	sum := 0
	for i, item := range counts {
		sum += item
		sums[i] = sum
	}
	result := make([]int, len(list))
	for _, item := range list {
		result[sums[item-1]] = item
		sums[item-1] += 1
	}

	copy(list, result)
}

func (b *CountSorter) SortPosition(list []int, pos uint) {
	zeroes := make([]int, 0, len(list))
	ones := make([]int, 0, len(list))
	var mask int = 1 << pos
	for _, item := range list {
		if item&mask == 0 {
			zeroes = append(zeroes, item)
		} else {
			ones = append(ones, item)
		}
	}
	copy(list, zeroes)
	copy(list[len(zeroes):], ones)
}
