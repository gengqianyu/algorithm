//thread binary tree
//线索化二叉树
package tbtree

import (
	"fmt"
)

// tree node
type Node struct {
	value        interface{}
	left, right  *Node
	lType, rType bool
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

type NodeFunc func(node *Node)

// binary tree
type TBtree struct {
	root *Node // root node
	len  int   //number of nodes
	high int   //layers
	pre  *Node
}

func New() *TBtree {
	return new(TBtree).init()
}

func (b *TBtree) init() *TBtree {
	b.root = new(Node)
	b.len = 7
	b.high = 0
	return b

}

//numbers of nodes in binary tree
func (b *TBtree) Len() int {
	return b.len
}

func (b *TBtree) insert(n *Node) {

}

// layers of TBtree
func (b *TBtree) High() int {
	return b.high
}

func (b *TBtree) Root() *Node {
	return b.root
}

//线索化树节点 中序线索化
func (b *TBtree) Thread(n *Node) {
	if n == nil {
		fmt.Println(1)
		return
	}
	//线索化左子树
	//第一次弹出（左子树递归栈）顶的Thread函数执行，参数是树左3的林冲所在树节点，（其实是第二次因为第一次林冲的left是nil，执行会被直接返回）
	//第二次弹出（右子树递归栈）顶的Thread函数执行，参数林冲的right为nil 直接被返回
	//第三次弹出（左子树递归栈）顶的Thread函数执行，参数是树左2的吴用所在树节点
	//第四次弹出（右子树递归栈）顶的Thread函数执行，参数是武松所在树节点
	//第五次弹出（左子树递归栈）顶的Thread函数执行，参数是宋江所在树节点
	b.Thread(n.left)
	fmt.Println(*n)
	//线索化当前节点
	//处理当前节点的前驱节点
	//第一次 林冲满足条件 处理林冲left 为nil
	//第三次 吴用不满足条件 跳过
	//第四次 武松满足条件 处理武松left 为吴用
	if n.left == nil {
		n.left = b.pre
		n.lType = true
	}
	//处理当前节点的后驱节点 其实是在下一次递归的时候执行
	//第一次pre为nil不处理此逻辑
	//第三次pre为林冲，处理林冲的right 为吴用
	//第四次pre为吴用 吴用的right不为nil 跳过
	if b.pre != nil && b.pre.right == nil {
		b.pre.right = n
		b.pre.rType = true
	}
	//重置pre
	//第一次被重置成林冲
	//第三次被重置成吴用
	//第四次pre被重置成武松
	b.pre = n
	//线索化右子树
	//第一次林冲的right为nil 右子树递归栈弹栈执行 参数也就nil 直接被返回  左递归栈 弹栈 递归栈顶变成吴用来执行
	//第三次吴用的right为武松  右子树递归栈弹栈执行 参数也就武松
	//第四次武松的right为nil 右子树递归栈弹栈执行 参数也就nil 直接被返回 左递归栈 弹栈 递归栈顶变成吴用来执行
	b.Thread(n.right)
}

//遍历线索化二叉树
func (b *TBtree) Traverse() {
	currentNode := b.root
	//第一轮从root 宋江开始查找
	//第二轮从 武松开始查找
	for currentNode != nil {
		//查找 node.lType 为true的节点 找到线索化节点头
		//第一轮循环找到林冲 林冲的lType是true
		//第二轮循环武松lType是true所以直接跳过下面循环currentNode还是武松
		for !currentNode.lType {
			currentNode = currentNode.left
		}

		//找到头节点以后就打印
		//第一轮打印林冲
		//第二轮打印武松
		fmt.Println(*currentNode)
		//如果当前节点的右指针是后继节点，就输出
		//第一轮林冲的rType是true 林冲的right指向吴用 打印吴用
		//第二轮武松的rType为true 武松的 right指向宋江 打印宋江
		for currentNode.rType {
			//	获取当前节点的后继节点
			currentNode = currentNode.right
			fmt.Println(*currentNode)
		}
		// 第一轮最后将吴用的right 武松作为下一轮的开始节点
		// 第一轮最后将宋江的right 公孙胜作为下一轮的开始节点 以此类推
		currentNode = currentNode.right
	}
}
