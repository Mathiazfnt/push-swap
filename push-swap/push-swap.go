package main

import (
	"fmt"

	"push-swap/checker"
	"push-swap/instructions/echanger"
	"push-swap/instructions/rotation_bas"
	"push-swap/instructions/rotation_haut"
)

func Push_swap(a []int, b []int) {
	if checker.EstTriee(a) {
		return
	}
	if len(a) == 2 {
		cas_2(a, b)
		return
	} else if len(a) == 3 {
		cas_3(a, b)
		return
	}
}

func cas_2(a []int, b []int) {
	if a[0] > a[1] {
		echanger.Sa()
		fmt.Println("sa")
	}
}

func cas_3(a []int, b []int) { 
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
