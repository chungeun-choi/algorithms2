package baekjoon

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func backtrack2(numbers []int, selected []int, start, depth int, writer *bufio.Writer) {
	if depth == 6 {
		for i, num := range selected {
			if i > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, num)
		}
		fmt.Fprintln(writer)
		return
	}

	for i := start; i < len(numbers); i++ {
		selected = append(selected, numbers[i])
		backtrack2(numbers, selected, i+1, depth+1, writer)
		selected = selected[:len(selected)-1]
	}
}

func baek6603() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for {
		scanner.Scan()
		line := scanner.Text()

		if line == "0" {
			break
		}

		parts := strings.Fields(line)
		k, _ := strconv.Atoi(parts[0])
		if k == 0 {
			break
		}

		numbers := make([]int, k)
		for i := 0; i < k; i++ {
			numbers[i], _ = strconv.Atoi(parts[i+1])
		}

		selected := make([]int, 0, 6)
		backtrack2(numbers, selected, 0, 0, writer)
		fmt.Fprintln(writer)
	}
}
