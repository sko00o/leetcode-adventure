package bfs

import (
	"github.com/sko00o/leetcode-adventure/queue-stack/queue"
)

// Node defines node with neighbors.
type Node struct {
	Val       int
	Neighbors []*Node
}

// NodeQueue is queue for Nodes.
type NodeQueue struct {
	queue.SliceQueue
}

// EnQueue insert Node into queue.
func (q *NodeQueue) EnQueue(n *Node) bool {
	return q.SliceQueue.EnQueue(n)
}

// Front get the front node from the queue.
func (q *NodeQueue) Front() *Node {
	val, ok := q.SliceQueue.Front().(*Node)
	if !ok {
		return nil
	}
	return val
}
