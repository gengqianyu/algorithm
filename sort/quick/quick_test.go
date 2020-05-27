package quick

import (
	"log"
	"math/rand"
	"testing"
	"time"
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

func BenchmarkSort(b *testing.B) {
	s := make([]int, 8000)
	for i := 0; i < 8000; i++ {
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(8000)
		s[i] = r
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sort(s, 0, len(s)-1)
	}
}
