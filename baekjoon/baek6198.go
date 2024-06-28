package baekjoon

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func baek6198() {
	// 표준 입력을 받아오기 위한 스캐너 생성
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	// 빌딩의 수 N을 읽음
	N, _ := strconv.Atoi(scanner.Text())

	// 빌딩의 높이를 저장할 slice
	buildings := make([]int, N)

	// 각 빌딩의 높이를 입력받아 저장
	for i := 0; i < N; i++ {
		scanner.Scan()
		buildings[i], _ = strconv.Atoi(scanner.Text())
	}

	// 결과를 저장할 변수
	result := 0
	// 스택을 이용해 오른쪽 빌딩들을 체크
	stack := []int{}

	// 빌딩을 순회
	for i := 0; i < N; i++ {
		// 현재 빌딩보다 높은 빌딩은 스택에서 제거
		for len(stack) > 0 && stack[len(stack)-1] <= buildings[i] {
			stack = stack[:len(stack)-1]
		}
		// 스택에 남아있는 빌딩의 수가 현재 빌딩에서 볼 수 있는 빌딩의 수
		result += len(stack)
		// 현재 빌딩을 스택에 추가
		stack = append(stack, buildings[i])
	}

	// 결과 출력
	fmt.Println(result)
}
