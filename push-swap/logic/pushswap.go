package logic

import "push-swap/piles"

func PushSwap() {
	if isSorted(piles.A) {
		return
	}
	switch {
	case len(piles.A) == 2:
		sortTwo()
	case len(piles.A) == 3:
		sortThree()
	case len(piles.A) <= 7:
		sortUpToSeven()
	default:
		radixSort()
	}
}

func isSorted(values []int) bool {
	for i := 1; i < len(values); i++ {
		if values[i-1] > values[i] {
			return false
		}
	}
	return true
}
