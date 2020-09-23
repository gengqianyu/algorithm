package graph

import (
	"fmt"
	"testing"
)

func TestGoMap_Add(t *testing.T) {
	//vertices := []string{"A", "B", "C", "D", "E"}
	vertices := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	goMap := New(len(vertices))
	//添加关系
	goMap.Add(vertices...)
	//添加边
	//goMap.SetEdge(0, 1, 1) //A-B
	//goMap.SetEdge(0, 2, 1) //A-C
	//goMap.SetEdge(0, 4, 1) //A-E
	//goMap.SetEdge(1, 2, 1) //B-C
	//goMap.SetEdge(1, 3, 1) //B-D
	//goMap.SetEdge(1, 4, 1) //B-E

	goMap.SetEdge(0, 1, 1) //A-B
	goMap.SetEdge(0, 2, 1) //A-C
	goMap.SetEdge(1, 3, 1) //B-D
	goMap.SetEdge(1, 4, 1) //B-E
	goMap.SetEdge(3, 7, 1)
	goMap.SetEdge(4, 7, 1)
	goMap.SetEdge(2, 5, 1)
	goMap.SetEdge(2, 6, 1)
	goMap.SetEdge(5, 6, 1)

	for _, edge := range goMap.Edges() {
		for _, v := range edge {
			fmt.Print(v)
		}
		fmt.Println()
	}
	//深度
	//for vertex := range goMap.DFS() {
	//	fmt.Println(vertex)
	//}

	//广度优先
	for vertex := range goMap.BFS() {
		fmt.Println(vertex)
	}
}
