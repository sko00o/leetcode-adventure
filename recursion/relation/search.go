package relation

// TreeNode is definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root.Val == val {
		return root
	}

	if root.Val > val {
		if root.Left != nil {
			return searchBST(root.Left, val)
		}
		return nil
	}

	if root.Val < val {
		if root.Right != nil {
			return searchBST(root.Right, val)
		}
		return nil
	}

	return nil
}
