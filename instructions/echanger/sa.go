package echanger

import "push-swap/piles"

// swap a : échange les 2 premiers éléments de la pile a

func Sa() {
	if len(piles.A) >= 2 {
		piles.A[0], piles.A[1] = piles.A[1], piles.A[0]
	}
}
