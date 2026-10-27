package pousser

import "push-swap/piles"

// Pa prend le premier élément au sommet de la pile b
// et le place au sommet de la pile a.

func Pa() {
	if len(piles.B) > 0 {
		piles.A = append([]int{piles.B[0]}, piles.A...)
		piles.B = piles.B[1:]
	}
}
