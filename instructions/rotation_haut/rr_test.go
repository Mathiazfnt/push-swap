package rotation_haut

import (
	"testing"

	"push-swap/piles"
)

func TestRr(t *testing.T) {
	piles.A = []int{1, 2, 3}
	piles.B = []int{4, 5, 6}

	Rr()

	wantA := []int{2, 3, 1}
	wantB := []int{5, 6, 4}
	if !equalInts(piles.A, wantA) {
		t.Errorf("Rr() pile A = %v, want %v", piles.A, wantA)
	}
	if !equalInts(piles.B, wantB) {
		t.Errorf("Rr() pile B = %v, want %v", piles.B, wantB)
	}
}

func TestRrWithFewerThanTwoElements(t *testing.T) {
	tests := []struct {
		name  string
		a     []int
		b     []int
		wantA []int
		wantB []int
	}{
		{
			name:  "both piles empty",
			a:     []int{},
			b:     []int{},
			wantA: []int{},
			wantB: []int{},
		},
		{
			name:  "both piles have one element",
			a:     []int{1},
			b:     []int{2},
			wantA: []int{1},
			wantB: []int{2},
		},
		{
			name:  "only pile A can rotate",
			a:     []int{1, 2},
			b:     []int{3},
			wantA: []int{2, 1},
			wantB: []int{3},
		},
		{
			name:  "only pile B can rotate",
			a:     []int{1},
			b:     []int{2, 3},
			wantA: []int{1},
			wantB: []int{3, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			piles.A = append([]int{}, tt.a...)
			piles.B = append([]int{}, tt.b...)

			Rr()

			if !equalInts(piles.A, tt.wantA) {
				t.Errorf("Rr() pile A = %v, want %v", piles.A, tt.wantA)
			}
			if !equalInts(piles.B, tt.wantB) {
				t.Errorf("Rr() pile B = %v, want %v", piles.B, tt.wantB)
			}
		})
	}
}

func equalInts(got, want []int) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}
