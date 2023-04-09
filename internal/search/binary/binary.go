package binary

type BinarySearcher struct{}

func New() *BinarySearcher {
	return &BinarySearcher{}
}

func (b *BinarySearcher) Search(list []int, target int) bool {
	low := 0
	hi := len(list)
	for low < hi {
		mid := low + (hi-low)/2
		value := list[mid]
		if value == target {
			return true
		}
		if value < target {
			low = mid + 1
		}
		if value > target {
			hi = mid
		}
	}

	return false
}
