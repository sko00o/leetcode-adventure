package treenode

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTreeNode_String(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want string
	}{
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: "[1,2,2,null,3,null,3]",
		},
		{
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 2,
					},
				},
			},
			want: "[1,null,3,2]",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case#%d", i), func(t *testing.T) {
			assert := require.New(t)
			assert.Equal(tt.want, tt.root.String())
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		nums NumSlice
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "a normal tree",
			args: args{
				nums: NumSlice{{1}, {2}, {2}, nil, {3}, nil, {3}},
			},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
		},
		{
			name: "empty",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
