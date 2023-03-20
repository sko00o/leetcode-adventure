package recursion

import (
	"fmt"
	"testing"

	"github.com/sko00o/leetcode-adventure/nary-tree/treenode"
)

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{root: treenode.ExampleTree1},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{root: treenode.ExampleTree2},
			want: 5,
		},
	}

	for idx, f := range []func(*Node) int{
		maxDepth,
		maxDepth1,
	} {
		t.Run(fmt.Sprintf("func#%d", idx), func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := f(tt.args.root); got != tt.want {
						t.Errorf("maxDepth() = %v, want %v", got, tt.want)
					}
				})
			}
		})
	}
}
