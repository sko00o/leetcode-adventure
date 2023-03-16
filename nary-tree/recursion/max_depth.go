package recursion

// "Top-down" Solusion
// preorder traversal
func maxDepth(root *Node) int {
	var dfs func(*Node, int) int
	dfs = func(node *Node, depth int) int {
		if node == nil {
			return 0
		}
		currDepth := depth + 1
		for _, child := range node.Children {
			if v := dfs(child, currDepth); v > depth {
				depth = v
			}
		}
		return depth
	}
	return dfs(root, 1)
}

// "Bottom-up" Solution
func maxDepth1(root *Node) int {
	var dfs func(*Node) int
	dfs = func(node *Node) int {
		if node == nil {
			return 0
		}
		max := 1
		for _, child := range node.Children {
			if currMax := dfs(child); currMax+1 > max {
				max = currMax + 1
			}
		}
		return max
	}
	return dfs(root)
}
