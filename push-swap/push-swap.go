package main

import (
	"push-swap/push-swap/gestion_cas"
)

func Push_swap(a []int, b []int) {
	if len(b) != 0 {
		return
	} else if EstTriee(a) {
		return
	} else if len(a) == 2 {
		gestion_cas.Cas2(a, b)
		return
	} else if len(a) == 3 {
		gestion_cas.Cas3(a, b)
		return
	} else if len(a) <= 5 {
		gestion_cas.Cas5(a, b)
	}else{
		gestion_cas.Cas_100(a, b)
	}
}

func EstTriee(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i-1] > a[i] {
			return false
		}
	}
	return true
}
