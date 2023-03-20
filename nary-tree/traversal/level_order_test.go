package traversal

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/sko00o/leetcode-adventure/nary-tree/treenode"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "example 1",
			args: args{
				root: treenode.ExampleTree1,
			},
			want: [][]int{{1}, {3, 2, 4}, {5, 6}},
		},
		{
			name: "example 2",
			args: args{
				root: treenode.ExampleTree2,
			},
			want: [][]int{{1}, {2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13}, {14}},
		},
		{
			name: "empty",
		},
	}

	for idx, f := range []func(*Node) [][]int{
		levelOrder,
	} {
		t.Run(fmt.Sprintf("func#%d", idx), func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := f(tt.args.root); !reflect.DeepEqual(got, tt.want) {
						t.Errorf("levelOrder() = %v, want %v", got, tt.want)
					}
				})
			}
		})
	}
}
