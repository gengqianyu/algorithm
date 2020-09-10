package prim

//定义图的边
type edge struct {
	s, e int     //起点终点
	d    float64 //距离
}

//堆节点
type Node struct {
	edge
	p, l, r *Node //父节点，左子节点，右子节点
}

//定义一个小堆
type MHeap []*Node

//实现 heap.Interface 的Len方法
func (h *MHeap) Len() int {
	return len(*h)
}

//实现 heap.Interface 的Less方法
func (h *MHeap) Less(i, j int) bool {
	return (*h)[i].d < (*h)[j].d
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
