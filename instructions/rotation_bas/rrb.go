package rotation_bas

import "push-swap/piles"

// rrb b : le dernier devient le premier
func Rrb() {
	if len(piles.B) < 2 {
		return
	}

	temp := []int{piles.B[len(piles.B)-1]}

	for i := 0; i < len(piles.B)-1; i++ {
		temp = append(temp, piles.B[i])
	}
	piles.B = temp
}
