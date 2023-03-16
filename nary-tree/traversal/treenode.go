package traversal

import (
	"strconv"
	"strings"
)

// Node is definition for a n-ary tree node.
type Node struct {
	Val      int
	Children []*Node
}

func (n *Node) String() string {
	return "[" + strings.Join(bfs(n), ",") + "]"
}

func bfs(n *Node) []string {
	var q []*Node
	q = append(q, n)
	var out []string
	var lq = []int{len(q)}
	for len(q) != 0 {
		if len(out) > 0 {
			out = append(out, "null")
		}
		for k := lq[0]; k > 0; k-- {
			curr := q[0]
			if curr != nil {
				out = append(out, strconv.Itoa(curr.Val))
				q = append(q, curr.Children...)
				lq = append(lq, len(curr.Children))
			}
			q = q[1:]
		}
		lq = lq[1:]
	}

	return out
}
