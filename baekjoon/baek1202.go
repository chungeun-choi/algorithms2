package baekjoon

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

// Jewel represents a jewel with weight and value
type Jewel struct {
	weight int
	value  int
}

// MaxHeap is a max-heap of int
type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Baek1202() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	fmt.Fscanf(reader, "%d %d\n", &n, &k)

	jewels := make([]Jewel, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d %d\n", &jewels[i].weight, &jewels[i].value)
	}

	bags := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscanf(reader, "%d\n", &bags[i])
	}

	// Sort jewels by weight
	sort.Slice(jewels, func(i, j int) bool {
		return jewels[i].weight < jewels[j].weight
	})

	// Sort bags by their weight capacity
	sort.Ints(bags)

	totalValue := 0
	jewelIndex := 0
	maxHeapObj := &maxHeap{}
	heap.Init(maxHeapObj)

	for _, bag := range bags {
		for jewelIndex < n && jewels[jewelIndex].weight <= bag {
			heap.Push(maxHeapObj, jewels[jewelIndex].value)
			jewelIndex++
		}
		if maxHeapObj.Len() > 0 {
			totalValue += heap.Pop(maxHeapObj).(int)
		}
	}

	fmt.Fprintln(writer, totalValue)
}
