package rel

import "testing"

func Test_searchBST(t *testing.T) {
	target := TreeNode{
		Val: 10,
	}

	root := TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 1},
			Right: &TreeNode{
				Val:  4,
				Left: &TreeNode{Val: 3},
			},
		},
		Right: &TreeNode{
			Val:  7,
			Left: &TreeNode{Val: 6},
			Right: &TreeNode{
				Val:   9,
				Left:  &TreeNode{Val: 8},
				Right: &target,
			},
		},
	}

	got := searchBST(&root, 10)
	expect := &target
	if got != expect {
		t.Errorf("expect: %v, got %v", expect, got)
	}
}
