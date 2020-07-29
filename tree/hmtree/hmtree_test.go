package hmtree

import (
	"log"
	"testing"
)

func TestHMTree_Traverse(t *testing.T) {
	s := []int{13, 7, 8, 3, 29, 6, 1}
	tree := New(s)
	tree.Create()
	log.Println(tree.len)
	log.Println(tree.root.Weight())
	for e := range tree.Traverse() {
		t.Log(e.Weight())
	}
}
