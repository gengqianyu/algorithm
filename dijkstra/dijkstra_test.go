package dijkstra

import (
	"algorithm/graph"
	"testing"
)

func TestDijkstra(t *testing.T) {
	vertices := []string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
	}
	//定义无向图矩阵
	MTX := [][]float64{
		{10000, 5, 7, 10000, 10000, 10000, 2},
		{5, 10000, 10000, 9, 10000, 10000, 3},
		{7, 10000, 10000, 10000, 8, 10000, 10000},
		{10000, 9, 10000, 10000, 10000, 4, 10000},
		{10000, 10000, 8, 10000, 10000, 5, 4},
		{10000, 10000, 10000, 4, 5, 10000, 6},
		{2, 3, 10000, 10000, 4, 6, 10000},
	}

	const N = 7
	//初始化图
	m := graph.New(N)
	m.Add(vertices...)
	m.SetEdges(MTX)

	for _, n := range Dijkstra(6, m) {
		t.Log(n)
	}
}
