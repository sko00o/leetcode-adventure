package bfs

import (
	"github.com/sko00o/leetcode-adventure/queue-stack/queue"
)

// Node defines node with neighbors.
type Node struct {
	Val       int
	Neighbors []*Node
}

// NodeQueue is queue for Node.
type NodeQueue struct {
	queue.Queue
}

// EnQueue insert Node into queue.
func (q *NodeQueue) EnQueue(n *Node) bool {
	return q.Queue.EnQueue(n)
}

// Front get the front node from the queue.
func (q *NodeQueue) Front() *Node {
	val, ok := q.Queue.Front().(*Node)
	if !ok {
		return nil
	}
	return val
}
