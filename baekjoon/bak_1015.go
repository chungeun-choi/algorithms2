package baekjoon

import (
	"fmt"
	"sort"
)

func bak1015() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	b := make([]int, n)
	copy(b, a)

	sort.Ints(b)

	p := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a[i] == b[j] {
				p[i] = j
				b[j] = -1 // 중복 방지
				break
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Print(p[i], " ")
	}
	fmt.Println()
}
