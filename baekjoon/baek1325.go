package baekjoon

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bfs1325(start int, N int, computer [][]int) int {
	visited := make([]bool, N+1)
	queue := []int{start}
	visited[start] = true
	cnt := 1

	for len(queue) > 0 {
		x := queue[0]
		queue = queue[1:]

		for _, nx := range computer[x] {
			if !visited[nx] {
				queue = append(queue, nx)
				visited[nx] = true
				cnt++
			}
		}
	}

	return cnt
}

func Baek1325() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력 받기
	input, _ := reader.ReadString('\n')
	parts := strings.Fields(input)
	N, _ := strconv.Atoi(parts[0])
	M, _ := strconv.Atoi(parts[1])

	computer := make([][]int, N+1)
	for i := 0; i < M; i++ {
		input, _ := reader.ReadString('\n')
		parts := strings.Fields(input)
		A, _ := strconv.Atoi(parts[0])
		B, _ := strconv.Atoi(parts[1])
		computer[B] = append(computer[B], A)
	}

	answer := []int{}
	maxHack := 0
	for i := 1; i <= N; i++ {
		result := bfs1325(i, N, computer)
		if maxHack < result {
			maxHack = result
			answer = []int{i}
		} else if maxHack == result {
			answer = append(answer, i)
		}
	}

	// 결과 출력
	for _, v := range answer {
		fmt.Fprintf(writer, "%d ", v)
	}
	fmt.Fprintln(writer)
}
