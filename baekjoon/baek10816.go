package baekjoon

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Baek10816() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscanln(reader, &n)
	cards := make(map[int]int)

	// Read and process the card numbers
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	strs := strings.Split(str, " ")

	for _, value := range strs {
		number, _ := strconv.Atoi(value)
		cards[number]++
	}

	fmt.Fscanln(reader, &m)

	// Read and process the numbers to count
	str, _ = reader.ReadString('\n')
	str = strings.TrimSpace(str)
	strs = strings.Split(str, " ")

	// Output the count of each number
	for _, value := range strs {
		number, _ := strconv.Atoi(value)
		fmt.Fprintf(writer, "%d ", cards[number])
	}
	fmt.Fprintln(writer) // Print a new line at the end
}
