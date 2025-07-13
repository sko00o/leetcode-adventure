package bst

import (
	"fmt"
	"math"
	"testing"
)

func Test_isValidBST(t *testing.T) {
	max := math.MaxInt
	min := math.MinInt

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				root: cc.Deserialize("[2,1,3]"),
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				root: cc.Deserialize("[5,1,4,null,null,3,6]"),
			},
			want: false,
		},
		{
			name: "Test 1",
			args: args{
				root: cc.Deserialize("[5,4,6,null,null,3,7]"),
			},
			want: false,
		},
		{
			name: "Test 2",
			args: args{
				root: cc.Deserialize("[2,2,2]"),
			},
			want: false,
		},
		{
			name: "Test Max",
			args: args{
				root: cc.Deserialize(fmt.Sprintf("[%d,%d,%d]", max-1, max-2, max)),
			},
			want: true,
		},
		{
			name: "Test Min",
			args: args{
				root: cc.Deserialize(fmt.Sprintf("[%d,%d,%d]", min+1, min, min+2)),
			},
			want: true,
		},
		{
			name: "Empty",
			want: true,
		},
	}

	for _, f := range []func(root *TreeNode) bool{
		isValidBST,
		isValidBST0,
		isValidBST1,
	} {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.root); got != tt.want {
					t.Errorf("isValidBST() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}
