package logic

import (
	"testing"

	"push-swap/piles"
)

func TestPushSwap(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		want []int
	}{
		{
			name: "already sorted",
			a:    []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "reverse pair",
			a:    []int{2, 1},
			want: []int{1, 2},
		},
		{
			name: "three values",
			a:    []int{3, 1, 2},
			want: []int{1, 2, 3},
		},
		{
			name: "four values",
			a:    []int{4, 1, 3, 2},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "five values",
			a:    []int{5, 1, 4, 2, 3},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "six values",
			a:    []int{6, 3, 1, 5, 2, 4},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "seven values",
			a:    []int{7, 4, 1, 6, 3, 5, 2},
			want: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name: "radix with negative and large values",
			a:    []int{1000, -20, 7, -300, 42, 0, -500, 35},
			want: []int{-500, -300, -20, 0, 7, 35, 42, 1000},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			piles.A = append([]int(nil), tt.a...)
			piles.B = nil

			PushSwap()

			if len(piles.B) != 0 {
				t.Fatalf("PushSwap() left B=%v, want an empty pile", piles.B)
			}
			if !equalInts(piles.A, tt.want) {
				t.Errorf("PushSwap() left A=%v, want %v", piles.A, tt.want)
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
