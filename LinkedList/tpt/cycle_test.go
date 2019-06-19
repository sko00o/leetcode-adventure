package tpt

import (
	"testing"
)

func makeLinkedListCycle(pos int, head ...int) (h *ListNode) {
	var c, p *ListNode
	for idx, v := range head {
		n := ListNode{
			Val: v,
		}

		if p != nil {
			p.Next = &n
			p = p.Next
		} else {
			// first node
			p = &n
			h = &n
		}

		// cycle start
		if idx == pos {
			c = &n
		}
		if idx == len(head)-1 {
			n.Next = c
		}
	}
	return
}

func getIndex(head *ListNode, pos *ListNode) int {
	for i, p := 0, head; p != nil; i, p = i+1, p.Next {
		if pos == p {
			return i
		}
	}

	return -1
}

func Test_hasCycle(t *testing.T) {
	tasks := []struct {
		head   []int
		pos    int
		expect bool
	}{
		{head: []int{3, 2, 0, -4}, pos: 1, expect: true},
		{head: []int{1, 2}, pos: 0, expect: true},
		{head: []int{1}, pos: -1, expect: false},
		{head: []int{1, 2, 3}, pos: 3, expect: false},
	}

	for i, task := range tasks {
		h := makeLinkedListCycle(task.pos, task.head...)
		if res := hasCycle(h); res != task.expect {
			t.Errorf("task #%d failed, output: %v, expect: %v", i, res, task.expect)
		}
	}
}

func Test_detectCycle(t *testing.T) {
	tasks := []struct {
		head   []int
		pos    int
		expect int
	}{
		{head: []int{3, 2, 0, -4}, pos: 1, expect: 1},
		{head: []int{1, 2}, pos: 0, expect: 0},
		{head: []int{1}, pos: -1, expect: -1},
		{head: []int{1, 2, 3}, pos: 3, expect: -1},
	}

	for i, task := range tasks {
		h := makeLinkedListCycle(task.pos, task.head...)
		if res := getIndex(h, detectCycle(h)); res != task.expect {
			t.Errorf("task #%d failed, output: %v, expect: %v", i, res, task.expect)
		}
	}
}
