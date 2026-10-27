package pousser

import (
	"reflect"
	"testing"

	"push-swap/piles"
)

func TestPa(t *testing.T) {
	piles.A = []int{1, 2}
	piles.B = []int{3, 4}

	Pa()

	wantA := []int{3, 1, 2}
	if !reflect.DeepEqual(piles.A, wantA) {
		t.Errorf("Pa() pile A = %v, want %v", piles.A, wantA)
	}
	wantB := []int{4}
	if !reflect.DeepEqual(piles.B, wantB) {
		t.Errorf("Pa() pile B = %v, want %v", piles.B, wantB)
	}
}

func TestPaWithEmptyB(t *testing.T) {
	piles.A = []int{1, 2}
	piles.B = []int{}

	Pa()

	wantA := []int{1, 2}
	if !reflect.DeepEqual(piles.A, wantA) {
		t.Errorf("Pa() pile A = %v, want %v", piles.A, wantA)
	}
	wantB := []int{}
	if !reflect.DeepEqual(piles.B, wantB) {
		t.Errorf("Pa() pile B = %v, want %v", piles.B, wantB)
	}
}

func TestPb(t *testing.T) {
	piles.A = []int{1, 2}
	piles.B = []int{3, 4}

	Pb()

	wantA := []int{2}
	if !reflect.DeepEqual(piles.A, wantA) {
		t.Errorf("Pb() pile A = %v, want %v", piles.A, wantA)
	}
	wantB := []int{1, 3, 4}
	if !reflect.DeepEqual(piles.B, wantB) {
		t.Errorf("Pb() pile B = %v, want %v", piles.B, wantB)
	}
}

func TestPbWithEmptyA(t *testing.T) {
	piles.A = []int{}
	piles.B = []int{3, 4}

	Pb()

	wantA := []int{}
	if !reflect.DeepEqual(piles.A, wantA) {
		t.Errorf("Pb() pile A = %v, want %v", piles.A, wantA)
	}
	wantB := []int{3, 4}
	if !reflect.DeepEqual(piles.B, wantB) {
		t.Errorf("Pb() pile B = %v, want %v", piles.B, wantB)
	}
}
