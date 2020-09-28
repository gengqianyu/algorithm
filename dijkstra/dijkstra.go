package dijkstra

import (
	"algorithm/graph"
)

const MaxWeigh int = 10000

//迪杰斯特拉求图中一个顶点到其他顶点的最短路径问题
func Dijkstra(start int, m *graph.GoMap) []int {
	l := len(m.GetVertices()) //顶点个数

	shortPath := make([]int, l) //defined shortPath record  起始顶点到各顶点的最小距离
	flag := make([]bool, l)     //记录顶点是否已经找到v0到vx的最小路径
	prePath := make([]int, l)   //记录当前顶点的前驱顶点，以便重新计算最小边距

	for i := range shortPath {
		//如果是出发顶点，就马上处理
		if i == start {
			shortPath[i] = 0
			flag[i] = true
			continue
		}
		shortPath[i] = MaxWeigh
	}

	for {
		//遍历邻接矩阵的起始顶点，和其他顶点的关系，更新起始start索引顶点，到周围顶点的距离和周围顶点的前驱顶点
		for j := 0; j < len(m.Edges()[start]); j++ {
			//如果j顶点没有被访问，并且出发start顶点到当前顶点的距离+当前顶点到j顶点距离之和 ，比出发顶点直接到j顶点的距离还要小，就更新
			if !flag[j] && (shortPath[start]+int(m.Edges()[start][j])) < shortPath[j] {
				//更新当前顶点的前驱节点
				prePath[j] = start
				//更新出发顶点到j顶点的距离
				shortPath[j] = shortPath[start] + int(m.Edges()[start][j])
			}
		}

		//	获取新顶点作为初始顶点，重复以上操作
		min := MaxWeigh
		end := start
		for i := 0; i < len(flag); i++ {
			if !flag[i] && shortPath[i] < min {
				min = shortPath[i]
				start = i
			}
		}
		//未获得新顶点就结束
		if start == end {
			break
		}
		//更新新顶点为已访问
		flag[start] = true
	}

	return shortPath
}
