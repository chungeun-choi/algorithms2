package baekjoon

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func recursive(n, row, col int) int {
	if n == 0 {
		return 0
	}
	curCount := 2*(row%2) + (col % 2)
	return 4*recursive(n-1, row/2, col/2) + curCount
}

func Baek1074() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)
	N, _ := strconv.Atoi(parts[0])
	r, _ := strconv.Atoi(parts[1])
	c, _ := strconv.Atoi(parts[2])

	result := recursive(N, r, c)
	fmt.Fprintln(writer, result)
}
