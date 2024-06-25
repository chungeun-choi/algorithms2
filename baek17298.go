package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func baek17298() {
	// Initialize reader and writer for fast input and output
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// Read the number of elements
	nStr, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nStr))

	// Read the elements into an array
	arr := make([]int, n)
	line, _ := reader.ReadString('\n')
	numStrs := strings.Fields(line)
	for i := 0; i < n; i++ {
		arr[i], _ = strconv.Atoi(numStrs[i])
	}

	// Result array to store the answers
	result := make([]int, n)
	for i := range result {
		result[i] = -1
	}

	// Stack to store indices of the elements
	stack := make([]int, 0, n)

	for i := 0; i < n; i++ {
		// Pop elements from the stack until we find a greater element
		for len(stack) > 0 && arr[stack[len(stack)-1]] < arr[i] {
			result[stack[len(stack)-1]] = arr[i]
			stack = stack[:len(stack)-1] // Pop the element
		}
		// Push current index to the stack
		stack = append(stack, i)
	}

	// Print the results
	for i := 0; i < n; i++ {
		fmt.Fprintf(writer, "%d ", result[i])
	}
}
