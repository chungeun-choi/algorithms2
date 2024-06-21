package main

import (
	"fmt"
)

// backtrack 함수는 현재 선택된 숫자들의 조합을 관리합니다.
func backtrack(N, M, start int, path []int) {
	// 선택한 숫자의 개수가 M과 같으면, 해당 조합을 출력합니다.
	if len(path) == M {
		for _, num := range path {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
		return
	}

	// start부터 N까지의 숫자를 탐색합니다.
	for i := start; i <= N; i++ {
		// 새로운 숫자를 path에 추가하고, 재귀 호출합니다.
		backtrack(N, M, i+1, append(path, i))
	}
}

// solve 함수는 초기 백트래킹을 시작하는 함수입니다.
func sol(N, M int) {
	// 백트래킹 시작: 1부터 시작하며 빈 path로 시작합니다.
	backtrack(N, M, 1, []int{})
}

func bacek15650() {
	var N, M int
	fmt.Scan(&N, &M) // 사용자로부터 N과 M을 입력받습니다.
	sol(N, M)        // 주어진 N과 M으로 조합을 생성합니다.
}
