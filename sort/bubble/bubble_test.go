package bubble

import (
	"math/rand"
	"testing"
	"time"
)

//冒泡排序 性能测试
//D:\go\algorithm\sort\bubble>go test --bench=".*" --benchmem -v
//goos: windows
//goarch: amd64
//pkg: algorithm/sort/bubble
//BenchmarkBubbleSort-8                  1        3971727900 ns/op               0 B/op
//0 allocs/op
//PASS
//ok      algorithm/sort/bubble   4.904s
func BenchmarkBubbleSort(b *testing.B) {
	var s []int
	for i := 0; i < 80000; i++ {
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(80000)
		s = append(s, r)
	}
	b.ResetTimer()
	//因为是个排序算法不用循环b.N次，测试1次就ok
	for j := 0; j < 1; j++ {
		BubbleSort(s)
	}
}
