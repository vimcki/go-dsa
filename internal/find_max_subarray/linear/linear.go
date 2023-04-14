package linear

type Linear struct{}

func New() *Linear {
	return &Linear{}
}

func (l *Linear) Find(prices []int) (int, int) {
	start := -1
	end := -1
	changes := make([]int, len(prices)-1)
	for i := 1; i < len(prices); i++ {
		old := prices[i-1]
		current := prices[i]
		changes[i-1] = current - old
	}

	currentStart := 0
	currentValue := 0
	value := 0

	for i, change := range changes {
		currentValue = currentValue + change
		if currentValue < 0 {
			currentStart = i + 1
			currentValue = 0
			continue
		}
		if currentValue > value {
			start = currentStart
			end = i + 1
			value = currentValue
		}
	}

	return start, end
}
