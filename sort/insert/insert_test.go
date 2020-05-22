package insert

import (
	"math/rand"
	"testing"
	"time"
)

func TestSort(t *testing.T) {
	s := [][]int{
		{3, 0, 7, 8, 9, 4, 1, 2, 6, 5},
		{10, 17, 11, 13, 12, 14, 15, 16, 18, 20, 19},
	}
	var r []int
	for _, e := range s {
		r = append(r, Sort(e)...)
		//log.Println(r)
	}

	for i := 0; i < len(r); i++ {
		if r[i] != i {
			t.Errorf("expected %d,actual %d", i, r[i])
		}
	}
}

func TestSort2(t *testing.T) {
	s := [][]int{
		{3, 0, 7, 8, 9, 4, 1, 2, 6, 5},
		{10, 17, 11, 13, 12, 14, 15, 16, 18, 20, 19},
	}
	var r []int
	for _, e := range s {
		r = append(r, Sort2(e)...)
		//log.Println(r)
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
		r := rand.Intn(len(s))
		s[i] = r
	}
	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		Sort2(s)
	}
}
