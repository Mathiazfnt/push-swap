package rotation_haut

import "push-swap/piles"

// rb b : 1er devient dernier

func Rb() {
	if len(piles.B) < 2 {
		return
	}

	temp := piles.B[1:]
	temp = append(temp, piles.B[0])
	piles.B = temp
}
