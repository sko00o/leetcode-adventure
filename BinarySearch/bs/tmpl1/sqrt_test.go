package tmpl1

import (
	"fmt"
	"testing"
)

func TestMySqrt(t *testing.T) {
	testMySqrt(t, mySqrt, mySqrt1)
}

func testMySqrt(t *testing.T, fs ...func(int) int) {
	tcs := []struct {
		input  int
		expect int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 1},
		{4, 2},
		{5, 2},
		{6, 2},
		{7, 2},
		{8, 2},
		{9, 3},
		{65535, 255},
		{65536, 256},
		{65537, 256},
		{1<<31 - 1, 46340},
		{1<<63 - 1, 3037000499},
	}

	for fIdx, f := range fs {
		for i, tc := range tcs {
			t.Run(fmt.Sprintf("func #%d, task #%d", fIdx, i), func(t *testing.T) {
				if actual := f(tc.input); actual != tc.expect {
					t.Errorf("actual: %d, expect: %d", actual, tc.expect)
				}
			})
		}
	}
}

func BenchmarkMySqrt(b *testing.B) {
	benchmarkMySqrt(b, mySqrt)
}

func BenchmarkMySqrt1(b *testing.B) {
	benchmarkMySqrt(b, mySqrt1)
}

func benchmarkMySqrt(b *testing.B, f func(int) int) {
	for i := 0; i < b.N; i++ {
		for n := 1<<63 - 1<<10; n < (1<<63 - 1); n += 1 {
			f(n)
		}
		for n := 1; n < 1<<10; n += 1 {
			f(n)
		}
	}
}

/*

goos: linux
goarch: amd64
pkg: adventure/BinarySearch/bs/tmpl1
BenchmarkMySqrt-8           5000            289058 ns/op
BenchmarkMySqrt1-8          3000            552816 ns/op
PASS

*/
