package baekjoon

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Bak10773() int {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N int
	fmt.Fscanln(reader, &N)

	stack := make([]int, 0, N)
	var result int

	for i := 0; i < N; i++ {
		var inputStr string
		inputStr, _ = reader.ReadString('\n')
		input, _ := strconv.Atoi(inputStr[:len(inputStr)-1])

		if input != 0 {
			stack = append(stack, input)
		} else if len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
	}

	for _, value := range stack {
		result += value
	}

	return result
}
