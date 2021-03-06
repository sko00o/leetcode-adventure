package impl3

import "github.com/sko00o/leetcode-adventure/queue-stack/queue"

/*
Notes:
○ 用队列模拟栈
	§ 可以只用一个队列，压栈操作时，判断队列长度是否大于 1，
	  如果大于1，将n-1个元素出队再入队，以达到将队尾元素排到队头的效果。
	  出栈操作时，直接出队即可。压栈操作的时间复杂度是 O(n) 。
*/

// Queue is a FIFO Data Structure.
type Queue struct {
	queue.SliceQueue
}

// MyStack is a stack using queue.
type MyStack struct {
	Q Queue
}

// Constructor return MyStack object.
func Constructor() MyStack {
	return MyStack{
		Q: Queue{},
	}
}

// Push element x onto stack.
// time complexity: O(n)
// space complexity : O(1)
func (s *MyStack) Push(x int) {
	s.Q.EnQueue(x)

	for i := len(s.Q.Data); i > 1; i-- {
		s.Q.EnQueue(s.Q.Front())
		s.Q.DeQueue()
	}
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
