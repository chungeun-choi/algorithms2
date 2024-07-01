package baekjoon

import (
	"bufio"
	"fmt"
	"os"
)

func Baek2164() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var (
		n     int
		queue []int
	)
	fmt.Fscanln(reader, &n)

	queue = make([]int, n)

	for i := range queue {
		queue[i] = i + 1
	}

	for len(queue) > 1 {
		// 1. 가장 위의 숫자를 버림
		queue = queue[1:]
		// 2. 가장 위의 숫자를 맨 뒨로 변경
		queue = append(queue[1:], queue[0])
	}

	fmt.Fprintln(writer, queue[0])

}
