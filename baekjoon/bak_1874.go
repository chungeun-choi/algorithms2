package baekjoon

import (
	"bufio"
	"fmt"
	"os"
)

func Bak1874() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// Read the number of elements
	var n int
	fmt.Fscanln(reader, &n)

	// Read the sequence of numbers
	sequence := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &sequence[i])
	}

	stack := []int{}
	operations := []string{}
	current := 1

	for _, num := range sequence {
		for current <= num {
			stack = append(stack, current)
			operations = append(operations, "+")
			current++
		}
		if stack[len(stack)-1] == num {
			stack = stack[:len(stack)-1]
			operations = append(operations, "-")
		} else {
			fmt.Fprintln(writer, "NO")
			return
		}
	}

	for _, op := range operations {
		fmt.Fprintln(writer, op)
	}
}
