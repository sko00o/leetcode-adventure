package problems

import (
	"fmt"
	"testing"
)

func Test_numTrees(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 2},
		{3, 5},
		{4, 14},
	}

	for idx, f := range []func(int) int{
		numTrees,
		numTrees1,
	} {
		t.Run(fmt.Sprintf("func#%d", idx), func(t *testing.T) {
			for _, tt := range tests {
				if got := f(tt.n); got != tt.want {
					t.Errorf("numTrees() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
