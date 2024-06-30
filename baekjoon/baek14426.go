package baekjoon

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func baek14426() {
	reader := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)

	S := make([]string, n)
	for i := 0; i < n; i++ {
		S[i], _ = reader.ReadString('\n')
		S[i] = strings.TrimSpace(S[i])
	}
	sort.Strings(S)

	s := make([]string, m)
	for i := 0; i < m; i++ {
		s[i], _ = reader.ReadString('\n')
		s[i] = strings.TrimSpace(s[i])
	}
	sort.Strings(s)

	answer := 0
	i, j := 0, 0
	for i < n && j < m {
		if strings.HasPrefix(S[i], s[j]) {
			answer++
			j++
		} else if S[i] > s[j] {
			j++
		} else {
			i++
		}
	}

	fmt.Println(answer)
}
