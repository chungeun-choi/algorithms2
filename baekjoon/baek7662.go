package baekjoon

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MinHeapItem struct {
	value int
	index int
}

type MaxHeapItem struct {
	value int
	index int
}

type MinHeap []MinHeapItem
type MaxHeap []MaxHeapItem

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].value < h[j].value }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(MinHeapItem))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].value > h[j].value }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(MaxHeapItem))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Baek7662() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscanf(reader, "%d\n", &t)

	for i := 0; i < t; i++ {
		var k int
		fmt.Fscanf(reader, "%d\n", &k)

		minHeap := &MinHeap{}
		maxHeap := &MaxHeap{}
		heap.Init(minHeap)
		heap.Init(maxHeap)

		check := make([]int, k)

		for j := 0; j < k; j++ {
			line, _ := reader.ReadString('\n')
			parts := strings.Fields(line)
			cal := parts[0]
			num, _ := strconv.Atoi(parts[1])

			if cal == "I" {
				heap.Push(minHeap, MinHeapItem{value: num, index: j})
				heap.Push(maxHeap, MaxHeapItem{value: num, index: j})
				check[j] = 1
			} else {
				if num == -1 {
					for minHeap.Len() > 0 && check[(*minHeap)[0].index] == 0 {
						heap.Pop(minHeap)
					}
					if minHeap.Len() > 0 {
						check[(*minHeap)[0].index] = 0
						heap.Pop(minHeap)
					}
				} else if num == 1 {
					for maxHeap.Len() > 0 && check[(*maxHeap)[0].index] == 0 {
						heap.Pop(maxHeap)
					}
					if maxHeap.Len() > 0 {
						check[(*maxHeap)[0].index] = 0
						heap.Pop(maxHeap)
					}
				}
			}

			for minHeap.Len() > 0 && check[(*minHeap)[0].index] == 0 {
				heap.Pop(minHeap)
			}
			for maxHeap.Len() > 0 && check[(*maxHeap)[0].index] == 0 {
				heap.Pop(maxHeap)
			}
		}

		if minHeap.Len() == 0 {
			fmt.Fprintln(writer, "EMPTY")
		} else {
			fmt.Fprintf(writer, "%d %d\n", (*maxHeap)[0].value, (*minHeap)[0].value)
		}
	}
}
