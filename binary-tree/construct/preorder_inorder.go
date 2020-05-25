package construct

func buildTreePreorderInorder(preorder []int, inorder []int) *TreeNode {
	var inSet = make(map[int]int)
	for idx, n := range inorder {
		inSet[n] = idx
	}

	var pIdx int
	return constructPreorder(preorder, inSet, 0, len(inorder)-1, &pIdx)
}

func constructPreorder(preorder []int, inSet map[int]int, start, end int, pIdx *int) *TreeNode {
	if start > end {
		return nil
	}

	curr := preorder[*pIdx]
	*pIdx++
	f := constructPreorder
	return &TreeNode{
		Val:   curr,
		Left:  f(preorder, inSet, start, inSet[curr]-1, pIdx),
		Right: f(preorder, inSet, inSet[curr]+1, end, pIdx),
	}
}
