package insert

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	s := [][]int{
		{3, 0, 7, 8, 9, 4, 1, 2, 6, 5},
		{10, 17, 11, 13, 12, 14, 15, 16, 18, 20, 19},
	}
	var r []int
	for _, e := range s {
		r = append(r, Sort(e)...)
		log.Println(r)
	}

	for i := 0; i < len(r); i++ {
		if r[i] != i {
			t.Errorf("expected %d,actual %d", i, r[i])
		}
	}
}