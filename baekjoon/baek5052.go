package baekjoon

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func Baek5052() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var (
		t, n int
	)

	fmt.Fscanln(reader, &t)
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &n)
		phoneNumbers := make([]string, n)

		for j := 0; j < n; j++ {
			var phoneNumber string
			fmt.Fscanln(reader, &phoneNumber)
			phoneNumbers[j] = phoneNumber
		}

		sort.Strings(phoneNumbers)
		consistent := true

		for k := 0; k < n-1; k++ {
			if strings.HasPrefix(phoneNumbers[k+1], phoneNumbers[k]) {
				consistent = false
				break
			}
		}

		if consistent {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}
