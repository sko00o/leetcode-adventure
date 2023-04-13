package codec

import (
	"strconv"
	"strings"
)

// Codec defines a codec structure.
type Codec struct {
}

// Constructor return a Codec.
func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	var q []*TreeNode
	q = append(q, root)
	var out []string
	for len(q) != 0 {
		for k := len(q); k > 0; k-- {
			curr := q[0]
			if curr != nil {
				out = append(out, strconv.Itoa(curr.Val))
				q = append(q, curr.Left, curr.Right)
			} else {
				out = append(out, "null")
			}
			q = q[1:]
		}
	}
	return "[" + strings.TrimRight(strings.Join(out, ","), ",null") + "]"
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	if len(data) < 2 {
		return nil
	}
	if data[0] != '[' && data[len(data)-1] != ']' {
		return nil
	}
	data = data[1 : len(data)-1]

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

	if len(nums) == 0 {
		return nil
	}

	if nums[0] == nil {
		return nil
	}
	var root = &TreeNode{Val: *nums[0]}
	var q []*TreeNode
	q = append(q, root)

	for i := 1; len(q) != 0; {
		curr := q[0]

		if curr != nil {
			var left, right *int
			if i < len(nums) {
				left = nums[i]
				if i+1 < len(nums) {
					right = nums[i+1]
				}
				i += 2
			}
			if left != nil {
				curr.Left = &TreeNode{
					Val: *left,
				}
				q = append(q, curr.Left)
			}
			if right != nil {
				curr.Right = &TreeNode{
					Val: *right,
				}
				q = append(q, curr.Right)
			}
		}

		q = q[1:]
	}

	return root
}
