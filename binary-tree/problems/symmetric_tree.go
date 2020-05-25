package problems

// recursively
// time complexity: O(N)
// space complexity: O(N)
func isSymmetric(root *TreeNode) bool {
	return cmp(root, root)
}

func cmp(n1, n2 *TreeNode) bool {
	if n1 == nil || n2 == nil {
		return n1 == n2
	}

	return n1.Val == n2.Val &&
		cmp(n1.Right, n2.Left) &&
		cmp(n1.Left, n2.Right)
}

// iteratively
// time complexity: O(N)
// space complexity: O(N)
func isSymmetric1(root *TreeNode) bool {
	var level = []*TreeNode{root}
	for len(level) != 0 {
		var newLevel []*TreeNode
		lLen := len(level)
		for i := 0; i <= lLen-1; i++ {
			if j := lLen - 1 - i; i < j {
				if level[i] == nil && level[j] == nil {
					continue
				}
				if level[i] == nil || level[j] == nil {
					return false
				}
				if level[i].Val != level[j].Val {
					return false
				}
			}

			if level[i] != nil {
				newLevel = append(newLevel, level[i].Left, level[i].Right)
			}
		}
		level = newLevel
	}
	return true
}

func isSymmetric2(root *TreeNode) bool {
	var q []*TreeNode
	q = append(q, root, root)
	for len(q) > 0 {
		var t1, t2 *TreeNode

		t1 = q[0]
		q = q[1:]
		if len(q) > 0 {
			t2 = q[0]
			q = q[1:]
		}

		if t1 == nil && t2 == nil {
			continue
		}

		if t1 == nil || t2 == nil {
			return false
		}

		if t1.Val != t2.Val {
			return false
		}

		q = append(q, t1.Left, t2.Right, t1.Right, t2.Left)
	}

	return true
}
