package ana

// TreeNode is definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root != nil {
		return dfs(1, root)
	}
	return 0
}

func dfs(d int, r *TreeNode) int {
	if r.Left == nil && r.Right == nil {
		return d
	}

	var ld, rd int
	if r.Left != nil {
		ld = dfs(d+1, r.Left)
	}

	if r.Right != nil {
		rd = dfs(d+1, r.Right)
	}

	if ld > rd {
		return ld
	}
	return rd
}
