package dijkstra

import (
	"algorithm/graph"
	"fmt"
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
	start := 2
	shortPath, prePath := Dijkstra(start, m)
	t.Log(prePath)
	for i, n := range shortPath {
		t.Logf("%s到%s的最短路径为：%d\r\n", m.GetVertices()[start], m.GetVertices()[i], n)
		for _, index := range prePath[i] {
			fmt.Print(m.GetVertices()[index], "->")
		}
		fmt.Println()
	}
}
