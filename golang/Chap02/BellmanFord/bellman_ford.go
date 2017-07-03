package main

import (
	"fmt"
)

type Edge struct {
	from int
	to   int
	cost int
}

var G []Edge

var d []int
var V, E int

const INF = 1000000

var Letters = map[string]int{
	"A": 0,
	"B": 1,
	"C": 2,
	"D": 3,
	"E": 4,
	"F": 5,
	"G": 6,
}

func shortestPath(s int) {
	for i := 0; i < V; i++ {
		d[i] = INF
	}
	d[s] = 0
	for {
		update := false
		for i := 0; i < E; i++ {
			e := G[i]
			if d[e.from] != INF && d[e.to] > d[e.from]+e.cost {
				d[e.to] = d[e.from] + e.cost
				update = true
			}
		}
		if !update {
			break
		}
	}
}

func main() {
	fmt.Scan(&V, &E)
	var from, to string
	var cost int
	G = make([]Edge, E)
	d = make([]int, V)
	for i := 0; i < E; i++ {
		fmt.Scan(&from, &to, &cost)
		G[i] = Edge{Letters[from], Letters[to], cost}
	}
	shortestPath(0)
	fmt.Println(d[Letters["G"]])
}
