package kruskal

import (
	"algorithm/graph"
	"container/heap"
)

//定义图的边
type edge struct {
	s, e int     //起点终点
	d    float64 //距离
}

func (n *Node) SetS(v int) {
	n.s = v
}

func (n *Node) GetS() int {
	return n.s
}

func (n *Node) SetE(v int) {
	n.e = v
}
func (n *Node) GetE() int {
	return n.e
}

//堆节点
type Node struct {
	edge
	//p, l, r *Node //父节点，左子节点，右子节点
}

func (n *Node) SetD(v interface{}) {
	n.d = v.(float64)
}
func (n *Node) GetD() float64 {
	return n.d
}

//定义一个小堆
type MHeap []*Node

//实现 heap.Interface 的Len方法
func (h *MHeap) Len() int {
	return len(*h)
}

//实现 heap.Interface 的Less方法
func (h *MHeap) Less(i, j int) bool {
	return (*h)[i].GetD() < (*h)[j].GetD()
}

//实现 heap.Interface 的Swap方法
func (h *MHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

//实现 heap.Interface 的Push方法
func (h *MHeap) Push(v interface{}) {
	*h = append(*h, v.(*Node))
}

//实现 heap.Interface 的Pop方法
func (h *MHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}

//利用克鲁斯卡尔算法解决最短路径问题
//1首先图中所有图节点都创建为单元素的集合；然后将图的所有边,按长短排序。
//2依次从小到大取出每一条边。如果某条边它的两个顶点点分属不同集合，则记录该边，并合并两个集合；否则丢弃该边。
//3重复1，2步骤，直到所有边均被处理完毕。
//用固定长度数组和伪链表、查询更新来实现。
func Kruskal(m *graph.GoMap) (s []*Node) {
	//初始化最小堆
	h := new(MHeap)
	heap.Init(h)

	//遍历矩阵, 将所有边加入最小堆中,小堆自然按边长排序
	for i, row := range m.Edges() {
		for j, col := range row {
			//跳过当前顶点到自己的边关系
			if i == j {
				continue
			}
			n := new(Node)
			n.SetS(i)
			n.SetE(j)
			n.SetD(col)
			heap.Push(h, n)
		}
	}

	//2依次从小到大取出每一条边。如果某条边它的两个顶点点分属不同集合，则记录该边，并合并两个集合；否则丢弃该边。

	return s
}
