package codec

import (
	"github.com/sko00o/leetcode-adventure/binary-tree/treenode"
)

type TreeNode = treenode.TreeNode

func (c *Codec) Serialize(root *TreeNode) string {
	return c.serialize(root)
}

func (c *Codec) Deserialize(data string) *TreeNode {
	return c.deserialize(data)
}
