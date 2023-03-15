package traversal

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_preorder(t *testing.T) {
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
				root: &Node{
					Val: 1,
					Children: []*Node{
						{
							Val: 3,
							Children: []*Node{
								{Val: 5},
								{Val: 6},
							},
						},
						{Val: 2},
						{Val: 4},
					},
				},
			},
			want: []int{1, 3, 5, 6, 2, 4},
		},
		{
			name: "example 2",
			args: args{
				root: &Node{
					Val: 1,
					Children: []*Node{
						{Val: 2},
						{
							Val: 3,
							Children: []*Node{
								{Val: 6},
								{Val: 7, Children: []*Node{
									{Val: 11, Children: []*Node{
										{Val: 14},
									}},
								}},
							},
						},
						{Val: 4, Children: []*Node{
							{Val: 8, Children: []*Node{
								{Val: 12},
							}},
						}},
						{Val: 5, Children: []*Node{
							{Val: 9, Children: []*Node{
								{Val: 13},
							}},
							{Val: 10},
						}},
					},
				},
			},
			want: []int{1, 2, 3, 6, 7, 11, 14, 4, 8, 12, 5, 9, 13, 10},
		},
		{
			name: "empty",
		},
	}

	for idx, f := range []func(*Node) []int{
		preorder,
		preorder1,
	} {
		t.Run(fmt.Sprintf("func#%d", idx), func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := f(tt.args.root); !reflect.DeepEqual(got, tt.want) {
						t.Errorf("preorder() = %v, want %v", got, tt.want)
					}
				})
			}
		})
	}
}
