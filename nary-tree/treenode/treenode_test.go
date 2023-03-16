package treenode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNode(t *testing.T) {
	tests := []struct {
		root *Node
		want string
	}{
		{
			root: ExampleTree1,
			want: "[1,null,3,2,4,null,5,6]",
		},
		{
			root: ExampleTree2,
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
