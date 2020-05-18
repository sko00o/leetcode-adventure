package tmpl2

import "github.com/sko00o/leetcode-adventure/queue-stack/stack/dfs"

// 避免函数递归过深造成栈溢出，显式使用栈。

// DFS return true if there is a path from cur to target.
func DFS(root, target *dfs.Node) bool {
	visited := make(map[*dfs.Node]bool)
	var stack dfs.NodeStack

	// add root to stack
	visited[root] = true
	stack.Push(root)

	for !stack.IsEmpty() {
		curr := stack.Top()
		if curr == target {
			return true
		}
		stack.Pop()
		for i := range curr.Neighbors {
			next := curr.Neighbors[i]
			if !visited[next] {
				visited[next] = true
				stack.Push(next)
			}
		}
	}

	return false
}
