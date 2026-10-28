package checker

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
			name:    "accepts already sorted stack",
			entiers: []int{1, 2, 3},
			want:    true,
		},
		{
			name:         "rejects unsorted stack",
			entiers:      []int{3, 1, 2},
			instructions: []string{"sa"},
			want:         false,
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
			if got := Check(tt.entiers, tt.instructions); got != tt.want {
				t.Errorf("Check(%v, %v) = %t, want %t",
					tt.entiers, tt.instructions, got, tt.want)
			}
		})
	}
}

func TestIsValidInstructionSequence(t *testing.T) {
	if !IsValidInstructionSequence([]string{"sa", "sb", "ss", "pa", "pb", "ra", "rb", "rr", "rra", "rrb", "rrr"}) {
		t.Error("IsValidInstructionSequence rejected a valid instruction")
	}
	if IsValidInstructionSequence([]string{"sa", "invalid"}) {
		t.Error("IsValidInstructionSequence accepted an unknown instruction")
	}
	if IsValidInstructionSequence([]string{" sa "}) {
		t.Error("IsValidInstructionSequence accepted an instruction with surrounding spaces")
	}
	if IsValidInstructionSequence([]string{""}) {
		t.Error("IsValidInstructionSequence accepted an empty instruction")
	}
}
