package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bak1270() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscanf(reader, "%d\n", &t)

	for i := 0; i < t; i++ {
		line, _ := reader.ReadString('\n')
		n, _ := strconv.Atoi(strings.TrimSpace(line))

		candidate := 0
		count := 0

		line, _ = reader.ReadString('\n')
		numbers := strings.Fields(line)

		for j := 0; j < n; j++ {
			num, _ := strconv.Atoi(numbers[j])
			if count == 0 {
				candidate = num
				count = 1
			} else if candidate == num {
				count++
			} else {
				count--
			}
		}

		// Verify if the candidate is indeed the majority element
		finalCount := 0
		for j := 0; j < n; j++ {
			num, _ := strconv.Atoi(numbers[j])
			if num == candidate {
				finalCount++
			}
		}

		if finalCount > n/2 {
			fmt.Fprintln(writer, candidate)
		} else {
			fmt.Fprintln(writer, "SYJKGW")
		}
	}
}
