package btree

type Node struct {
	value       interface{}
	left, right *Node
}

func (n *Node) Value() interface{} {
	return n.value
}

func (n *Node) Left() *Node {
	return n.left
}

func (n *Node) Right() *Node {
	return n.right
}

func (n *Node) Traverse(order uint8, c container) {
	*c = append(*c, n.value)
	n.left.Traverse(order, c)
	n.right.Traverse(order, c)
}

func (n *Node) TraverseByClosure(order uint8, f NodeFunc) {
	if n == nil {
		return
	}

	if order == 0 {
		f(n)
	}

	n.left.TraverseByClosure(order, f)

	if order == 1 {
		f(n)
	}

	n.right.TraverseByClosure(order, f)

	if order == 2 {
		f(n)
	}

}

type NodeFunc func(node *Node)

type Btree struct {
	root *Node // root node
	len  int   //number of nodes
	high int   //layers
}

func New() *Btree {
	return new(Btree).init()
}

func (b *Btree) init() *Btree {
	b.root = new(Node)
	b.len = 0
	b.high = 0
	return b
}

//numbers of nodes in binary tree
func (b *Btree) Len() int {
	return b.len
}

func (b *Btree) insert(n *Node) {

}

// layers of Btree
func (b *Btree) High() int {
	return b.high
}

func (b *Btree) Root() *Node {
	return b.root
}

// defied enum of traverse order
const (
	PRE = iota
	POST
	IN
)

// traverse binary tree
func (b *Btree) Traverse(order uint8) <-chan *Node {
	out := make(chan *Node)

	o := order
	go func(o uint8) {
		b.root.TraverseByClosure(o, func(n *Node) {
			out <- n
		})
		close(out)
	}(o)

	return out
}

type container *[]interface{}
