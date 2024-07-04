package baekjoon

import (
	"bufio"
	"fmt"
	"os"
)

func Baek13144() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	lastOccurrence := make(map[int]int)

	left := 0
	answer := int64(0)

	for right := 0; right < n; right++ {
		if last, found := lastOccurrence[arr[right]]; found && last >= left {
			left = last + 1
		}

		lastOccurrence[arr[right]] = right
		answer += int64(right - left + 1)
	}

	fmt.Fprintln(writer, answer)
}
