package problems

/*
Follow up:
1. only use constant extra space.
2. recursive solution is fine.
*/

// normal binary tree
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	p := root.Next

	// next subtree's most left child
	for p != nil {
		if p.Left != nil {
			p = p.Left
			break
		}

		if p.Right != nil {
			p = p.Right
			break
		}

		p = p.Next
	}

	// connect root's childs
	if root.Right != nil {
		root.Right.Next = p
	}

	if root.Left != nil {
		if root.Right != nil {
			root.Left.Next = root.Right
		} else {
			root.Left.Next = p
		}
	}

	f := connect
	f(root.Right)
	f(root.Left)
	return root
}

// use a queue, is not good.
func connect1(root *Node) *Node {
	if root == nil {
		return nil
	}

	var q []*Node
	q = append(q, root)

	for len(q) != 0 {
		var nq []*Node
		for i := 0; i+1 < len(q); i++ {
			q[i].Next = q[i+1]

			if n := q[i].Left; n != nil {
				nq = append(nq, n)
			}
			if n := q[i].Right; n != nil {
				nq = append(nq, n)
			}
		}

		if n := q[len(q)-1].Left; n != nil {
			nq = append(nq, n)
		}
		if n := q[len(q)-1].Right; n != nil {
			nq = append(nq, n)
		}
		q = nq
	}

	return root
}
