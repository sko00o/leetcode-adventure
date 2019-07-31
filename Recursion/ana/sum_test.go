package ana

import "testing"

var list = make([]int, 10000)

func Benchmark_sum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sum(list)
	}
}

func Benchmark_sum1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sum1(list)
	}
}
