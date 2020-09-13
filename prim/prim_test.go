package prim

import (
	"algorithm/graph"
	"testing"
)

func TestMHeap_Prim(t *testing.T) {
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
		{10, 5, 7, 10, 10, 10, 2},
		{5, 10, 10, 9, 10, 10, 3},
		{7, 10, 10, 10, 8, 10, 10},
		{10, 9, 10, 10, 10, 4, 10},
		{10, 10, 8, 10, 10, 5, 4},
		{10, 10, 10, 4, 5, 10, 6},
		{2, 3, 10, 10, 4, 6, 10},
	}

	const N = 7
	m := graph.New(N)
	m.Add(vertices...)
	m.SetEdges(MTX)

	for _, n := range Prim(m) {
		t.Logf("%2d(%s) -> %2d(%s): %.4f\r\n", n.GetS(), m.GetVertices()[n.GetS()], n.GetE(), m.GetVertices()[n.GetE()], n.GetD())
	}
}
