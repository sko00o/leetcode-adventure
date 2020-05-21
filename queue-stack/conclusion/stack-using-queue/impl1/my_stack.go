package impl1

import "github.com/sko00o/leetcode-adventure/queue-stack/queue"

// Queue is a FIFO Data Structure.
type Queue struct {
	queue.SliceQueue
}

// MyStack is a stack using queue.
type MyStack struct {
	Q   *Queue
	tmp *Queue
	top int
}

// Constructor return MyStack object.
func Constructor() MyStack {
	return MyStack{
		Q:   &Queue{},
		tmp: &Queue{},
	}
}

// Push element x onto stack.
// time complexity: O(1)
// space complexity : O(1)
func (s *MyStack) Push(x int) {
	s.top = x
	s.Q.EnQueue(s.top)
}

// Pop removes the element on top of the stack and returns that element.
// time complexity: O(n)
// space complexity : O(1)
func (s *MyStack) Pop() int {
	if s.Empty() {
		return -1
	}

	for len(s.Q.Data) > 1 {
		s.top = s.Q.Front().(int)
		s.Q.DeQueue()
		s.tmp.EnQueue(s.top)
	}
	res := s.Q.Front().(int)
	s.Q.DeQueue()
	s.Q, s.tmp = s.tmp, s.Q
	return res
}

// Top get the top element.
// time complexity: O(1)
// space complexity : O(1)
func (s *MyStack) Top() int {
	if s.Empty() {
		return -1
	}

	return s.top
}

// Empty returns whether the queue is empty.
// time complexity: O(1)
// space complexity : O(1)
func (s *MyStack) Empty() bool {
	return s.Q.IsEmpty()
}
