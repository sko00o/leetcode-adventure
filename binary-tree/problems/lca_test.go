package problems

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	n := []*TreeNode{
		{Val: 0},
		{Val: 1},
		{Val: 2},
		{Val: 3},
		{Val: 4},
		{Val: 5},
		{Val: 6},
	}
	n[0].Left = n[1]
	n[0].Right = n[2]
	n[1].Left = n[3]
	n[1].Right = n[4]
	n[2].Left = n[5]
	n[3].Right = n[6]

	/*
			   0
			 /   \
			1     2
		   / \   /
		  3   4 5
		   \
		    6
	*/

	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "test 1",
			args: args{
				root: n[0],
				p:    n[6],
				q:    n[4],
			},
			want: n[1],
		},
		{
			name: "test 2",
			args: args{
				root: n[0],
				p:    n[6],
				q:    n[1],
			},
			want: n[1],
		},
		{
			name: "test 3",
			args: args{
				root: n[0],
				p:    n[5],
				q:    n[3],
			},
			want: n[0],
		},
		{
			name: "empty",
		},
	}

	for idx, f := range []func(root, p, q *TreeNode) *TreeNode{
		lowestCommonAncestor,
		lowestCommonAncestor1,
		lowestCommonAncestor2,
	} {
		t.Run(fmt.Sprintf("func#%d", idx), func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := f(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
						t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
					}
				})
			}
		})
	}
}
