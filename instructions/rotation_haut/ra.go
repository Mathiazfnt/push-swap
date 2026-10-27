package rotation_haut

import "push-swap/piles"

// ra a : 1er devient dernier

func Ra() {
	if len(piles.A) < 2 {
		return
	}

	temp := piles.A[1:]
	temp = append(temp, piles.A[0])
	piles.A = temp
}
