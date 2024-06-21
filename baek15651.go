package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n, m   int
	s      []int
	writer *bufio.Writer
)

func sol2() {
	// s의 길이가 m과 같으면 s를 출력합니다.
	if len(s) == m {
		for _, num := range s {
			writer.WriteString(fmt.Sprintf("%d ", num))
		}
		writer.WriteString("\n")
		return
	}

	for i := 1; i <= n; i++ {
		s = append(s, i)
		sol2()
		s = s[:len(s)-1]
	}
}

func baek15651() {
	fmt.Scan(&n, &m)

	writer = bufio.NewWriter(os.Stdout)

	sol2()

	writer.Flush()
}
