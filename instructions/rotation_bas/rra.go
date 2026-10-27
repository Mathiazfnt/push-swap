package rotation_bas

import "push-swap/piles"

// rra a : le dernier devient le premier
func Rra() {
	if len(piles.A) < 2 {
		return
	}

	temp := []int{piles.A[len(piles.A)-1]}

	for i := 0; i < len(piles.A)-1; i++ {
		temp = append(temp, piles.A[i])
	}
	piles.A = temp
}
