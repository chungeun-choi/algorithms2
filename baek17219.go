package main

import (
	"bufio"
	"fmt"
	"os"
)

func baek17219() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var (
		n, m    int
		keyList map[string]string = make(map[string]string)
	)

	fmt.Fscanf(reader, "%d %d\n", &n, &m)

	for i := 0; i < n; i++ {
		var key, value string
		fmt.Fscanf(reader, "%s %s\n", &key, &value)

		keyList[key] = value
	}

	for i := 0; i < m; i++ {
		var key string
		fmt.Fscanln(reader, &key)
		fmt.Fprintln(writer, keyList[key])
	}

}
