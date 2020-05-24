package traversal

// recursively
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	f := postorderTraversal
	return append(append(f(root.Left), f(root.Right)...), root.Val)
}

// iteratively
func postorderTraversal1(root *TreeNode) (out []int) {
	var stack []*TreeNode
	var last *TreeNode
	for curr := root; curr != nil || len(stack) != 0; {
		if curr != nil {
			stack = append(stack, curr) // push
			curr = curr.Left
		} else {
			top := stack[len(stack)-1] // top
			if top.Right == nil || top.Right == last {
				out = append(out, top.Val)
				stack = stack[:len(stack)-1] // pop
				last = top
			} else {
				curr = top.Right
			}
		}
	}

	return out
}

// Morris traversal
// time complexity: O(N)
// space complexity: O(1)
func postorderTraversal2(root *TreeNode) (out []int) {
	var prev *TreeNode
	// use for access right subtree.
	var curr = &TreeNode{Left: root}
	for curr != nil {
		if curr.Left == nil {
			curr = curr.Right
		} else {
			// Find in-order predecessor for curr node.
			prev = curr.Left
			for prev.Right != nil && prev.Right != curr {
				prev = prev.Right
			}

			// Create links to the in-order successor.
			if prev.Right == nil {
				prev.Right = curr
				curr = curr.Left
			} else {
				// Revert the changes to restore original tree.
				out = append(out, visitReverse(curr.Left, prev)...)
				prev.Right = nil
				curr = curr.Right
			}
		}
	}
	return
}

func visitReverse(n1, n2 *TreeNode) (out []int) {
	reverse(n1, n2)
	defer reverse(n2, n1)
	for curr := n2; curr != n1; curr = curr.Right {
		out = append(out, curr.Val)
	}
	out = append(out, n1.Val)
	return
}

func reverse(start, end *TreeNode) {
	var prev = start
	var curr = prev.Right
	var next = curr.Right
	for prev != end {
		curr.Right = prev
		prev = curr
		curr = next
		next = next.Right
	}
}
