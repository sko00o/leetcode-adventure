package tmpl2

import (
	"github.com/sko00o/leetcode-adventure/queue-stack/queue/bfs"
)

// BFS return the length of the shortest path between root and target node.
// Pro: we never visit a node twice.
func BFS(root, target *bfs.Node) int {
	// store all nodes which are waiting to be processed
	var queue bfs.NodeQueue

	// store all the nodes that we've visited
	visited := make(map[*bfs.Node]struct{})

	// number of steps neeeded from root to current node
	var step int

	// initialize
	queue.EnQueue(root)
	visited[root] = struct{}{}

	for !queue.IsEmpty() {
		// iterate the nodes which are already in the queue
		for size := len(queue.Data); size != 0; size-- {
			n := queue.Front()
			if n == target {
				return step
			}

			for i := range n.Neighbors {
				if _, ok := visited[n.Neighbors[i]]; !ok {
					queue.EnQueue(n.Neighbors[i])
					visited[n.Neighbors[i]] = struct{}{}
				}
			}
			queue.DeQueue()
		}

		step++
	}

	// there is no path from root to target
	return -1
}
