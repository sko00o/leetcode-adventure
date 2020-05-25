package traversal

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var queue []*TreeNode
	var out [][]int
	queue = append(queue, root)
	for len(queue) != 0 {
		var level []int
		for k := len(queue); k > 0; k-- {
			top := queue[0]
			level = append(level, top.Val)
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
			queue = queue[1:]
		}
		out = append(out, level)
	}
	return out
}
