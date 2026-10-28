package main

// push-swap

// Calcule et affiche la plus courte liste
// d'instructions qui trie la pile

// • Entrée : liste d'entiers
// • Sortie : une instruction par ligne
// • Objectif : minimiser le nombre de coups

import (
	"fmt"
	"os"
	"push-swap/entiers"
	"push-swap/piles"
	"push-swap/push-swap/logic"
)

func main() {
	// 1. Pas d'argument -> on n'affiche rien
	if len(os.Args) < 2 {
		return
	}

	// 2. Récupérer les nombres (gère "3 2 1 0" ET 3 2 1 0)
	nombres, err := entiers.Parse(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	piles.A = nombres
	piles.B = nil

	logic.PushSwap()
}
