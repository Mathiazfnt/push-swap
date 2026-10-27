package rotation_bas

import (
	"reflect"
	"testing"

	"push-swap/piles"
)

func TestRrr(t *testing.T) {
	piles.A = []int{1, 2, 3}
	piles.B = []int{4, 5, 6}

	Rrr()

	wantA := []int{3, 1, 2}
	wantB := []int{6, 4, 5}
	if !reflect.DeepEqual(piles.A, wantA) {
		t.Errorf("Rrr() pile A = %v, want %v", piles.A, wantA)
	}
	if !reflect.DeepEqual(piles.B, wantB) {
		t.Errorf("Rrr() pile B = %v, want %v", piles.B, wantB)
	}
}

func TestRrrWithFewerThanTwoElements(t *testing.T) {
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

			Rrr()

			if !reflect.DeepEqual(piles.A, tt.wantA) {
				t.Errorf("Rrr() pile A = %v, want %v", piles.A, tt.wantA)
			}
			if !reflect.DeepEqual(piles.B, tt.wantB) {
				t.Errorf("Rrr() pile B = %v, want %v", piles.B, tt.wantB)
			}
		})
	}
}
