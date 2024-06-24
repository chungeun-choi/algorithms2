package main

import (
	"bufio"
	"fmt"
	"os"
)

func sol3() {
	// s의 길이가 m과 같으면 s를 출력합니다.
	if len(s) == m {
		for _, num := range s {
			writer.WriteString(fmt.Sprintf("%d ", num))
		}
		writer.WriteString("\n")
		return
	}

	for i := 1; i <= n; i++ {
		if len(s) > 0 {
			if s[len(s)-1] > i {
				continue
			}
		}

		s = append(s, i)

		sol3()
		s = s[:len(s)-1]
	}
}

func baek15652() {
	fmt.Scan(&n, &m)

	writer = bufio.NewWriter(os.Stdout)

	sol3()

	writer.Flush()
}
