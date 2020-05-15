package tmpl1

import "github.com/sko00o/leetcode-adventure/queue-stack/stack/dfs"

// DFS return true if there is a path from cur to target.
func DFS(curr, target *dfs.Node) bool {
	return dfsRecursion(curr, target, make(map[*dfs.Node]bool))
}

func dfsRecursion(curr, target *dfs.Node, visited map[*dfs.Node]bool) bool {
	if curr == target {
		return true
	}

	for i := range curr.Neighbors {
		if visited[curr.Neighbors[i]] {
			continue
		}
		visited[curr.Neighbors[i]] = true
		if dfsRecursion(curr.Neighbors[i], target, visited) {
			return true
		}
	}

	return false
}
