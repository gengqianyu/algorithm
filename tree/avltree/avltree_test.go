package avltree

import (
	"log"
	"testing"
)

func TestAVLTree_Height(t *testing.T) {
	//
	//s := []int{10, 11, 7, 6, 8, 9}
	//
	//s := []int{10, 11, 7, 6, 8, 9}
	//先左旋，再右旋
	//s := []int{10, 11, 7, 6, 8, 9}
	//先右旋，再左旋
	s := []int{2, 1, 6, 5, 7, 3}
	AVL := New(nil)
	for _, e := range s {
		AVL.Add(&Node{value: e})
	}

	for n := range AVL.Traverse() {
		t.Log(n.Value())
	}

	log.Println(AVL.Height())
	log.Println(AVL.Root().LeftHeight())
	log.Println(AVL.Root().RightHeight())
}
