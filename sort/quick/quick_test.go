package quick

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	s := [][]int{
		{3, 5, 7, 8, 9, 4, 1, 2, 6, 0},
		{10, 17, 11, 13, 12, 14, 15, 16, 18, 20, 19},
	}
	var r []int
	for _, e := range s {
		r = append(r, Sort(e, 0, len(e)-1)...)
	}
	for _, e := range r {
		log.Println(e)
	}

	for i := 0; i < len(r); i++ {
		if r[i] != i {
			t.Errorf("expected %d,actual %d", i, r[i])
		}
	}
}
