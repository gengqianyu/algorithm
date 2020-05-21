package bubble

import (
	"math/rand"
	"testing"
	"time"
)

// 冒泡排序测试
func TestSort(t *testing.T) {
	s := [][]int{
		{3, 5, 7, 8, 9, 4, 1, 2, 6, 0},
		{10, 17, 11, 13, 12, 14, 15, 16, 18, 20, 19},
	}
	var r []int
	for _, e := range s {
		r = append(r, Sort(e)...)
	}

	for i := 0; i < len(r); i++ {
		if r[i] != i {
			t.Errorf("expected %d,actual %d", i, r[i])
		}
	}
}

//冒泡排序 性能测试
//D:\go\algorithm\sort\bubble>go test --bench=".*" --benchmem -v
//goos: windows
//goarch: amd64
//pkg: algorithm/sort/bubble
//BenchmarkBubbleSort-8                  1        3971727900 ns/op               0 B/op
//0 allocs/op
//PASS
//ok      algorithm/sort/bubble   4.904s
func BenchmarkSort(b *testing.B) {
	var s []int
	for i := 0; i < 80000; i++ {
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(80000)
		s = append(s, r)
	}
	b.ResetTimer()
	//因为是个排序算法不用循环b.N次，测试1次就ok
	for j := 0; j < 1; j++ {
		Sort(s)
	}
}
