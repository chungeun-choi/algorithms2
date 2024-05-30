package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func bak2577() {
	var a, b, c int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &b)
	fmt.Fscanln(reader, &c)

	multi := a * b * c
	stringMulti := strconv.Itoa(multi)

	var countMap []int = make([]int, 10)

	for i := 0; i < len(stringMulti); i++ {
		number, _ := strconv.Atoi(string(stringMulti[i]))
		countMap[number]++
	}

	for i := 0; i < len(countMap); i++ {
		fmt.Fprintln(writer, countMap[i])
	}

}
