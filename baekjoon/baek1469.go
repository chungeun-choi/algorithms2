package baekjoon

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const BLANK = -1

// DFS function to find the valid sequence
func dfs(count int, N int, S []int, X []int) {
	if count == N {
		// Print the sequence if valid
		for _, num := range S {
			fmt.Print(num, " ")
		}
		fmt.Println()
		os.Exit(0)
	}

	for _, i := range X {
		// Check if the number is already used in the sequence
		if contains(S, i) {
			continue
		}
		// Find the next blank index
		nextIdx := indexOf(S, BLANK)
		if nextIdx+i+1 >= N*2 {
			break
		}
		// Check if the position to place the second i is available
		if S[nextIdx+i+1] != BLANK {
			continue
		}

		// Place the number i at the calculated positions
		S[nextIdx] = i
		S[nextIdx+i+1] = i
		// Recursively call DFS with updated count
		dfs(count+1, N, S, X)
		// Backtrack and reset the positions to BLANK
		S[nextIdx] = BLANK
		S[nextIdx+i+1] = BLANK
	}
}

// Check if a slice contains a value
func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// Find the index of the first occurrence of a value in a slice
func indexOf(slice []int, value int) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

// Solve function to initiate DFS
func solve(N int, X []int) {
	S := make([]int, N*2)
	for i := range S {
		S[i] = BLANK
	}
	dfs(0, N, S, X)
	fmt.Println(-1)
}

// Main function to read input and call solve
func baek1469() {
	reader := bufio.NewReader(os.Stdin)

	// Read the integer N
	nInput, _ := reader.ReadString('\n')
	nInput = strings.TrimSpace(nInput)
	N, _ := strconv.Atoi(nInput)

	// Read the integers X
	xInput, _ := reader.ReadString('\n')
	xInput = strings.TrimSpace(xInput)
	xStrings := strings.Split(xInput, " ")
	X := make([]int, N)
	for i, s := range xStrings {
		X[i], _ = strconv.Atoi(s)
	}

	// Sort X to process in ascending order
	sort.Ints(X)
	solve(N, X)
}
