package stack

import (
	"math/rand"
	"testing"
	"time"
)

var answer = []string{
	"php",
	"go",
	"c++",
	"js",
	"java",
	"dd",
	"css",
	"html",
}

//基准测试的代码文件必须以_test.go结尾
//基准测试的函数必须以Benchmark开头，必须是可导出的
//基准测试函数必须接受一个指向Benchmark类型的指针作为唯一参数
//基准测试函数不能有返回值
//b.ResetTimer是重置计时器，这样可以避免for循环之前的初始化代码的干扰
//最后的for循环很重要，被测试的代码要放到循环里
//b.N是基准测试框架提供的，表示循环的次数，因为需要反复调用测试的代码，才可以评估性能
func BenchmarkStack_Push(b *testing.B) {
	l := len(answer)
	s := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(l)
		s.Push(NewElement(answer[r]))
	}
}

//				   代表所有基准测试  显示每次操作分配内存次数和字节数	  默认执行测试时间为1s，可以用此参数改变
//运行命令 go test  -bench=".*"      -benchmem -v              -benchtime=3s
//结合pprof go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out
//goos: windows 系统
//goarch: amd64 位数
//pkg: algorithm/container/stack 包名
// benchmark 名字 - cpu核数			循环次数	 平均每次执行次数	 每次操作分配的字节数 每次执行内存分配次数
//BenchmarkStack_Push-8             118849   10455 ns/op       48 B/op			2 allocs/op
//2 allocs/op
//BenchmarkStack_Pop-8            1000000000  0.306 ns/op   	0 B/op			0 allocs/op
//0 allocs/op
//PASS
//   			执行目录					     总时间
//ok      algorithm/container/stack       7.956s

func BenchmarkStack_Pop(b *testing.B) {
	l := len(answer)
	s := New()
	for i := 0; i < 100000; i++ {
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(l)
		s.Push(NewElement(answer[r]))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}

func TestStack_Push(t *testing.T) {
	s := New()
	for _, e := range answer {
		s.Push(NewElement(e))
	}
	s.Pop()
	if (len(answer) - 1) != s.Len() {
		t.Errorf("expected stack lenght %d,actual %d", len(answer)-1, s.Len())
	}
	top := s.Peek()
	v := top.Value()
	// type assertion
	a, ok := v.(string)
	if !ok {
		t.Error("type assertion error")
	}

	if a != "css" {
		t.Error("expected html, actual", a)
	}

}

//逆波兰表达式
func PolandNotation() {

}
