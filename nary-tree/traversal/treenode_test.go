package traversal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	exampleTree1 = &Node{
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
	}

	exampleTree2 = &Node{
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
	}
)

func TestNode(t *testing.T) {
	tests := []struct {
		root *Node
		want string
	}{
		{
			root: exampleTree1,
			want: "[1,null,3,2,4,null,5,6]",
		},
		{
			root: exampleTree2,
			want: "[1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case#%d", i), func(t *testing.T) {
			assert := require.New(t)
			assert.Equal(tt.want, tt.root.String())
		})
	}
}
