package main

func Verifier(tab []string) bool {
	instructions := []string{
		"sa", "sb", "ss",
		"pa", "pb",
		"ra", "rb", "rr",
		"rra", "rrb", "rrr",
	}
	for _, elt := range tab {
		trouve := false
		for _, instruction := range instructions {
			if elt == instruction {
				trouve = true
				break
			}
		}
		if !trouve {
			return false
		}
	}
	return true
}
