package problems

// Node is definition for a Node.
type Node struct {
	Val               int
	Left, Right, Next *Node
}

func connect(root *Node) *Node {
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
