package traversal

import (
	"fmt"
	"reflect"
	"testing"
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
			want: []int{5, 6, 3, 2, 4, 1},
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
