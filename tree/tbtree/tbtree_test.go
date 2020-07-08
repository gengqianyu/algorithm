package tbtree

import (
	"testing"
)

type hero struct {
	id   int
	name string
}

var marshalTests []hero

//D:\go\algorithm\tree\btree>go test -v(详情)
//=== RUN   TestBtree_Traverse
//{1 宋江}
//{2 吴用}
//{4 林冲}
//{5 武松}
//{3 公孙胜}
//{6 卢俊义}
//{7 鲁智深}
//--- PASS: TestBtree_Traverse (0.00s)
//PASS
//ok      algorithm/tree/btree    0.187s

func TestBtree_Traverse(t *testing.T) {
	tBtree := New()
	root := tBtree.Root()
	root.value = hero{1, "宋江"}
	root.left = &Node{value: hero{2, "吴用"}}
	root.right = &Node{value: hero{3, "公孙胜"}}
	root.right.left = &Node{value: hero{6, "卢俊义"}}
	root.right.right = &Node{value: hero{7, "鲁智深"}}
	root.left.left = &Node{value: hero{4, "林冲"}}
	root.left.right = &Node{value: hero{5, "武松"}}

	tBtree.Thread(root)
	h, ok := root.Left().Left().Right().Value().(hero)
	if !ok {
		t.Errorf("type error")
	}
	if h.name != "吴用" {
		t.Errorf("excpected:%s,actual:%s", "吴用", h.name)
	}

	tBtree.Traverse()
}
