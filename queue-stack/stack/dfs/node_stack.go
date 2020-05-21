package dfs

import "github.com/sko00o/leetcode-adventure/queue-stack/stack"

// Node defines node with neighbors.
type Node struct {
	Val       int
	Neighbors []*Node
}

// NodeStack is stack for Nodes.
type NodeStack struct {
	stack.SliceStack
}

// Push insert a Node into the stack.
func (s *NodeStack) Push(n *Node) {
	s.SliceStack.Push(n)
}

// Top get the top Node from the stack.
func (s *NodeStack) Top() *Node {
	if s.IsEmpty() {
		return nil
	}
	n, ok := s.Data[len(s.Data)-1].(*Node)
	if !ok {
		return nil
	}
	return n
}
