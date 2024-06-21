package main

import (
	"bufio"
	"fmt"
	"os"
)

func baek9663() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)
	// 내부 변수 선언 및 초기화
	var (
		issued1 = make([]bool, n)
		issued2 = make([]bool, 2*n-1)
		issued3 = make([]bool, 2*n-1)
		cnt     int
	)

	// 내부 함수 val 정의
	var val func(cur int)
	val = func(cur int) {
		if cur == n {
			cnt++
			return
		}

		for i := 0; i < n; i++ {
			if issued1[i] || issued2[cur+i] || issued3[cur-i+(n-1)] {
				continue
			}
			issued1[i] = true
			issued2[cur+i] = true
			issued3[cur-i+(n-1)] = true
			val(cur + 1)
			issued1[i] = false
			issued2[cur+i] = false
			issued3[cur-i+(n-1)] = false
		}
	}

	val(0)
	fmt.Println(cnt)
}
