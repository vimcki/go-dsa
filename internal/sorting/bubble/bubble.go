package bubble

type BubbleSorter struct{}

func New() *BubbleSorter {
	return &BubbleSorter{}
}

func (b *BubbleSorter) Sort(list []int) {
	for i := len(list) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if list[j] > list[i] {
				swap(list, j, i)
			}
		}
	}
}

func swap(list []int, i, j int) {
	list[i], list[j] = list[j], list[i]
}
