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

//在计算机科学中，并查集是一种树型的数据结构，用于处理一些不交集（Disjoint Sets）的合并及查询问题。
//有一个联合-查找算法（union-find algorithm）定义了两个用于此数据结构的操作：
//Find：确定元素属于哪一个子集。它可以被用来确定两个元素是否属于同一子集。
//Union：将两个子集合并成同一个集合。
//为了更加精确的定义这些方法，需要定义如何表示集合。一种常用的策略是为每个集合选定一个固定的元素(老大)，称为代表，以表示整个集合。
//接着，Find(x) 返回x所属集合的代表，而Union使用两个集合的代表作为参数。

//定义一个不交集，（合并查找集合）并查集
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
func (s *Set) Find(i int) int {
	//第一次i(索引)元素的老大就是他自己，经过几次合并之后，就有可能别的索引x也指向了索引i指向的值, 也就是x认i做老大
	//i索引代表小弟，(*s)[i]代表老大
	r := (*s)[i] //找到i的老大r
	//找root老大，索引i小弟 和老大r相等的时候就说明找到老大了（自己指向自己），否则将i的老大r,作为小弟递归继续找
	if r == i {
		return i
	} else {
		return s.Find(r)
	}
}

//功能同s.Find 非递归实现，节点非常多的话效率会很高
func (s *Set) Trace(i int) int {
	//p记录最小小弟
	//i索引代表小弟，(*s)[i]代表老大
	p, r := i, (*s)[i]
	//如果r老大比i小，说明i需要用自己的老大r作为小弟继续去找老大，不停的重置小弟i和老大r，知道i==r，老大找到
	for r < i {
		i = r       //重置小弟
		r = (*s)[i] //重置老大
	}
	//进行路径压缩,让最小小弟直接拜老大，减小下次寻找层级 ，可以用for重做上面步骤进行每一层的路径压缩
	(*s)[p] = i // 此操作是为了在扁平化并查集，在节点数量较大时节省查询时间，就是让孙子辈的节点直接拜老大
	return i
}

// 将两个子集合并成同一个集合
func (s *Set) Union(i, j int) {
	//让j集合拜i集合为老大
	//因为所有边是从小到大处理的，所以i，j谁小谁是老大
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
//3,如果联通，跳过，如果不连通，则使用union（并查集合并）将两个顶点合并。提取或处理此边(可以储存或者计算数值)。

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
		//通过两个顶点，分别去找老大，判断两个顶点在图中是否有回路
		p := set.Find(n.GetS())
		q := set.Find(n.GetE())
		//判断边的两点是否联通。如此不连通，就记录边，合并两个不交集
		//如果双方找到的老大不是同一个，就代表图的这两个顶点没有回路，记录边。否则老大是一个就不处理，说明有回路
		if p != q {
			s = append(s, n)
			//fmt.Printf("%2d => %2d: %.4f\n", n.GetS(), n.GetE(), n.GetD())
			//拜老大
			set.Union(p, q)
		}
	}

	return s
}
