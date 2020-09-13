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

//在计算机科学中，并查集是一种树型的数据结构，用于处理一些不交集（Disjoint Sets）的合并及查询问题。
//有一个联合-查找算法（union-find algorithm）定义了两个用于此数据结构的操作：
//Find：确定元素属于哪一个子集。它可以被用来确定两个元素是否属于同一子集。
//Union：将两个子集合并成同一个集合。
//为了更加精确的定义这些方法，需要定义如何表示集合。一种常用的策略是为每个集合选定一个固定的元素，称为代表，以表示整个集合。
//接着，Find(x) 返回x所属集合的代表，而Union使用两个集合的代表作为参数。

//定义一个不交集，（合并查找集合）
type Set []int

//创建一个所有顶点集的总集，因为每个元素都等于下标，即都是单元素集合
func (s *Set) Init(n int) *Set {
	for n--; n >= 0; n-- {
		(*s)[n] = n
	}
	return s
}

//创建单元素集合
func MakeSet(n int) *Set {
	set := Set(make([]int, n))
	return set.Init(n)
}

//确定元素属于哪一个子集。它可以被用来确定两个元素是否属于同一子集。
func (s *Set) Find(e int) int {
	i := e
	v := (*s)[e]

	for v < e {
		e = v
		v = (*s)[e]
	}

	(*s)[i] = e
	return e
}

// 将两个子集合并成同一个集合
func (s *Set) Union(i, j int) {
	if i < j {
		(*s)[j] = i
	} else {
		(*s)[i] = j
	}
}

//利用克鲁斯卡尔算法解决最短路径问题
//Kruskal算法进行调度的单位是边,它的信仰为:所有边能小则小，算法的实现方面和并查集(不相交集合)很像，要用到并查集判断两点是否在同一集合。
//1,将边(以及2顶点)的对象依次加入集合(优先队列)q1中。初始所有点相互独立。

//2,取出当前q1最小边，判断边的两点是否联通。
//3,如果联通，跳过，如果不连通，则使用union（并查集合并）将两个顶点合并。这条边被使用(可以储存或者计算数值)。

//重复2，3操作直到集合（优先队列）q1为空。此时被选择的边构成最小生成树。
func Kruskal(m *graph.GoMap) (s []*Node) {
	//初始化最小堆
	h := new(MHeap)
	heap.Init(h)

	//初始化集合
	set := MakeSet(len(m.GetVertices()))

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

	//依次从小到大取出每一条边,判断边的两点是否联通。
	//如果某条边它的两个顶点点分属不同集合，则记录该边，并,合并两个集合；否则丢弃该边。
	for h.Len() > 0 {
		n := heap.Pop(h).(*Node)
		p := set.Find(n.GetS())
		q := set.Find(n.GetE())
		//判断边的两点是否联通。
		if p != q {
			fmt.Printf("%2d => %2d: %.4f\n", n.GetS(), n.GetE(), n.GetD())
			set.Union(p, q)
		}
	}

	return s
}
