package baekjoon

import (
	"bufio"
	"fmt"
	"os"
)

// 방향 이동을 위한 배열
var directions = [4][2]int{
	{-1, 0}, // 상
	{1, 0},  // 하
	{0, -1}, // 좌
	{0, 1},  // 우
}

// isConnected 함수: BFS를 통해 선택된 7개의 자리가 인접해 있는지 확인
func isConnected(grid [5][5]rune, positions []int) bool {
	visited := make([]bool, 7)
	queue := []int{positions[0]}
	visited[0] = true
	connectedCount := 1

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		x, y := current/5, current%5

		for _, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]
			if nx < 0 || ny < 0 || nx >= 5 || ny >= 5 {
				continue
			}
			next := nx*5 + ny
			for j := 0; j < 7; j++ {
				if !visited[j] && positions[j] == next {
					visited[j] = true
					queue = append(queue, next)
					connectedCount++
				}
			}
		}
	}

	return connectedCount == 7
}

// backtrack 함수: 백트래킹을 통해 25개 자리 중 7개를 선택
func backtrack3(grid [5][5]rune, start, depth, sCount int, selected []int, result *int) {
	if depth == 7 {
		if sCount >= 4 && isConnected(grid, selected) {
			*result++
		}
		return
	}

	for i := start; i < 25; i++ {
		x, y := i/5, i%5
		newSCount := sCount
		if grid[x][y] == 'S' {
			newSCount++
		}
		selected = append(selected, i)
		backtrack3(grid, i+1, depth+1, newSCount, selected, result)
		selected = selected[:len(selected)-1]
	}
}

func baek1941() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var grid [5][5]rune

	// 5x5 격자를 입력받음
	for i := 0; i < 5; i++ {
		scanner.Scan()
		line := scanner.Text()
		for j, ch := range line {
			grid[i][j] = ch
		}
	}

	result := 0
	selected := make([]int, 0, 7)
	backtrack3(grid, 0, 0, 0, selected, &result)

	fmt.Fprintln(writer, result)
}
