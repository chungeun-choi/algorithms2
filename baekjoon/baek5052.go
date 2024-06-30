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
		var (
			pi, pj       int = 0, 1
			phoneNumbers []string
		)

		phoneNumbers = make([]string, n)
		for index, _ := range phoneNumbers {
			var phoneNumber string
			fmt.Fscanln(reader, &phoneNumber)
			phoneNumbers[index] = phoneNumber
		}

		sort.Strings(phoneNumbers)

		for pi < n && pj < n {
			if strings.HasPrefix(phoneNumbers[pj], phoneNumbers[pi]) {
				fmt.Println("NO")
				break
			} else if phoneNumbers[pi] < phoneNumbers[pj] {
				pj++
			} else {
				pi++
			}
		}
		fmt.Println("YES")
	}

}
