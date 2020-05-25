package problems

func hasPathSum(node *TreeNode, sum int) bool {
	if node == nil {
		return false
	}
	if node.Left == nil && node.Right == nil {
		return sum == node.Val
	}

	var dfs = hasPathSum
	return dfs(node.Left, sum-node.Val) || dfs(node.Right, sum-node.Val)
}
