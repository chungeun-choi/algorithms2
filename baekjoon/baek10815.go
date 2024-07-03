package baekjoon

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Baek10815() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var (
		n, m     int
		cardsMap map[int]int
	)

	fmt.Fscanln(reader, &n)

	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	words := strings.Fields(line)

	cardsMap = make(map[int]int)

	for _, value := range words {
		number, _ := strconv.Atoi(value)
		cardsMap[number]++
	}

	fmt.Fscanln(reader, &m)

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	words = strings.Fields(line)

	for _, value := range words {
		number, _ := strconv.Atoi(value)
		if _, ok := cardsMap[number]; ok {
			fmt.Fprintf(writer, "%d ", cardsMap[number])
		} else {
			fmt.Fprintf(writer, "%d ", 0)
		}
	}
}
