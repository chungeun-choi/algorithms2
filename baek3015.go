package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func baek3015() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력 받기
	nStr, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(nStr[:len(nStr)-1])

	oasis := make([]int, n)
	for i := 0; i < n; i++ {
		heightStr, _ := reader.ReadString('\n')
		height, _ := strconv.Atoi(heightStr[:len(heightStr)-1])
		oasis[i] = height
	}

	// 스택 및 결과값 초기화
	var stack []pair
	var result int64 = 0

	// 오아시스 재결합 알고리즘
	for _, o := range oasis {
		// 스택 끝 값보다 키가 크면 pop
		for len(stack) > 0 && stack[len(stack)-1].height < o {
			result += int64(stack[len(stack)-1].count)
			stack = stack[:len(stack)-1]
		}

		// 스택이 비어있으면 해당 키 추가하고 continue
		if len(stack) == 0 {
			stack = append(stack, pair{o, 1})
			continue
		}

		// 스택 끝 값의 키와 같을 때
		if stack[len(stack)-1].height == o {
			cnt := stack[len(stack)-1].count
			stack = stack[:len(stack)-1]
			result += int64(cnt)

			if len(stack) > 0 {
				result++
			}
			stack = append(stack, pair{o, cnt + 1})
		} else {
			// 스택 끝 값의 키보다 작을 때
			stack = append(stack, pair{o, 1})
			result++
		}
	}

	// 결과값 출력
	fmt.Fprintln(writer, result)
}

// 키와 같은 키의 사람 수를 저장하기 위한 구조체
type pair struct {
	height int
	count  int
}
