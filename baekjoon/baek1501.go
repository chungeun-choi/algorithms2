package baekjoon

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Function to generate the key for each word
func generateKey(word string) string {
	length := len(word)
	if length <= 2 {
		return fmt.Sprintf("%d%s", length, word)
	}
	middle := []rune(word[1 : length-1])
	sort.Slice(middle, func(i, j int) bool {
		return middle[i] < middle[j]
	})
	return fmt.Sprintf("%d%c%s%c", length, word[0], string(middle), word[length-1])
}

func baek1501() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// Read the number of words
	nInput, _ := reader.ReadString('\n')
	nInput = strings.TrimSpace(nInput)
	n, _ := strconv.Atoi(nInput)

	// Create a dictionary to count occurrences of each key
	dic := make(map[string]int)

	// Read each word and store its key in the dictionary
	for i := 0; i < n; i++ {
		wordInput, _ := reader.ReadString('\n')
		word := strings.TrimSpace(wordInput)
		key := generateKey(word)
		dic[key]++
	}

	// Read the number of queries
	mInput, _ := reader.ReadString('\n')
	mInput = strings.TrimSpace(mInput)
	m, _ := strconv.Atoi(mInput)

	// Process each query
	for i := 0; i < m; i++ {
		queryInput, _ := reader.ReadString('\n')
		query := strings.TrimSpace(queryInput)
		words := strings.Split(query, " ")

		result := 1
		check := 0
		for _, word := range words {
			key := generateKey(word)
			if count, exists := dic[key]; exists {
				result *= count
				check++
			}
		}

		if check == 0 {
			fmt.Fprintln(writer, 0)
		} else {
			fmt.Fprintln(writer, result)
		}
	}
}
