package problems

import (
	"fmt"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	tests := []struct {
		input  numSlice
		output int
	}{
		{
			input: numSlice{
				{3},
				{9},
				{20},
				nil,
				nil,
				{15},
				{7},
			},
			output: 3,
		},
		{
			input: numSlice{
				{3},
			},
			output: 1,
		},
	}

	for idx, f := range []func(*TreeNode) int{
		maxDepth,
		maxDepth1,
	} {
		t.Run(fmt.Sprintf("func#%d", idx), func(t *testing.T) {
			for _, tst := range tests {
				got := f(makeTree(tst.input))
				if got != tst.output {
					t.Errorf("expect: %d, got: %d", tst.output, got)
				}
			}
		})
	}
}
