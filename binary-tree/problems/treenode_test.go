package problems

import (
	"reflect"
	"testing"
)

func Test_makeTree(t *testing.T) {
	type args struct {
		nums numSlice
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "a normal tree",
			args: args{
				nums: numSlice{{1}, {2}, {2}, nil, {3}, nil, {3}},
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
			if got := makeTree(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
