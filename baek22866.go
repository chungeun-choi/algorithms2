package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func baek22866() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	// 입력 받기
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		arr[i] = num
	}

	// 보이는 것만 담을 스택
	answer := make([][]int, n+1)
	for i := range answer {
		answer[i] = []int{0, 0}
	}

	leftStack := [][]int{}
	rightStack := [][]int{}

	// 왼쪽에서 오른쪽으로 순회
	for i := 1; i <= n; i++ {
		for len(leftStack) > 0 && leftStack[len(leftStack)-1][1] <= arr[i] {
			leftStack = leftStack[:len(leftStack)-1]
		}

		if len(leftStack) > 0 {
			answer[i][1] = leftStack[len(leftStack)-1][0]
		}
		answer[i][0] += len(leftStack)
		leftStack = append(leftStack, []int{i, arr[i]})
	}

	// 오른쪽에서 왼쪽으로 순회
	for i := n; i >= 1; i-- {
		for len(rightStack) > 0 && rightStack[len(rightStack)-1][1] <= arr[i] {
			rightStack = rightStack[:len(rightStack)-1]
		}

		if len(rightStack) > 0 {
			if answer[i][0] == 0 {
				answer[i][1] = rightStack[len(rightStack)-1][0]
			} else {
				prevDist := abs(i - answer[i][1])
				nowDist := abs(i - rightStack[len(rightStack)-1][0])
				if nowDist < prevDist {
					answer[i][1] = rightStack[len(rightStack)-1][0]
				}
			}
		}
		answer[i][0] += len(rightStack)
		rightStack = append(rightStack, []int{i, arr[i]})
	}

	// 결과 출력 (버퍼를 사용하여 빠르게)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for i := 1; i <= n; i++ {
		if answer[i][0] == 0 {
			writer.WriteString("0\n")
		} else {
			writer.WriteString(fmt.Sprintf("%d %d\n", answer[i][0], answer[i][1]))
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
