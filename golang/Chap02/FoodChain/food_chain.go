package main

import (
	"fmt"
)

type UnionFind struct {
	Par  []int
	Rank []int
}

func Init(n int) UnionFind {
	par := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		par[i] = i
		rank[i] = 0
	}
	return UnionFind{par, rank}
}

func (uf UnionFind) Find(x int) int {
	if uf.Par[x] == x {
		return x
	}
	uf.Par[x] = uf.Find(uf.Par[x])
	return uf.Par[x]
}

func (uf UnionFind) Unite(x, y int) {
	x = uf.Find(x)
	y = uf.Find(y)
	if x == y {
		return
	}
	if uf.Rank[x] < uf.Rank[y] {
		uf.Par[x] = y
	} else {
		uf.Par[y] = x
		if uf.Rank[x] == uf.Rank[y] {
			uf.Rank[x]++
		}
	}
}

func (uf UnionFind) Same(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	t := make([]int, k)
	x := make([]int, k)
	y := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Scan(&t[i], &x[i], &y[i])
	}
	uf := Init(n * 3)
	ans := 0
	for i := 0; i < k; i++ {
		tt := t[i]
		xx := x[i] - 1
		yy := y[i] - 1

		if xx < 0 || n <= xx || yy < 0 || n <= yy {
			ans++
			continue
		}
		if tt == 1 {
			if uf.Same(xx, yy+n) || uf.Same(xx, yy+2*n) {
				ans++
			} else {
				uf.Unite(xx, yy)
				uf.Unite(xx+n, yy+n)
				uf.Unite(xx+n*2, yy+n*2)
			}
		} else {
			if uf.Same(xx, yy) || uf.Same(xx, yy+2*n) {
				ans++
			} else {
				uf.Unite(xx, yy+n)
				uf.Unite(xx+n, yy+2*n)
				uf.Unite(xx+2*n, yy)
			}
		}
	}
	fmt.Println(ans)
}
