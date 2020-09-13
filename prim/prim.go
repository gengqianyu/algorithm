package prim

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

//清空堆
func (h *MHeap) clear() {
	for h.Len() > 0 {
		h.Pop()
	}
}

//普里姆算法，
//1寻找图中任意顶点，以它为起点，它的所有边n加入集合(优先队列)h,设置一个boolean数组bool[]标记该顶点是否已访问。
//2从集合h中找到距离最小的那个边t，并判断，边t的终点顶点e是否被访问过， 如果终点e被访问过那么直接跳过此顶点；继续寻找下一个最小边。
// 如果未被标(访问)记那么标记该终点e,并且将与该终点顶点e相连的关系顶点(未被标记)构成的边加入集合(优先队列)h，边t就是此轮中要找的最小边.
//重复1，2步骤，直到h优先队列为空，构成最小生成树 ！
func Prim(m *graph.GoMap) (s []*Node) {
	var i, j int                         //i当前顶点索引，j表示当前顶点的关系顶点索引
	have := make([]bool, len(m.Edges())) //用来标记顶点是否已访问

	h := new(MHeap) //定义一个小堆，模拟优先队列
	heap.Init(h)    //堆化

	for {
		//初始化当前起始顶点为已访问
		have[i] = true

		//以i索引对应的顶点为起始顶点，将和起始顶点连接的其他顶点全部放入堆（优先队列）中
		vertxNum := len(m.GetVertices())
		for j = 0; j < vertxNum; j++ {
			//起始顶点i到关系顶点j未访问过，就加入优先队列 ，否则直接不处理
			if have[j] == false {
				v := m.Edges()[i][j]
				n := new(Node)
				n.SetS(i)
				n.SetE(j)
				n.SetD(v)
				heap.Push(h, n)
			}
		}
		//顶点都已经访问h就不会添加数据了
		if h.Len() == 0 {
			break
		}

		//获取最小顶点
		t := heap.Pop(h).(*Node)
		//如果最小边t的的终点e已经被访问过了就直接跳过，否则提取该边作为，本轮寻找的最小边保存
		if have[t.GetE()] == true {
			continue
		} else {
			//保存每一轮最小边节点
			s = append(s, t)
			//将目的边距最小的顶点设置为起始顶点，继续找下一轮的目的顶点
			i = t.GetE()
			//fmt.Printf("%2d -> %2d: %.4f\n", t.GetS(), t.GetE(), t.GetD())
		}
	}
	return s
}
