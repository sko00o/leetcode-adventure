package classic

import (
	"testing"
)

func TestEqual(t *testing.T) {
	tcs := []struct {
		a, b   []int
		expect bool
	}{
		{
			a:      []int{1, 2, 3},
			b:      []int{1, 2, 3},
			expect: true,
		},
		{
			a:      []int{1, 2, 3},
			b:      []int{1, 3, 3},
			expect: false,
		},
		{
			a:      []int{1, 2, 3},
			b:      []int{1, 2},
			expect: false,
		},
		{
			a:      []int{},
			b:      []int{},
			expect: true,
		},
	}

	for _, tc := range tcs {
		if actual := Equal(tc.a, tc.b); actual != tc.expect {
			t.Errorf("actual: %v, expect: %v", actual, tc.expect)
		}
	}
}
