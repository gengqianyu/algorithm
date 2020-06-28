package btree

import (
	"reflect"
)

// tree node
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
	if n == nil {
		return
	}
	if order == PRE {
		*c = append(*c, n.value)
	}
	n.left.Traverse(order, c)
	if order == IN {
		*c = append(*c, n.value)
	}
	n.right.Traverse(order, c)
	if order == POST {
		*c = append(*c, n.value)
	}
}

// traverse binary tree with closure func
func (n *Node) TraverseByClosure(order uint8, f NodeFunc) {
	if n == nil {
		return
	}

	if order == PRE {
		f(n)
	}

	n.left.TraverseByClosure(order, f)

	if order == IN {
		f(n)
	}

	n.right.TraverseByClosure(order, f)

	if order == POST {
		f(n)
	}

}

type NodeFunc func(node *Node)

// binary tree
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
	PRE  = iota //前序
	IN          //中序
	POST        //后序
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

func (b *Btree) SelectById(order uint8, id uint8, f TypeAssertionFunc) *Node {
	out := b.Traverse(order)
	for node := range out {
		if f(reflect.ValueOf(node.Value()), id) {
			return node
		}
	}
	return nil
}

type container *[]interface{}

type TypeAssertionFunc func(v reflect.Value, id uint8) bool
