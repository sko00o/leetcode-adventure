package traversal

// recursively
func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var res []int
	for _, child := range root.Children {
		res = append(res, postorder(child)...)
	}
	return append(res, root.Val)
}

// literatively
func postorder1(root *Node) (out []int) {
	var stack []*Node
	for curr := root; curr != nil || len(stack) != 0; {
		if curr != nil {
			out = append([]int{curr.Val}, out...)   // prepend
			stack = append(stack, curr.Children...) // push
			curr = nil
			continue
		}
		curr = stack[len(stack)-1]   // top
		stack = stack[:len(stack)-1] // pop
	}
	return out
}
