//赫夫曼树
package hmtree

type Node struct {
	weight      int   //权值
	left, right *Node //左右子节点
}

func (n *Node) init(v int) *Node {
	n.weight = v
	return n
}

func NewNode(v int) *Node {
	return new(Node).init(v)
}

func (n *Node) Left() *Node {
	return n.left
}

func (n *Node) Right() *Node {
	return n.right
}

func (n *Node) SetLeft(e *Node) {
	n.left = e
}

func (n *Node) SetRight(e *Node) {
	n.right = e
}

func (n *Node) Weight() int {
	return n.weight
}

func (n *Node) Traverse(f func(e *Node)) {
	if n == nil {
		return
	}
	n.Left().Traverse(f)
	f(n)
	n.Right().Traverse(f)
}

type HMTree struct {
	root      *Node
	container []*Node
	len       int
}

func New(c []int) *HMTree {
	return new(HMTree).init(c)
}

func (h *HMTree) init(c []int) *HMTree {
	for _, e := range c {
		h.container = append(h.container, NewNode(e))
	}
	h.len = len(h.container)
	return h
}

func (h *HMTree) Len() int {
	return h.len
}

func (h *HMTree) Create() *HMTree {
	for len(h.container) > 1 {
		//排序
		h.Sort(0, len(h.container)-1)
		//左右子节点
		l, r := h.container[0], h.container[1]
		h.container = h.container[2:]
		//新父节点
		w := l.Weight() + r.Weight()
		n := NewNode(w)
		n.SetLeft(l)
		n.SetRight(r)
		h.container = append(h.container, n)
		h.len++
	}
	h.root = h.container[0]
	return h
}

func (h *HMTree) Traverse() <-chan *Node {
	out := make(chan *Node)
	go func() {
		h.root.Traverse(func(n *Node) {
			out <- n
		})
		close(out)
	}()
	return out
}

func (h *HMTree) Sort(l, r int) []*Node {
	//中轴值
	m := h.container[l]
	//中轴位置
	p := l
	//定义before 和after 前后两个起始位置
	b, a := l, r
	for b <= a {
		//找右边一组找到一个比中轴值小的权值
		for a >= p && h.container[a].Weight() >= m.Weight() {
			a--
		}
		if a >= p {
			h.container[p] = h.container[a]
			p = a
		}
		//找左边一组，找到一个比中轴值大的权值
		for b <= p && h.container[b].Weight() <= m.Weight() {
			b++
		}
		if b <= p {
			h.container[p] = h.container[b]
			p = b
		}
	}
	//这里的p的位置已经是b的位置了，将m
	h.container[p] = m
	//中值不动，两边切片递归排序
	//左递归
	if p-l > 1 {
		h.Sort(0, p-1)
	}
	//右递归
	if r-p > 1 {
		h.Sort(p+1, r)
	}
	return h.container
}
