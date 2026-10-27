package main

import "testing"

func TestChecker(t *testing.T) {
	tests := []struct {
		name         string
		entiers      []int
		instructions []string
		want         bool
	}{
		{
			name:         "sorts stack A",
			entiers:      []int{2, 1},
			instructions: []string{"sa"},
			want:         true,
		},
		{
			name:         "rejects unknown instruction",
			entiers:      []int{1, 2},
			instructions: []string{"invalid"},
			want:         false,
		},
		{
			name:    "executes every known instruction",
			entiers: []int{3, 2, 1},
			instructions: []string{
				"sa", "sb", "ss", "pa", "pb", "ra",
				"rb", "rr", "rra", "rrb", "rrr",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Checker(tt.entiers, tt.instructions); got != tt.want {
				t.Errorf("Checker(%v, %v) = %t, want %t",
					tt.entiers, tt.instructions, got, tt.want)
			}
		})
	}
}

func TestVerifier(t *testing.T) {
	if !Verifier([]string{"sa", "sb", "ss", "pa", "pb", "ra", "rb", "rr", "rra", "rrb", "rrr"}) {
		t.Error("Verifier rejected a valid instruction")
	}
	if Verifier([]string{"sa", "invalid"}) {
		t.Error("Verifier accepted an unknown instruction")
	}
}
