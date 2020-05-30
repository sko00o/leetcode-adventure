package problems

/*
Notes:
最近公共祖先。
迭代法，不借助指向祖先节点的指针 (lowestCommonAncestor2) ：
1. 从 root 节点开始，入栈。
2. 栈中保存节点以及其状态，状态区分 root 的一个孩子还是两个孩子需要遍历。
3. 查看栈顶元素。
4. 遍历任何节点的子节点前，先检查当前节点是不是 p 或者 q 。
5. 一但我们找到了 p 或者 q ，标记已找到一个，lca 索引指向该节点，
   此时栈里所有元素都是当前元素的祖先。
6. 第二次找到 p 或 q，那就找齐了，返回 lca 。
7. 无论何时访问子节点，都更新栈顶的祖先节点的状态。
8. 当前节点的左右子树都遍历完时，将其从栈中弹出，如果此时已找到过一个目标节点，
   此时栈顶节点可能是找到的祖先之一，将 lca 指向新的栈顶节点。
*/

// recursive
// time complexity: O(N)
// space complexity: O(N)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var lca *TreeNode

	var recur func(root, p, q *TreeNode) bool
	recur = func(curr, p, q *TreeNode) bool {
		if curr == nil {
			return false
		}

		var cnt int
		if recur(curr.Left, p, q) {
			cnt++
		}
		if recur(curr.Right, p, q) {
			cnt++
		}
		if curr == p || curr == q {
			cnt++
		}
		if cnt >= 2 {
			lca = curr
		}
		return cnt > 0
	}
	recur(root, p, q)
	return lca
}

// iterative
// time complexity: O(N)
// space complexity: O(N)
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var stk []*TreeNode
	var parent = make(map[*TreeNode]*TreeNode)

	stk = append(stk, root)

	var gotCnt int
	for gotCnt < 2 && len(stk) != 0 {
		node := stk[len(stk)-1]
		stk = stk[:len(stk)-1]

		if node.Left != nil {
			if node.Left == p || node.Left == q {
				gotCnt++
			}
			parent[node.Left] = node
			stk = append(stk, node.Left)
		}
		if node.Right != nil {
			if node.Right == p || node.Right == q {
				gotCnt++
			}
			parent[node.Right] = node
			stk = append(stk, node.Right)
		}
	}

	// backtracking process.
	var ancestors = make(map[*TreeNode]bool)
	for p != nil {
		ancestors[p] = true
		p = parent[p]
	}
	for !ancestors[q] {
		q = parent[q]
	}

	return q
}

// iterative without parent pointers
// time complexity: O(N)
// space complexity: O(N)
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	var stk []*pair
	if root != nil {
		stk = append(stk, &pair{N: root, S: bothPending})
	}

	var foundOne bool
	var lca *TreeNode

	for len(stk) != 0 {
		parent := stk[len(stk)-1]

		if parent.S == bothDone {
			// pop
			stk = stk[:len(stk)-1]
			// point to last ancestor
			if parent.N == lca && foundOne && len(stk) > 0 {
				lca = stk[len(stk)-1].N
			}
			continue
		}

		var childNode *TreeNode
		if parent.S == bothPending {
			if parent.N == p || parent.N == q {
				if foundOne {
					return lca
				}

				foundOne = true
				lca = parent.N
			}

			// if both pending, traverse the left child
			childNode = parent.N.Left
		} else {
			childNode = parent.N.Right
		}

		parent.S++ // -> leftDone -> bothDone

		if childNode != nil {
			stk = append(stk, &pair{N: childNode, S: bothPending})
		}
	}

	return nil
}

type status int

const (
	// 待处理
	bothPending status = iota
	// 左侧处理过了
	leftDone
	// 两子树都处理过了
	bothDone
)

type pair struct {
	N *TreeNode
	S status
}
