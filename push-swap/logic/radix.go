package logic

import (
	"fmt"
	"math/bits"
	"sort"

	"push-swap/instructions/pousser"
	"push-swap/instructions/rotation_haut"
	"push-swap/piles"
)

func radixSort() {
	sorted := append([]int(nil), piles.A...)
	sort.Ints(sorted)
	ranks := make(map[int]int, len(sorted))
	for rank, value := range sorted {
		ranks[value] = rank
	}

	for bit := 0; bit < bits.Len(uint(len(piles.A)-1)); bit++ {
		count := len(piles.A)
		for i := 0; i < count; i++ {
			if (ranks[piles.A[0]]>>bit)&1 == 0 {
				pousser.Pb()
				fmt.Println("pb")
			} else {
				rotation_haut.Ra()
				fmt.Println("ra")
			}
		}

		for len(piles.B) > 0 {
			pousser.Pa()
			fmt.Println("pa")
		}
	}
}
