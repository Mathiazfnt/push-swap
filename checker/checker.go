package checker

import (
	"push-swap/instructions/echanger"
	"push-swap/instructions/pousser"
	"push-swap/instructions/rotation_bas"
	"push-swap/instructions/rotation_haut"
	"push-swap/piles"
)

var liste_operations = map[string]func(){
	"sa":  echanger.Sa,
	"sb":  echanger.Sb,
	"ss":  echanger.Ss,
	"pa":  pousser.Pa,
	"pb":  pousser.Pb,
	"ra":  rotation_haut.Ra,
	"rb":  rotation_haut.Rb,
	"rr":  rotation_haut.Rr,
	"rra": rotation_bas.Rra,
	"rrb": rotation_bas.Rrb,
	"rrr": rotation_bas.Rrr,
}

func Checker(entiers []int, instructions []string) bool {
	piles.A = append([]int(nil), entiers...)
	piles.B = nil

	for _, instr := range instructions {
		operation, ok := liste_operations[instr]
		if !ok {
			return false
		}
		operation() // operation deviens la fonction qui effectue une instruction 
	}
	return len(piles.B) == 0 && EstTriee(piles.A)
}

func EstTriee(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i-1] > a[i] {
			return false
		}
	}
	return true
}
