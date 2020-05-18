package practice

// TreeNode for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var out []int
	var stack []*TreeNode
	var visited = make(map[*TreeNode]bool)

	stack = append(stack, root)
	for len(stack) != 0 {
		curr := stack[len(stack)-1]
		if curr.Left != nil && !visited[curr.Left] {
			stack = append(stack, curr.Left)
			continue
		}
		out = append(out, curr.Val)
		visited[curr] = true
		stack = stack[:len(stack)-1]
		if curr.Right != nil && !visited[curr.Right] {
			stack = append(stack, curr.Right)
		}
	}
	return out
}

func inorderTraversal1(root *TreeNode) (out []int) {
	var stack []*TreeNode
	for curr := root; curr != nil || len(stack) != 0; {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		if len(stack) != 0 {
			curr = stack[len(stack)-1]
			out = append(out, curr.Val)
			stack = stack[:len(stack)-1]
			curr = curr.Right
		}
	}
	return
}

func inorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	f := inorderTraversal2
	return append(append(f(root.Left), root.Val), f(root.Right)...)
}
