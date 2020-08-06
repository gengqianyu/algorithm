package bstree

import (
	"log"
	"testing"
)

func TestBSTree_Traverse(t *testing.T) {
	//这是一个人为的平衡树
	s := []int{7, 3, 10, 12, 11, 5, 1, 9, 2}
	bst := New(nil)
	for _, e := range s {
		bst.Add(&Node{value: e})
	}
	for n := range bst.Traverse() {
		t.Log(n.Value())
	}
	log.Println()

	//log.Println(bst.Parent(9))

	//删除叶子节点
	//bst.Delete(5)
	//bst.Delete(9)
	//bst.Delete(12)

	//删除只有一个子树节点的节点
	//bst.Delete(1)

	//删除有两个子节点的节点
	//bst.Delete(10)

	//删除root节点
	//bst.Delete(7)

	//全部删除测试
	bst.Delete(7)
	bst.Delete(3)
	bst.Delete(12)
	bst.Delete(11)
	bst.Delete(5)
	bst.Delete(9)
	bst.Delete(2)
	bst.Delete(10)
	bst.Delete(1)
	for n := range bst.Traverse() {
		t.Log(n.Value())
	}
	log.Println(bst.Number())
}
