package construct

func buildTreeInorderPostorder(inorder []int, postorder []int) *TreeNode {
	var inSet = make(map[int]int)
	for idx, n := range inorder {
		inSet[n] = idx
	}

	var pIdx = len(inorder)-1
	return constructPostorder(postorder, inSet, 0, len(inorder)-1, &pIdx)
}

func constructPostorder(postorder []int, inSet map[int]int, start, end int, pIdx *int)  *TreeNode {
	if start > end {
		return nil
	}

	curr := postorder[*pIdx]
	*pIdx--
	f := constructPostorder
	return &TreeNode {
		Val: curr,
		Right: f(postorder, inSet, inSet[curr]+1, end, pIdx),
		Left:  f(postorder, inSet, start, inSet[curr]-1, pIdx),
	}
}
