package traversal

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/sko00o/leetcode-adventure/nary-tree/treenode"
)

func Test_postorder(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{
				root: treenode.ExampleTree1,
			},
			want: []int{5, 6, 3, 2, 4, 1},
		},
		{
			name: "example 2",
			args: args{
				root: treenode.ExampleTree2,
			},
			want: []int{2, 6, 14, 11, 7, 3, 12, 8, 4, 13, 9, 10, 5, 1},
		},
		{
			name: "empty",
		},
	}

	for idx, f := range []func(*Node) []int{
		postorder,
		postorder1,
	} {
		t.Run(fmt.Sprintf("func#%d", idx), func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := f(tt.args.root); !reflect.DeepEqual(got, tt.want) {
						t.Errorf("postorder() = %v, want %v", got, tt.want)
					}
				})
			}
		})
	}
}
