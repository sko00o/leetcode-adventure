package conclusion

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

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return genTrees(1, n)
}

func genTrees(start, end int) []*TreeNode {
	var out []*TreeNode
	if start > end {
		out = append(out, nil)
	}

	for i := start; i <= end; i++ {
		left := genTrees(start, i-1)
		right := genTrees(i+1, end)

		for _, l := range left {
			for _, r := range right {
				out = append(out, &TreeNode{
					Val:   i,
					Left:  l,
					Right: r,
				})
			}
		}
	}

	return out
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
