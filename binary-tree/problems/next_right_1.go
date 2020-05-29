package problems

/*
Follow up:
1. only use constant extra space.
2. recursive solution is fine.
*/

// Node is definition for a Node.
type Node struct {
	Val               int
	Left, Right, Next *Node
}

// for perfect binary tree only
func connect0(root *Node) *Node {
	if root != nil {
		conn(root.Left, root.Right)
	}
	return root
}

func conn(n1, n2 *Node) {
	if n1 == nil || n2 == nil {
		return
	}
	n1.Next = n2

	conn(n1.Right, n2.Left)
	conn(n1.Left, n1.Right)
	conn(n2.Left, n2.Right)
}

func connect01(root *Node) *Node {
	node := root 
	rowRoot := node 
	for node != nil && node.Left != nil {
		node.Left.Next = node.Right
		if node.Next != nil {
			node.Right.Next = node.Next.Left
			node= node.Next
		} else {
			// go next level
			node = rowRoot.Left
			rowRoot = node
		}
	}
	return root
}

