package problems

import (
	"strconv"
	"strings"

	"github.com/sko00o/leetcode-adventure/nary-tree/treenode"
)

type Node = treenode.Node

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *Node) string {
	var q []*Node
	q = append(q, root)
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
	return "[" + strings.Join(out, ",") + "]"
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *Node {
	// precheck
	if len(data) < 2 {
		return nil
	}
	if data[0] != '[' && data[len(data)-1] != ']' {
		return nil
	}
	data = data[1 : len(data)-1]

	// strings to numbers
	var nums []*int
	for _, s := range strings.Split(data, ",") {
		if s == "null" {
			nums = append(nums, nil)
		} else {
			num, err := strconv.Atoi(s)
			if err == nil {
				nums = append(nums, &num)
			}
		}
	}

	return makeTree(nums)
}

func makeTree(nums []*int) *Node {
	// special case
	if len(nums) == 0 {
		return nil
	}
	if nums[0] == nil {
		return nil
	}
	var root = &Node{Val: *nums[0]}
	var q []*Node
	q = append(q, root)

	for i := 1; len(q) != 0 && i < len(nums); {
		curr := q[0]
		isEnd := false
		j := 0
		for ; j < len(nums[i:]); j++ {
			if v := nums[i+j]; v == nil {
				if !isEnd {
					isEnd = true
				} else {
					break
				}
			} else {
				n := &Node{Val: *v}
				curr.Children = append(curr.Children, n)
				q = append(q, n)
			}
		}
		i += j
		q = q[1:]
	}
	return root
}
