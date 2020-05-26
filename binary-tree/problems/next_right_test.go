package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_connect(t *testing.T) {
	tests := []struct {
		root *Node
	}{
		{
			root: &Node{
				Val: 1,
				Left: &Node{
					Val: 2,
					Left: &Node{
						Val: 4,
					},
					Right: &Node{
						Val: 5,
					},
				},
				Right: &Node{
					Val: 3,
					Left: &Node{
						Val: 6,
					},
					Right: &Node{
						Val: 7,
					},
				},
			},
		},
		{},
	}

	for _, tt := range tests {
		assert := require.New(t)
		root := connect(tt.root)
		assert.True(check(root))
	}
}

func check(root *Node) bool {
	if root == nil {
		return true
	}
	var q []*Node
	q = append(q, root)

	for len(q) != 0 {
		var nq []*Node
		for i := 0; i+1 < len(q); i++ {
			if q[i].Next != q[i+1] {
				return false
			}

			if n := q[i].Left; n != nil {
				nq = append(nq, n)
			}
			if n := q[i].Right; n != nil {
				nq = append(nq, n)
			}
		}
		if q[len(q)-1].Next != nil {
			return false
		}

		if n := q[len(q)-1].Left; n != nil {
			nq = append(nq, n)
		}
		if n := q[len(q)-1].Right; n != nil {
			nq = append(nq, n)
		}
		q = nq
	}
	return true
}
