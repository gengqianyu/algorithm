package htree

import (
	"math/rand"
	"testing"
	"time"
)

func TestHTree_Sort(t *testing.T) {
	s := [][]int{
		{3, 5, 7, 8, 9, 4, 1, 2, 6, 0},
		{10, 17, 11, 13, 12, 14, 15, 16, 18, 20, 19},
	}
	var r []int
	for _, e := range s {
		r = append(r, New().init(e).Sort()...)
	}

	for i := 0; i < len(r); i++ {
		//t.Log(r[i])
		if r[i] != i {
			t.Errorf("expected %d,actual %d", i, r[i])
		}
	}
}

//D:\go\algorithm\tree\htree>go test --bench=".*" --benchmem -v
//=== RUN   TestHTree_Sort
//--- PASS: TestHTree_Sort (0.00s)
//goos: windows
//goarch: amd64
//pkg: algorithm/tree/htree
//BenchmarkHTree_Sort
//BenchmarkHTree_Sort-8               2920            416599 ns/op               0 B/op          0 allocs/op
//PASS
//ok      algorithm/tree/htree    1.679s

func BenchmarkHTree_Sort(b *testing.B) {
	s := make([]int, 8000)
	for i := 0; i < len(s); i++ {
		rand.Seed(time.Now().UnixNano())
		s[i] = rand.Intn(len(s))
	}
	h := New().init(s)

	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		h.Sort()
	}
}
