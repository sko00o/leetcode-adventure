package treenode

import (
	"strconv"
	"strings"
)

// TreeNode is definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (n *TreeNode) String() string {
	return "[" + strings.Join(bfs(n), ",") + "]"
}

func bfs(n *TreeNode) []string {
	var q []*TreeNode
	q = append(q, n)
	var out []string
	for len(q) != 0 {
		for k := len(q); k > 0; k-- {
			curr := q[0]
			if curr != nil {
				out = append(out, strconv.Itoa(curr.Val))
				if curr.Left != nil || curr.Right != nil {
					q = append(q, curr.Left, curr.Right)
				}
			} else {
				out = append(out, "null")
			}
			q = q[1:]
		}
	}

	i := len(out)
	for out[i-1] == "null" {
		i--
	}
	return out[:i]
}

type NumSlice []*struct {
	N int
}

type tagNode struct {
	p      *TreeNode
	isLeft bool
}

func New(nums NumSlice) *TreeNode {
	queue := make(chan tagNode, 1+len(nums))
	var root *TreeNode
	for i, v := range nums {
		var node tagNode
		if i != 0 {
			node = <-queue
		} else {
			root = &TreeNode{Val: v.N}
			queue <- tagNode{p: root, isLeft: true}
			queue <- tagNode{p: root, isLeft: false}
			continue
		}

		if v != nil {
			var p *TreeNode
			if node.isLeft {
				p = &TreeNode{Val: v.N}
				node.p.Left = p
			} else {
				p = &TreeNode{Val: v.N}
				node.p.Right = p
			}
			queue <- tagNode{p: p, isLeft: true}
			queue <- tagNode{p: p, isLeft: false}
		}
	}
	return root
}
