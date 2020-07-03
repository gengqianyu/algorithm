package btree

import (
	"fmt"
	"reflect"
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
	btree := New()
	root := btree.Root()
	root.value = hero{1, "宋江"}
	root.left = &Node{value: hero{2, "吴用"}}
	root.right = &Node{value: hero{3, "公孙胜"}}
	root.right.left = &Node{value: hero{6, "卢俊义"}}
	root.right.right = &Node{value: hero{7, "鲁智深"}}
	root.left.left = &Node{value: hero{4, "林冲"}}
	root.left.right = &Node{value: hero{5, "武松"}}

	out := btree.Traverse(POST)

	for node := range out {
		if v, ok := node.Value().(hero); ok {
			fmt.Println(v)
		}

	}
	fmt.Println("----------------------------------------")

	fieldName := "id"
	fieldValue := 7
	node, err := btree.Select(fieldName, fieldValue, POST)

	if err != nil {
		t.Error(err.Error())
	}

	h := node.value.(hero)
	if h.name != "鲁智深" {
		t.Errorf("expected:%s,actual:%s", "鲁智深", h.name)
	}

	id := 5
	node, err = btree.SelectById(id, func(value reflect.Value, id interface{}) bool {
		v, ok := value.Interface().(hero)
		if !ok {
			return false
		}
		if reflect.DeepEqual(v.id, id) {
			return true
		}
		return false
	}, POST)

	if err != nil {
		t.Error(err.Error())
	}

	h = node.value.(hero)
	if h.name != "武松" {
		t.Errorf("expected:%s,actual:%s", "武松", h.name)
	}

}
