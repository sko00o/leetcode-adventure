package problems

// TreeNode is definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type numSlice []*struct {
	n int
}

type tagNode struct {
	p      *TreeNode
	isLeft bool
}

func makeTree(nums numSlice) *TreeNode {
	queue := make(chan tagNode, 1+len(nums))
	var root *TreeNode
	for i, v := range nums {
		var node tagNode
		if i != 0 {
			node = <-queue
		} else {
			root = &TreeNode{Val: v.n}
			queue <- tagNode{p: root, isLeft: true}
			queue <- tagNode{p: root, isLeft: false}
			continue
		}

		if v != nil {
			var p *TreeNode
			if node.isLeft {
				p = &TreeNode{Val: v.n}
				node.p.Left = p
			} else {
				p = &TreeNode{Val: v.n}
				node.p.Right = p
			}
			queue <- tagNode{p: p, isLeft: true}
			queue <- tagNode{p: p, isLeft: false}
		}
	}
	return root
}
