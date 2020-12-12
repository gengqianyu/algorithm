package avltree

import (
	"errors"
	"log"
	"math"
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

// traversing AVLTree
func (n *Node) Traverse(f func(n *Node)) {
	if n == nil {
		return
	}
	n.left.Traverse(f)
	f(n)
	n.right.Traverse(f)
}

// find node from AVLTree
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

// find parent node from AVLTree
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

//返回以当前节点n,为root节点树的高度
//利用递归来求解树的高度，就是一个计数，发现一个节点就加+1
func (n *Node) Height() int {
	//从左右子树中找到一个最高的，然后加上root本身就是root的树高。
	return int(math.Max(float64(n.LeftHeight()), float64(n.RightHeight()))) + 1
}

//左子树高度
func (n *Node) LeftHeight() int {
	if n.left != nil {
		return n.left.Height()
	}
	return 0
}

//右子树高度
func (n *Node) RightHeight() int {
	if n.right != nil {
		return n.right.Height()
	}
	return 0
}

//旋转
func (n *Node) Rotate() {
	//获取左右子树高度
	lHeight, rHeight := n.LeftHeight(), n.RightHeight()
	//如果右子树比左子树高
	if rHeight-lHeight == 2 {
		//右子树的右子树，比右子树的左子树高，直接对root节点n进行左转，
		//否则先让root节点n的左子树先进行右转，再对root节点n左转
		if n.right.RightHeight() > n.right.LeftHeight() {
			log.Println("左璇")
			n.LeftRotate()
		} else {
			log.Println("先右璇，再左旋")
			n.RightThenLeftRotate()
		}
		return
	}
	//如果左子树比右子树高
	if lHeight-rHeight == 2 {
		if n.left.LeftHeight() > n.left.RightHeight() {
			log.Println("右旋")
			n.RightRotate()
		} else {
			log.Println("先左璇，再右旋")
			n.LeftThenRightRotate()
		}
		return
	}
}

//左旋
func (n *Node) LeftRotate() {
	//以当前root节点的值，创建一个新节点node
	node := &Node{value: n.value}
	//把新结点node的左子树，设置成root节点的左子树
	node.left = n.left
	//把新结点node的右子树，设置成root节点右子树的左子树
	node.right = n.right.left
	//把当前root节点的值修改成右子节点的值
	n.value = n.right.value
	//把当前root节点的右子树，设置成右子节点的右子树
	n.right = n.right.right
	//把当前节点root的左子节点，设置成新节点node
	n.left = node
}

//右旋
func (n *Node) RightRotate() {
	//创建一个新结点。新结点的值为root节点的值
	node := &Node{value: n.value}
	//让node节点的右子节点指向，根节点n的右节点
	node.right = n.right
	//让node节点的左子节点指向,根节点n的左子节点的右子树
	node.left = n.left.right
	//让root节点n的值等于，root节点的左子节点的值
	n.value = n.left.value
	//让root节点n的左子节点指向，根节点n的左子节点的左子树
	n.left = n.left.left
	//让root节点n的右节点，指向新结点node
	n.right = node
}

//左旋之后，右旋
func (n *Node) LeftThenRightRotate() {
	//以失衡点n，左子结点先左旋转
	n.left.LeftRotate()
	//再以失衡点右旋
	n.RightRotate()
}

//右璇之后，左旋
func (n *Node) RightThenLeftRotate() {
	//以失衡点n，右子结点先右旋转
	n.right.RightRotate()
	//再以失衡点n，左旋转
	n.LeftRotate()
}

// defined BD+STree struct
type AVLTree struct {
	root   *Node
	number int
}

func (a *AVLTree) init(n *Node) *AVLTree {
	a.root = n
	a.number = 1
	return a
}

func New(n *Node) *AVLTree {
	return new(AVLTree).init(n)
}

func (a *AVLTree) Root() *Node {
	return a.root
}

func (a *AVLTree) Number() int {
	return a.number
}

//add a node to AVLTree
func (a *AVLTree) Add(n *Node) {
	if a.root == nil {
		a.init(n)
	} else {
		a.root.Add(n)
		a.root.Rotate()
		a.number++
	}
}

func (a *AVLTree) Traverse() <-chan *Node {
	out := make(chan *Node)
	go func() {
		//Traverse递归全部完成，才关闭通道out
		a.root.Traverse(func(n *Node) {
			out <- n
			//close(out)//如果在函数内关闭通道，第一个发送完之后，后面的递归数据就没有通道可发了
		})
		close(out)
	}()
	return out
}

func (a *AVLTree) Find(v int) (*Node, error) {
	if a.number == 0 {
		return nil, errors.New("the AVLTree is empty")
	}
	if n := a.root.Find(v); n != nil {
		return n, nil
	} else {
		return nil, errors.New("not find node of value equal " + strconv.Itoa(v))
	}

}

func (a *AVLTree) Parent(v int) (*Node, error) {
	if a.number == 0 {
		return nil, errors.New("the AVLTree is empty")
	}
	if p := a.root.Parent(v); p != nil {
		return p, nil
	} else {
		return nil, errors.New("not find parent node of value equal " + strconv.Itoa(v))
	}
}

func (a *AVLTree) Delete(v int) bool {
	//查找要删除的节点
	n, e := a.Find(v)
	if e != nil {
		return false
	}

	//如果AVLTree只有一个节点，并且root节点就是要删除的节点直接删除
	if a.number == 1 && a.root.value == n.value {
		a.root = nil
		a.number--
		return true
	}
	//查找要删除节点的父节点
	p := a.root.Parent(v)

	//如果n为叶子节点，删除叶子节点
	if n.Left() == nil && n.Right() == nil {
		//n是父节点p的左子节点
		if p.Left() != nil && p.Left().Value() == n.Value() {
			p.SetLeft(nil)
			a.root.Rotate()
			a.number--
			return true
		}
		//n是父节点p的右子节点
		if p.Right() != nil && p.Right().Value() == n.Value() {
			p.SetRight(nil)
			a.root.Rotate()
			a.number--
			return true
		}
	}

	//如果删除的节点n只有一颗左子树
	if n.Left() != nil && n.Right() == nil {
		if p == nil {
			a.root = n.Left()
			a.number--
			return true
		}
		//n是父节p点的左子节点
		if p.Left() != nil && p.Left().Value() == n.Value() {
			p.SetLeft(n.Left())
			a.root.Rotate()
			a.number--
			return true
		}
		//n是父节p点的右子节点
		if p.Right() != nil && p.Right().Value() == n.Value() {
			p.SetRight(n.Left())
			a.root.Rotate()
			a.number--
			return true
		}
	}
	//如果删除的节点n只有一颗右子树
	if (n.Right() != nil) && (n.Left() == nil) {
		if p == nil {
			a.root = n.Right()
			a.number--
			return true
		}
		//n是父节点的左子节点
		if p.Left() != nil && p.Left().Value() == n.Value() {
			p.SetLeft(n.Right())
			a.root.Rotate()
			a.number--
			return true
		}
		//n是父节点的右子节点
		if p.Right() != nil && p.Right().Value() == n.Value() {
			p.SetRight(n.Right())
			a.root.Rotate()
			a.number--
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
	a.Delete(m.Value())
	n.SetValue(m.Value())
	return true
}

func (a *AVLTree) Height() int {
	if a.number == 0 {
		return 0
	}
	return a.root.Height()
}
