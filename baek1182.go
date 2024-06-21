package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func baek1182() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력 읽기
	var n, s int
	fmt.Fscanf(reader, "%d %d\n", &n, &s)

	// 수열 입력 받기
	numberList := make([]int, n)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	values := strings.Split(line, " ")
	for i, value := range values {
		numberList[i], _ = strconv.Atoi(value)
	}

	// 부분 수열의 합 계산 및 결과 저장
	var cnt int

	// 백트래킹 함수 정의
	var findSubsequences func(idx, currentSum int)
	findSubsequences = func(idx, currentSum int) {
		if idx == n {
			if currentSum == s {
				cnt++
			}
			return
		}

		// 현재 인덱스의 값을 포함하지 않음
		findSubsequences(idx+1, currentSum)

		// 현재 인덱스의 값을 포함함
		findSubsequences(idx+1, currentSum+numberList[idx])
	}

	// 부분 수열의 합 계산 시작
	findSubsequences(0, 0)

	// S가 0인 경우, 공집합이 포함되므로 하나를 빼줍니다.
	if s == 0 {
		cnt--
	}

	fmt.Fprintln(writer, cnt)
}
