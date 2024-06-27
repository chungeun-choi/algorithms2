package main

import (
	"fmt"
)

// modExp는 base^exp % mod를 계산하는 함수입니다.
func modExp(base, exp, mod int) int {
	if exp == 0 {
		return 1
	}
	half := modExp(base, exp/2, mod)
	half = (half * half) % mod
	if exp%2 != 0 {
		half = (half * base) % mod
	}
	return half
}

func baek1629() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	fmt.Println(modExp(a, b, c))
}
