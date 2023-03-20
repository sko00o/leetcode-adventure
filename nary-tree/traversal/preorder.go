package traversal

// recursively
// time complexity: O(N), N is the number of nodes in the given tree
// space complexity: O(H), H is the height of the given tree
func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var res = []int{root.Val}
	for _, child := range root.Children {
		res = append(res, preorder(child)...)
	}
	return res
}

// iteratively using stack
// time complexity: O(N), N is the number of nodes in the given tree
// space complexity: O(H), H is the height of the given tree
func preorder1(root *Node) (out []int) {
	var stack []*Node
	for curr := root; curr != nil || len(stack) != 0; {
		if curr != nil {
			out = append(out, curr.Val)
			for i := len(curr.Children) - 1; i >= 0; i-- {
				stack = append(stack, curr.Children[i]) // push
			}
			curr = nil
			continue
		}
		curr = stack[len(stack)-1]   // top
		stack = stack[:len(stack)-1] // pop
	}
	return out
}
