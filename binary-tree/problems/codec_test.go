package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCodec(t *testing.T) {

	tests := []struct {
		root *TreeNode
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
					Val: 4,
					Right: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val: 6,
						},
					},
				},
			},
		},
		{}, // empty test
	}

	for _, tt := range tests {
		assert := require.New(t)
		codec := Constructor()
		str := codec.serialize(tt.root)
		got := codec.deserialize(str)
		assert.Equal(tt.root, got)
	}
}
