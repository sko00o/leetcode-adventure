package codec

import (
	"github.com/sko00o/leetcode-adventure/nary-tree/treenode"
)

type Node = treenode.Node

func (c *Codec) Serialize(root *Node) string {
	return c.serialize(root)
}

func (c *Codec) Deserialize(data string) *Node {
	return c.deserialize(data)
}
