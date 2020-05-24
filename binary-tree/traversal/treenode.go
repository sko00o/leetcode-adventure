package traversal

import (
	"strconv"
	"strings"
)

// TreeNode is definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (n *TreeNode) String() string {
	return "[" + strings.TrimRight(strings.Join(bfs(n), ","), ",null") + "]"
}

func bfs(n *TreeNode) []string {
	var q []*TreeNode
	q = append(q, n)
	var out []string
	for len(q) != 0 {
		for k := len(q); k > 0; k-- {
			curr := q[0]
			if curr != nil {
				out = append(out, strconv.Itoa(curr.Val))
				if curr.Left != nil || curr.Right != nil {
					q = append(q, curr.Left, curr.Right)
				}
			} else {
				out = append(out, "null")
			}
			q = q[1:]
		}
	}

	return out
}
