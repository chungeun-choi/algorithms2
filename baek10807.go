package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func baek10807() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var (
		n         int
		v         int
		numberMap map[int]int = make(map[int]int)
	)

	fmt.Fscanf(reader, "%d\n", &n)

	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	numbers := strings.Split(str, " ")

	for _, value := range numbers {
		number, _ := strconv.Atoi(value)

		if _, ok := numberMap[number]; !ok {
			numberMap[number] = 1
		} else {
			numberMap[number]++
		}
	}

	fmt.Fscanf(reader, "%d\n", &v)

	fmt.Fprintln(writer, numberMap[v])

}
