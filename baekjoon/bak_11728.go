package baekjoon

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func bak11728() {
	var a, b int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanf(reader, "%d %d\n", &a, &b)

	if a == 0 && b == 0 {
		writer.WriteString("0")
	}

	var arr1 []int = make([]int, a)
	var arr2 []int = make([]int, b)

	line1, _ := reader.ReadString('\n')
	line1 = strings.TrimSpace(line1)
	strArr1 := strings.Split(line1, " ")
	for i := range strArr1 {
		arr1[i], _ = strconv.Atoi(strArr1[i])
	}

	line2, _ := reader.ReadString('\n')
	line2 = strings.TrimSpace(line2)
	strArr2 := strings.Split(line2, " ")
	for i := range strArr2 {
		arr2[i], _ = strconv.Atoi(strArr2[i])
	}

	sort.Ints(arr1)
	sort.Ints(arr2)

	n, m := len(arr1), len(arr2)
	var result []int = make([]int, 0)
	var i, j int = 0, 0

	for i < n && j < m {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}

	for i < n {
		result = append(result, arr1[i])
		i++
	}

	for j < m {
		result = append(result, arr2[j])
		j++
	}

	// Output the result using writer
	for i, res := range result {
		if i > 0 {
			writer.WriteString(" ")
		}
		writer.WriteString(strconv.Itoa(res))
	}
}
