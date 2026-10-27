package pousser

import "push-swap/piles"

// push b : prend le premier élément au sommet de la pile a
// et le place au sommet de la pile b

func Pb() {
	if len(piles.A) > 0 {
		piles.B = append([]int{piles.A[0]}, piles.B...)
		piles.A = piles.A[1:]
	}
}
