package problems

import (
	"fmt"
	"testing"

	"github.com/sko00o/leetcode-adventure/binary-tree/treenode"
)

func Test_maxDepth(t *testing.T) {
	tests := []struct {
		input  treenode.NumSlice
		output int
	}{
		{
			input: treenode.NumSlice{
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
			input: treenode.NumSlice{
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
				got := f(treenode.New(tst.input))
				if got != tst.output {
					t.Errorf("expect: %d, got: %d", tst.output, got)
				}
			}
		})
	}
}
