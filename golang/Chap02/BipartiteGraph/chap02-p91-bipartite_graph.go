package main

import (
	"fmt"
)

type Vertex struct {
	edge []int
}

var G []Vertex
var color []int

func dfs(v, c int) bool {
	color[v] = c
	for i := 0; i < len(G[v].edge); i++ {
		if color[G[v].edge[i]] == c {
			return false
		}
		if color[G[v].edge[i]] == 0 && !dfs(G[v].edge[i], -c) {
			return false
		}
	}
	return true
}

func solve(n, m int) {
	color = make([]int, n)
	G = make([]Vertex, n)
	for i := 0; i < n; i++ {
		color[i] = 0
	}
	var v1, v2 int
	for i := 0; i < m; i++ {
		fmt.Scan(&v1, &v2)
		G[v1].edge = append(G[v1].edge, v2)
		G[v2].edge = append(G[v2].edge, v1)
	}
	for i := 0; i < n; i++ {
		if color[i] == 0 {
			if !dfs(i, 1) {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
}

func main() {
	var t, n, m int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&n, &m)
		solve(n, m)
	}
}
