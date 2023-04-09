package linear

type LinearSearcher struct{}

func New() *LinearSearcher {
	return &LinearSearcher{}
}

func (l *LinearSearcher) Search(list []int, target int) bool {
	for _, value := range list {
		if value == target {
			return true
		}
	}
	return false
}
