package main

import "fmt"
import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	var n int
	fmt.Scan(&n)
	l := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&l[i])
	}

	que := &IntHeap{}
	heap.Init(que)

	ans := 0

	for i := 0; i < n; i++ {
		heap.Push(que, l[i])
	}
	for que.Len() > 1 {
		l1 := (*que)[0]
		heap.Pop(que)
		l2 := (*que)[0]
		heap.Pop(que)
		ans += (l1 + l2)
		heap.Push(que, l1+l2)
	}
	fmt.Println(ans)
}
