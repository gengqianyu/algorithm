package binary

import (
	"algorithm/sort/quick"
	"log"
	"testing"
)

func TestSearch(t *testing.T) {
	s := [][]int{
		{3, 5, 7, 8, 9, 4, 1, 2, 6, 0},
		{10, 17, 11, 13, 12, 14, 15, 16, 18, 20, 19, 31, 56},
	}
	var r []int
	for _, e := range s {
		r = append(r, quick.Sort(e, 0, len(e)-1)...)
	}

	//递归法
	//i := Search(r, 0, len(r)-1, 31)

	//非递归法
	i := Search2(r, 31)
	if i == -1 || i != 21 {
		t.Errorf("expected:%d;actual:%d", 21, i)
	} else {
		log.Printf("index:%d,value:%d", i, r[i])
	}

}
