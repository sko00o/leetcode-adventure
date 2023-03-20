package codec

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sko00o/leetcode-adventure/nary-tree/treenode"
)

func TestCodec(t *testing.T) {
	tests := []struct {
		root *Node
	}{
		{
			root: treenode.ExampleTree1,
		},
		{
			root: treenode.ExampleTree2,
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
