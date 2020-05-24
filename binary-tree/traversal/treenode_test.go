package traversal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTreeNode(t *testing.T) {
	n := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
			},
		},
	}

	assert := require.New(t)
	assert.Equal("[1,null,3,2]", n.String())
}
