package traversal

// recursively
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	f := preorderTraversal
	return append(append([]int{root.Val}, f(root.Left)...), f(root.Right)...)
}

// iteratively
func preorderTraversal1(root *TreeNode) (out []int) {
	var stack []*TreeNode
	for curr := root; curr != nil || len(stack) != 0; {
		if curr != nil {
			out = append(out, curr.Val)
			stack = append(stack, curr) // push
			curr = curr.Left
		} else {
			curr = stack[len(stack)-1]   // top
			stack = stack[:len(stack)-1] // pop
			curr = curr.Right
		}
	}

	return out
}

// Morris traversal
// time complexity: O(N)
// space complexity: O(1)
func preorderTraversal2(root *TreeNode) (out []int) {
	var prev *TreeNode
	for curr := root; curr != nil; {
		if curr.Left == nil {
			out = append(out, curr.Val)
			curr = curr.Right
		} else {
			// Find in-order predecessor for curr node.
			prev = curr.Left
			for prev.Right != nil && prev.Right != curr {
				prev = prev.Right
			}

			// Create links to the in-order successor.
			if prev.Right == nil {
				out = append(out, curr.Val) // *
				prev.Right = curr
				curr = curr.Left
			} else {
				// Revert the changes to restore original tree.
				prev.Right = nil
				curr = curr.Right
			}
		}
	}
	return
}
