package problems

// "Top-down" Solution
// preorder traversal
func maxDepth(root *TreeNode) int {
	var ans int
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		// a leaf node
		if node.Left == nil && node.Right == nil {
			// update answer
			if depth > ans {
				ans = depth
			}
		}

		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}
	dfs(root, 1)

	return ans
}

// "Bottom-up" Solution
// postorder traversal
func maxDepth1(root *TreeNode) int {
	return dfs(root)
}

func dfs(node *TreeNode) int {
	if node == nil {
		return 0
	}
	ld := dfs(node.Left)
	rd := dfs(node.Right)
	if ld > rd {
		return ld + 1
	}
	return rd + 1
}
