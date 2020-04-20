package ll2

import (
	"testing"

	"github.com/sko00o/leetcode-adventure/linked-list/dll"
	"github.com/sko00o/leetcode-adventure/linked-list/test"
)

func TestLinkedList(t *testing.T) {
	f := func() dll.LinkedList {
		c := Constructor()
		return &c
	}
	test.CommonTest(t, f)
}
