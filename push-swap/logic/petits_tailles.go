package logic

import (
	"fmt"

	"push-swap/instructions/echanger"
	"push-swap/instructions/pousser"
	"push-swap/instructions/rotation_bas"
	"push-swap/instructions/rotation_haut"
	"push-swap/piles"
)

func sortTwo() {
	if piles.A[0] > piles.A[1] {
		echanger.Sa()
		fmt.Println("sa")
	}
}

func sortThree() {
	max := piles.A[0]
	for _, value := range piles.A {
		if max < value {
			max = value
		}
	}

	if max == piles.A[0] {
		rotation_haut.Ra()
		fmt.Println("ra")
	} else if max == piles.A[1] {
		rotation_bas.Rra()
		fmt.Println("rra")
	}
	if piles.A[0] > piles.A[1] {
		echanger.Sa()
		fmt.Println("sa")
	}
}

func minPosition() int {
	pos := 0
	for i := 1; i < len(piles.A); i++ {
		if piles.A[i] < piles.A[pos] {
			pos = i
		}
	}
	return pos
}

func sortUpToSeven() {
	for len(piles.A) > 3 {
		for position := minPosition(); position != 0; position = minPosition() {
			if position == len(piles.A)-1 {
				rotation_bas.Rra()
				fmt.Println("rra")
			} else if position == 1 {
				echanger.Sa()
				fmt.Println("sa")
			} else {
				rotation_haut.Ra()
				fmt.Println("ra")
			}
		}
		pousser.Pb()
		fmt.Println("pb")
	}
	sortThree()
	for len(piles.B) > 0 {
		pousser.Pa()
		fmt.Println("pa")
	}
}
