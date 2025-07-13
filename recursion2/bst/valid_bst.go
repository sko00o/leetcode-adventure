package bst

func isValidBST(root *TreeNode) bool {
	return isValid(root, nil, nil)
}
func isValid(node *TreeNode, min *int, max *int) bool {
	if node == nil {
		return true
	}
	if min != nil && node.Val <= *min {
		return false
	}
	if max != nil && node.Val >= *max {
		return false
	}
	return isValid(node.Left, min, &node.Val) && isValid(node.Right, &node.Val, max)
}

func isValidBST0(root *TreeNode) bool {
	return isValid0(root, nil, nil)
}
func isValid0(node, min, max *TreeNode) bool {
	if node == nil {
		return true
	}
	if min != nil && node.Val <= min.Val {
		return false
	}
	if max != nil && node.Val >= max.Val {
		return false
	}
	return isValid0(node.Left, min, node) && isValid0(node.Right, node, max)
}

func isValidBST1(root *TreeNode) bool {
	_, _, ok := getMinMax(root)
	return ok
}
func getMinMax(root *TreeNode) (min *int, max *int, ok bool) {
	if root == nil {
		return nil, nil, true
	}
	min, max = &root.Val, &root.Val

	if root.Left != nil {
		lmin, lmax, ok := getMinMax(root.Left)
		if !ok || (lmax != nil && *lmax >= root.Val) {
			return nil, nil, false
		}
		if lmin != nil {
			min = lmin
		}
	}
	if root.Right != nil {
		rmin, rmax, ok := getMinMax(root.Right)
		if !ok || (rmin != nil && *rmin <= root.Val) {
			return nil, nil, false
		}
		if rmax != nil {
			max = rmax
		}
	}
	return min, max, true
}
