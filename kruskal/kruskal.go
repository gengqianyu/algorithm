package kruskal

import (
	"algorithm/graph"
	"container/heap"
	"fmt"
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

type Set []int

//创建一个所有顶点集的总集，因为每个元素都等于下标，即都是单元素集合
func (s *Set) Init(n int) *Set {
	for n--; n >= 0; n-- {
		(*s)[n] = n
	}
	return s
}

func NewSet(n int) *Set {
	set := Set(make([]int, n))
	return set.Init(n)
}

// 追溯某个集合，直至头部，头部的下标即表示集合的标识符
func (s *Set) Trace(n int) int {
	i, v := n, (*s)[n]
	for v < n {
		n = v
		v = (*s)[n]
	}
	// 此操作是为了在扁平化追溯链条，在节点数量较大时节省查询时间
	(*s)[i] = n
	return n
}

// 合并两个集合，注意需保证i, j均为Trace返回的头部元素（索引下标）
func (s *Set) Merge(i, j int) {
	if i < j {
		(*s)[j] = i
	} else {
		(*s)[i] = j
	}
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

	//初始化集合
	set := NewSet(len(m.GetVertices()))

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

	//依次从小到大取出每一条边。
	//如果某条边它的两个顶点点分属不同集合，则记录该边，并合并两个集合；否则丢弃该边。
	for h.Len() > 0 {
		n := heap.Pop(h).(*Node)
		s := set.Trace(n.GetS())
		e := set.Trace(n.GetE())
		if s != e {
			fmt.Printf("%-2d => %2d: %.4f\n", n.GetS(), n.GetE(), n.GetD())
			set.Merge(s, e)
		}
	}

	return s
}
