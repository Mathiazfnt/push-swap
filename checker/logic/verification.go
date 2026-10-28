package checker

func IsValidInstructionSequence(sequence []string) bool {
	instructions := map[string]struct{}{
		"sa": {}, "sb": {}, "ss": {},
		"pa": {}, "pb": {},
		"ra": {}, "rb": {}, "rr": {},
		"rra": {}, "rrb": {}, "rrr": {},
	}
	for _, instruction := range sequence {
		if _, ok := instructions[instruction]; !ok {
			return false
		}
	}
	return true
}
