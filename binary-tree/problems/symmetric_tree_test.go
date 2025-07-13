package problems

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/sko00o/leetcode-adventure/binary-tree/treenode"
)

func Test_isSymmetric(t *testing.T) {
	tests := []struct {
		nums treenode.NumSlice
		want bool
	}{
		{
			nums: treenode.NumSlice{{1}, {2}, {2}, {3}, {4}, {4}, {3}},
			want: true,
		},
		{
			nums: treenode.NumSlice{{1}, {2}, {2}, nil, {3}, nil, {3}},
			want: false,
		},
		{
			nums: treenode.NumSlice{},
			want: true,
		},
	}

	for idx, f := range []func(*TreeNode) bool{
		isSymmetric,
		isSymmetric1,
		isSymmetric2,
	} {
		t.Run(fmt.Sprintf("func#%d", idx), func(t *testing.T) {
			for _, tt := range tests {
				if got := f(treenode.New(tt.nums)); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
