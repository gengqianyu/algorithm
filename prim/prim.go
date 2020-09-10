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

//普雷姆算法，是在图的关系矩阵中，从一个顶点开始，找边距离最小的目的顶点，
//找到目的顶点后，在一此目的顶点作为起始顶点继续寻找下一个目的顶点，
//以此类推，直到所有顶点被访问过。退出。
func Prim(m *graph.GoMap) (s []*Node) {
	var i, j int                         //i当前顶点索引，j表示当前顶点的关系顶点索引
	have := make([]byte, len(m.Edges())) //用来记录
	h := new(MHeap)                      //定义一个小堆

	for {
		//初始化当前顶点为已访问
		have[i] = 1
		//以i索引对应的顶点为起始顶点，将和起始顶点连接的其他顶点全部放入堆中，然后找到一个最小顶点，这就是要找的目的顶点
		vertxNum := len(m.GetVertices())

		for j = 0; j < vertxNum; j++ {
			//have[j]为零，就表示i索引对应的起始节点到j索引对应的关系顶点未访问过，就处理
			if have[j] == 0 {
				v := m.Edges()[i][j]
				n := new(Node)
				n.SetS(i)
				n.SetE(j)
				n.SetD(v)
				h.Push(n)
			}
		}

		//顶点都已经访问h就不会添加数据了
		if h.Len() == 0 {
			break
		}
		//堆化
		heap.Init(h)
		//获取最小顶点
		t := heap.Pop(h).(*Node)
		//保存每一轮最小边节点
		s = append(s, t)
		//将目的边距最小的顶点设置为起始顶点，继续找下一轮的目的顶点
		i = t.GetE()

		//fmt.Printf("%2d -> %2d: %.4f\n", t.GetS(), t.GetE(), t.GetD())
		//清空小堆
		h.clear()
	}
	return s
}
