package radix

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestSort(t *testing.T) {
	s := []int{
		99, 100, 890, 3, 5, 1234, 7, 8, 9, 4, 1, 2, 6, 0, 10, 17, 11, 13, 12, 14, 15, 16, 18, 20, 19,
	}

	s = Sort(s)

	for _, e := range s {
		log.Println(e)
	}

	if s[len(s)-1] != 890 {
		t.Errorf("expected %d,actual %d", 890, s[len(s)-1])
	}

}

func BenchmarkSort(b *testing.B) {
	s := make([]int, 8000)
	l := len(s)
	for i := 0; i < l; i++ {
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(l)
		s[i] = r
	}

	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		Sort(s)
	}
}
