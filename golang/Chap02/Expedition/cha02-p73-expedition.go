package main

import "fmt"
import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
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
	var n, l, p int
	fmt.Scan(&n, &l, &p)
	a := make([]int, n+1)
	b := make([]int, n+1)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	a = append(a, 1)
	for i := 0; i < n; i++ {
		fmt.Scan(&b[i])
	}
	b = append(b, 0)
	n++

	que := &IntHeap{}
	heap.Init(que)

	ans := 0
	pos := 0
	tank := p

	for i := 0; i < n; i++ {
		d := a[i] - pos
		for tank-d < 0 {
			if que.Len() == 0 {
				fmt.Println(-1)
				return
			}
			tank += (*que)[que.Len()-1]
			heap.Pop(que)
			ans++
		}
		tank -= d
		pos = a[i]
		heap.Push(que, b[i])
	}
	fmt.Println(ans)
}
