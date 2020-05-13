package tmpl1

import (
	"github.com/sko00o/leetcode-adventure/queue-stack/queue/bfs"
)

// BFS return the length of the shortest path between root and target node.
func BFS(root, target *bfs.Node) int {
	// store all nodes which are waiting to be processed
	var queue bfs.NodeQueue

	// number of steps needed from root to current node
	var step int

	// initialize
	queue.EnQueue(root)

	for !queue.IsEmpty() {
		// iterate the nodes which are already in the queue
		for size := len(queue.Data); size != 0; size-- {
			n := queue.Front()
			if n == target {
				return step
			}

			for i := range n.Neighbors {
				queue.EnQueue(n.Neighbors[i])
			}
			queue.DeQueue()
		}

		step++
	}

	// there is no path from root to target
	return -1
}
