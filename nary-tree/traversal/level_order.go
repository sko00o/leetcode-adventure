package traversal

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	var queue []*Node
	queue = append(queue, root)
	var out [][]int
	for len(queue) != 0 {
		var tmp []int
		for k := len(queue); k > 0; k-- {
			curr := queue[0]
			if curr != nil {
				tmp = append(tmp, curr.Val)
				queue = append(queue, curr.Children...)
			}
			queue = queue[1:]
		}
		out = append(out, tmp)
	}
	return out
}
