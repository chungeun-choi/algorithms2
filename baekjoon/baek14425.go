package baekjoon

import (
	"bufio"
	"fmt"
	"os"
)

func Baek14425() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	sSet := make(map[string]struct{})
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < n; i++ {
		scanner.Scan()
		sSet[scanner.Text()] = struct{}{}
	}

	count := 0
	for i := 0; i < m; i++ {
		scanner.Scan()
		if _, exists := sSet[scanner.Text()]; exists {
			count++
		}
	}

	fmt.Println(count)
}
