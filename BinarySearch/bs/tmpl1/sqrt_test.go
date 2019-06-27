package tmpl1

import (
	"fmt"
	"testing"
)

func TestMySqrt(t *testing.T) {
	testMySqrt(t, mySqrt /*mySqrt1, mySqrt2*/)
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
