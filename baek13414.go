package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func baek13414() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var K, L int
	fmt.Fscanf(reader, "%d %d\n", &K, &L)

	students := make(map[string]int)
	for i := 0; i < L; i++ {
		var id string
		fmt.Fscanf(reader, "%s\n", &id)
		students[id] = i
	}

	type student struct {
		id    string
		order int
	}

	orderList := make([]student, 0, len(students))
	for id, order := range students {
		orderList = append(orderList, student{id, order})
	}

	sort.Slice(orderList, func(i, j int) bool {
		return orderList[i].order < orderList[j].order
	})

	if K > len(orderList) {
		K = len(orderList)
	}

	for i := 0; i < K; i++ {
		fmt.Fprintln(writer, orderList[i].id)
	}

	// Flush the writer buffer to ensure all output is written
	writer.Flush()
}
