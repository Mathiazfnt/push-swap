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
	"push-swap/entiers"
	"strings"
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
		ligne := strings.TrimSpace(scanner.Text())
		if ligne == "" {
			continue
		}
		instructions = append(instructions, ligne)
	}

	// ICI : vérifier que toutes les instructions sont connues
	if !Verifier(instructions) {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	if Checker(nombres, instructions) {
		print("OK")
	} else {
		print("KO")
	}
}
