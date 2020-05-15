package practice

import (
	"fmt"
	"testing"
)

func Test_cloneGraph(t *testing.T) {
	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "example 1",
			args: args{
				node: newGraph([][]int{
					{2, 4}, {1, 3}, {2, 4}, {1, 3},
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cloneGraph(tt.args.node); !isCopyed(got, tt.args.node) {
				t.Error("not copyed correct")
			}
		})
	}
}

func Test_isCopyed(t *testing.T) {
	tests := []struct {
		al1, al2 [][]int
		want     bool
	}{
		{
			al1: [][]int{
				{2, 3},
				{1, 3},
				{1, 2},
			},
			al2: [][]int{
				{2, 3},
				{1, 3},
				{1, 2},
			},
			want: true,
		},
		{
			al1: [][]int{
				{3, 4},
				{4},
				{1},
				{1, 2},
			},
			al2: [][]int{
				{3, 4},
				{3},
				{1, 2},
				{1},
			},
			want: false,
		},
	}

	for idx, tt := range tests {
		t.Run(fmt.Sprintf("test %d", idx), func(t *testing.T) {
			if got := isCopyed(newGraph(tt.al1), newGraph(tt.al2)); got != tt.want {
				t.Errorf("expected: %v, got: %v", tt.want, got)
			}
		})
	}
}
