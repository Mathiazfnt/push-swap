package gestion_cas

import (
	"fmt"

	"push-swap/instructions/echanger"
	"push-swap/instructions/pousser"
	"push-swap/instructions/rotation_bas"
	"push-swap/instructions/rotation_haut"
)

func Cas2(a []int, b []int) {
	if a[0] > a[1] {
		echanger.Sa()
		fmt.Println("sa")
	}
}

func Cas3(a []int, b []int) {
	max := a[0]
	for _, elt := range a {
		if max < elt {
			max = elt
		}
	}

	if max == a[0] {
		rotation_haut.Ra()
		fmt.Println("ra")
	} else if max == a[1] {
		rotation_bas.Rra()
		fmt.Println("rra")
	}
	if a[0] > a[1] {
		echanger.Sa()
		fmt.Println("sa")
	}
}

func positionDuMin(a []int) int {
	pos := 0
	for i := 1; i < len(a); i++ {
		if a[i] < a[pos] {
			pos = i
		}
	}
	return pos
}

func Cas5(a []int, b []int) {
	// 2. On place les deux plus petites cartes dans B.
	for len(a) > 3 {
		for positionDuMin(a) != 0 {
			if positionDuMin(a) == len(a)-1 {
				rotation_bas.Rra()
				println("rra")
			} else if positionDuMin(a) == 1 {
				echanger.Sa()
				println("sa")
			} else {
				rotation_haut.Ra()
				print("ra")
			}
		}
		pousser.Pb()
		println("pb")
	}
	Cas3(a, b)
	pousser.Pa()
	println("pa")
	pousser.Pa()
	println("pa")
}
