package impl2

import "github.com/sko00o/leetcode-adventure/queue-stack/queue"

/*
Notes:
○ 用队列模拟栈
	§ 若使用两个队列，分别记为 Q 和 tmp ，压栈操作时，元素入队 tmp，
	  然后 Q 元素全部出队再入队 tmp， 最后交换 Q 和 tmp；出栈操作时，
	  保持所有元素都在 Q，方便维护，从 Q 出队。
*/

// Queue is a FIFO Data Structure.
type Queue struct {
	queue.SliceQueue
}

// MyStack is a stack using queue.
type MyStack struct {
	Q   *Queue
	tmp *Queue
}

// Constructor return MyStack object.
func Constructor() MyStack {
	return MyStack{
		Q:   &Queue{},
		tmp: &Queue{},
	}
}

// Push element x onto stack.
// time complexity: O(n)
// space complexity : O(1)
func (s *MyStack) Push(x int) {
	s.tmp.EnQueue(x)
	for !s.Q.IsEmpty() {
		s.tmp.EnQueue(s.Q.Front())
		s.Q.DeQueue()
	}

	s.Q, s.tmp = s.tmp, s.Q
}

// Pop removes the element on top of the stack and returns that element.
// time complexity: O(1)
// space complexity : O(1)
func (s *MyStack) Pop() int {
	if s.Q.IsEmpty() {
		return -1
	}

	res := s.Q.Front().(int)
	s.Q.DeQueue()
	return res
}

// Top get the top element.
// time complexity: O(1)
// space complexity : O(1)
func (s *MyStack) Top() int {
	if s.Q.IsEmpty() {
		return -1
	}

	return s.Q.Front().(int)
}

// Empty returns whether the queue is empty.
// time complexity: O(1)
// space complexity : O(1)
func (s *MyStack) Empty() bool {
	return s.Q.IsEmpty()
}
