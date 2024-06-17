package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func baek1245() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var (
		n, m int
		dirs [][2]int = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	)

	fmt.Fscanf(reader, "%d %d\n", &n, &m)

	// mountain 초기화
	mountain := make([][]int, n)
	visited := make([][]bool, n)

	for i := range mountain {
		mountain[i] = make([]int, m)
		visited[i] = make([]bool, m)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		texts := strings.Split(text, " ")

		for index, value := range texts {
			number, _ := strconv.Atoi(value)
			mountain[i][index] = number
		}
	}

	isPeak := func(x, y int) bool {
		queue := [][2]int{{x, y}}
		visited[x][y] = true
		isPeak := true

		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			currX, currY := current[0], current[1]

			for _, dir := range dirs {
				dx, dy := currX+dir[0], currY+dir[1]
				if 0 <= dx && dx < n && 0 <= dy && dy < m {
					if mountain[dx][dy] > mountain[currX][currY] {
						isPeak = false
					}
					if !visited[dx][dy] && mountain[dx][dy] == mountain[currX][currY] {
						visited[dx][dy] = true
						queue = append(queue, [2]int{dx, dy})
					}
				}
			}
		}
		return isPeak
	}

	cnt := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !visited[i][j] && isPeak(i, j) {
				cnt++
			}
		}
	}

	fmt.Fprintln(writer, cnt)
}
