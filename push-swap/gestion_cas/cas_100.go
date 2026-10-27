package gestion_cas

import (
	"push-swap/instructions/pousser"
	"push-swap/instructions/rotation_haut"
)

func chiffre(carte, tour int) int {
	return (carte >> tour) & 1
}

func Cas_100(a []int, b []int) {
	// Si A est déjà rangée, on ne fait rien.

	n := len(a)

	// 7 chiffres binaires suffisent pour les nombres de 0 à 99.
	for tour := 0; tour < 7; tour++ {
		for k := 0; k < n; k++ {
			// On regarde le chiffre binaire de la carte au sommet.
			if chiffre(a[0], tour) == 0 {
				pousser.Pb()
				println("pb")
			} else {
				rotation_haut.Ra()
				println("ra")
			}
		}

		// On remet toutes les cartes de B sur A.
		for len(b) > 0 {
			pousser.Pa()
			println("pa")
		}
	}
}