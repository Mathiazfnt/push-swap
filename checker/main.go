package main

// checker

// Lit des instructions, les exécute et
// vérifie le résultat

// • Entrée : entiers + instructions (stdin)
// • Sortie : OK ou KO
// • OK = a triée et b vide

import (
	"bufio"
	"fmt"
	"os"
	checkerlogic "push-swap/checker/logic"
	"push-swap/entiers"
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

	// 3. Lire les instructions sur stdin
	var instructions []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	if !checkerlogic.IsValidInstructionSequence(instructions) {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	if checkerlogic.Check(nombres, instructions) {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
