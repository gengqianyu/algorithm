//二叉排序树BST
//BST：Binary Sort Tree 对于二叉排序树的任何一个非叶子节点，左子节点的值比当前节点的值小，右子节点的值比当前节点值大。
//如果值相等，左右节点都可以
package bstree

import (
	"errors"
	"strconv"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) Value() int {
	return n.value
}

func (n *Node) SetValue(v int) {
	n.value = v
}

func (n *Node) Left() *Node {
	return n.left
}
func (n *Node) SetLeft(e *Node) {
	n.left = e
}

func (n *Node) Right() *Node {
	return n.right
}

func (n *Node) SetRight(e *Node) {
	n.right = e
}

//add a element to node
func (n *Node) Add(e *Node) {
	if n == nil || e == nil {
		return
	}
	if e.value < n.value {
		if n.left == nil {
			n.left = e
		} else {
			n.left.Add(e)
		}
	} else {
		if n.right == nil {
			n.right = e
		} else {
			n.right.Add(e)
		}
	}
}

// traversing tree
func (n *Node) Traverse(f func(n *Node)) {
	if n == nil {
		return
	}
	n.left.Traverse(f)
	f(n)
	n.right.Traverse(f)
}

// find node from BSTree
func (n *Node) Find(v int) *Node {
	//如果相等，返回找到节点
	if n.value == v {
		return n
	}
	//如果节点值比要找的值大，同时满足节点的左子节点不为nil，就向左递归
	if n.value > v && n.left != nil {
		return n.left.Find(v)
	}
	//如果节点值比要找的值小，同时满足节点的右子节点不为nil，就向右递归
	if n.value < v && n.right != nil {
		return n.right.Find(v)
	}
	//如果左右节点为nil表示未找到直接返回nil
	return nil
}

// find parent node from BSTree
func (n *Node) Parent(v int) *Node {
	//如果一个节点的左右节点的值等于v，那这个节点就是要找节点的父节点
	if (n.left != nil && n.left.value == v) || (n.right != nil && n.right.value == v) {
		return n
	}
	//如果节点值比要找的值大，同时满足节点的左子节点不为nil，就向左递归
	if n.left != nil && n.value > v {
		return n.left.Parent(v)
	}
	//如果节点值比要找的值小，同时满足节点的右子节点不为nil，就向右递归
	if n.right != nil && n.value <= v {
		return n.right.Parent(v)
	}
	//如果左右节点为nil表示未找到直接返回nil
	return nil
}

// defined BD+STree struct
type BSTree struct {
	root   *Node
	number int
}

func (b *BSTree) init(n *Node) *BSTree {
	b.root = n
	b.number = 1
	return b
}

func New(n *Node) *BSTree {
	return new(BSTree).init(n)
}

func (b *BSTree) Root() *Node {
	return b.root
}

func (b *BSTree) Number() int {
	return b.number
}

//add a node to BSTree
func (b *BSTree) Add(n *Node) {
	if b.root == nil {
		b.init(n)
	} else {
		b.root.Add(n)
		b.number++
	}
}

func (b *BSTree) Traverse() <-chan *Node {
	out := make(chan *Node)
	go func() {
		b.root.Traverse(func(n *Node) {
			out <- n
		})
		close(out)
	}()
	return out
}

func (b *BSTree) Find(v int) (*Node, error) {
	if b.number == 0 {
		return nil, errors.New("the BSTree is empty")
	}
	if n := b.root.Find(v); n != nil {
		return n, nil
	} else {
		return nil, errors.New("not find node of value equal " + strconv.Itoa(v))
	}

}

func (b *BSTree) Parent(v int) (*Node, error) {
	if b.number == 0 {
		return nil, errors.New("the BSTree is empty")
	}
	if p := b.root.Parent(v); p != nil {
		return p, nil
	} else {
		return nil, errors.New("not find parent node of value equal " + strconv.Itoa(v))
	}
}

func (b *BSTree) Delete(v int) bool {
	//查找要删除的节点
	n, e := b.Find(v)
	if e != nil {
		return false
	}

	//如果BSTree只有一个节点，并且root节点就是要删除的节点直接删除
	if b.number == 1 && b.root.value == n.value {
		b.root = nil
		b.number--
		return true
	}
	//查找要删除节点的父节点
	p := b.root.Parent(v)

	//如果n为叶子节点，删除叶子节点
	if n.Left() == nil && n.Right() == nil {
		//n是父节点p的左子节点
		if p.Left() != nil && p.Left().Value() == n.Value() {
			p.SetLeft(nil)
			b.number--
			return true
		}
		//n是父节点p的右子节点
		if p.Right() != nil && p.Right().Value() == n.Value() {
			p.SetRight(nil)
			b.number--
			return true
		}
	}

	//如果删除的节点n只有一颗左子树
	if n.Left() != nil && n.Right() == nil {
		if p == nil {
			b.root = n.Left()
			b.number--
			return true
		}
		//n是父节p点的左子节点
		if p.Left() != nil && p.Left().Value() == n.Value() {
			p.SetLeft(n.Left())
			b.number--
			return true
		}
		//n是父节p点的右子节点
		if p.Right() != nil && p.Right().Value() == n.Value() {
			p.SetRight(n.Left())
			b.number--
			return true
		}
	}
	//如果删除的节点n只有一颗右子树
	if (n.Right() != nil) && (n.Left() == nil) {
		if p == nil {
			b.root = n.Right()
			b.number--
			return true
		}
		//n是父节点的左子节点
		if p.Left() != nil && p.Left().Value() == n.Value() {
			p.SetLeft(n.Right())
			b.number--
			return true
		}
		//n是父节点的右子节点
		if p.Right() != nil && p.Right().Value() == n.Value() {
			p.SetRight(n.Right())
			b.number--
			return true
		}
	}

	//如果删除的节点n，有左右两颗子树

	//从n的右子树查找一个权值最小的节点
	//将最小节点的权值用临时变量保存
	//将最小节点删除
	//将最小节点的值赋值给n

	//第二种方法是从左子树中找到最大的一个进行类似操作

	m := n.Right()
	//找到n右子树中最小的节点
	for m.Left() != nil {
		m = m.Left()
	}
	b.Delete(m.Value())
	n.SetValue(m.Value())
	return true
}
