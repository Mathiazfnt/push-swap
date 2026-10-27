package echanger

import "push-swap/piles"

// swap b : échange les 2 premiers éléments de la pile b

func Sb() {
	if len(piles.B) >= 2 {
		piles.B[0], piles.B[1] = piles.B[1], piles.B[0]
	}
}
