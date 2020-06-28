package btree

import (
	"fmt"
	"testing"
)

type hero struct {
	id   uint8
	name string
}

var marshalTests []hero

func TestBtree_Traverse(t *testing.T) {
	btree := New()
	root := btree.Root()
	root.value = hero{1, "宋江"}
	root.left = &Node{value: hero{2, "吴用"}}
	root.right = &Node{value: hero{3, "公孙胜"}}
	root.right.left = &Node{value: hero{6, "卢俊义"}}
	root.right.right = &Node{value: hero{7, "鲁智深"}}
	root.left.left = &Node{value: hero{4, "林冲"}}
	root.left.right = &Node{value: hero{5, "武松"}}

	out := btree.Traverse(PRE)

	for node := range out {
		if v, ok := node.Value().(hero); ok {
			fmt.Println(v)
		}

	}
}
