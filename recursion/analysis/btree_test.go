package ana

import "testing"

func Test_maxDepth(t *testing.T) {
	tests := []struct {
		input  []*num
		output int
	}{
		{
			input: []*num{
				{3},
				{9},
				{20},
				nil,
				nil,
				{15},
				{7},
			},
			output: 3,
		},
		{
			input: []*num{
				{3},
			},
			output: 1,
		},
	}
	for _, tst := range tests {
		got := maxDepth(makeTree(tst.input))
		if got != tst.output {
			t.Errorf("expect: %d, got: %d", tst.output, got)
		}
	}
}

type num struct {
	n int
}

type makeQueue struct {
	p      *TreeNode
	isLeft bool
}

func makeTree(nums []*num) *TreeNode {
	queue := make(chan makeQueue, 1+len(nums))
	var root *TreeNode
	for i, v := range nums {
		var node makeQueue
		if i != 0 {
			node = <-queue
		} else {
			root = &TreeNode{Val: v.n}
			queue <- makeQueue{p: root, isLeft: true}
			queue <- makeQueue{p: root, isLeft: false}
			continue
		}

		if v != nil {
			var p *TreeNode
			if node.isLeft {
				p = &TreeNode{Val: v.n}
				node.p.Left = p
			} else {
				p = &TreeNode{Val: v.n}
				node.p.Right = p
			}
			queue <- makeQueue{p: p, isLeft: true}
			queue <- makeQueue{p: p, isLeft: false}
		}
	}
	return root
}
