package entiers

import (
	"errors"
	"strconv"
	"strings"
)

func Parse(args []string) ([]int, error) {
	// "3 2 1 0" ou 3 2 1 0 -> ["3","2","1","0"]
	tab := strings.Fields(strings.Join(args, " "))

	entiers := make([]int, 0, len(tab))
	seen := make(map[int]bool)

	for _, f := range tab {
		n, err := strconv.Atoi(f)
		if err != nil {
			return nil, errors.New("not an integer")
		}
		if seen[n] {
			return nil, errors.New("duplicate")
		}
		seen[n] = true
		entiers = append(entiers, n)
	}
	return entiers, nil
}
