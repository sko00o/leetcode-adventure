package ll1

import (
	"testing"

	"adventure/LinkedList/dll"
	"adventure/LinkedList/dll/test"
)

func TestLinkedList(t *testing.T) {
	f := func() dll.LinkedList {
		c := Constructor()
		return &c
	}
	test.CommonTest(t, f)
}
