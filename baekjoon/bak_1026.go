package baekjoon

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func bak1026() {
	reader := bufio.NewReader(os.Stdin)

	var (
		n      int
		listA  []int
		listB  []int
		result int
	)

	fmt.Fscanf(reader, "%d\n", &n)
	listA = make([]int, n)
	listB = make([]int, n)

	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	strs := strings.Split(str, " ")
	for i, value := range strs {
		listA[i], _ = strconv.Atoi(value)
	}

	str2, _ := reader.ReadString('\n')
	str2 = strings.TrimSpace(str2)
	strs2 := strings.Split(str2, " ")
	for i, value := range strs2 {
		listB[i], _ = strconv.Atoi(value)
	}

	sort.Slice(listA, func(i, j int) bool {
		return listA[i] < listA[j]
	})

	sort.Slice(listB, func(i, j int) bool {
		return listB[i] > listB[j]
	})

	for i := 0; i < n; i++ {
		result += listA[i] * listB[i]
	}

	fmt.Println(result)
}
