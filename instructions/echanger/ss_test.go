package echanger

import (
	"reflect"
	"testing"

	"push-swap/piles"
)

func TestSs(t *testing.T) {
	piles.A = []int{3, 1, 2}
	piles.B = []int{4, 2, 5}

	Ss()

	wantA := []int{1, 3, 2}
	if !reflect.DeepEqual(piles.A, wantA) {
		t.Errorf("Ss() pile A = %v, want %v", piles.A, wantA)
	}

	wantB := []int{2, 4, 5}
	if !reflect.DeepEqual(piles.B, wantB) {
		t.Errorf("Ss() pile B = %v, want %v", piles.B, wantB)
	}
}

func TestSsWithFewerThanTwoElements(t *testing.T) {
	piles.A = []int{1}
	piles.B = []int{}

	Ss()

	if !reflect.DeepEqual(piles.A, []int{1}) {
		t.Errorf("pile A = %v, want [1]", piles.A)
	}
	if !reflect.DeepEqual(piles.B, []int{}) {
		t.Errorf("pile B = %v, want []", piles.B)
	}
}
