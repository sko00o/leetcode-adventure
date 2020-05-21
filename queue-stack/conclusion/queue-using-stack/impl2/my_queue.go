package impl2

import "github.com/sko00o/leetcode-adventure/queue-stack/stack"

// Stack is a LIFO Data Structure.
type Stack struct {
	stack.SliceStack
}

// MyQueue is a queue using stack.
type MyQueue struct {
	S     [2]Stack
	front int
}

// Constructor return MyQueue object.
func Constructor() MyQueue {
	return MyQueue{
		S: [2]Stack{},
	}
}

// Push element x to the back of queue.
// time complexity: O(1)
// space complexity : O(n), we need additional memory to store the queue elements.
func (q *MyQueue) Push(x int) {
	if q.S[0].IsEmpty() {
		q.front = x
	}
	q.S[0].Push(x)
}

// Pop removes the element from in front of queue and returns that element.
// time complexity: Amortized O(1), Worst-case O(n)
// space complexity: O(1)
func (q *MyQueue) Pop() int {
	if q.Empty() {
		return -1
	}

	if q.S[1].IsEmpty() {
		for !q.S[0].IsEmpty() {
			q.S[1].Push(q.S[0].Top())
			q.S[0].Pop()
		}
	}

	res := q.S[1].Top().(int)
	q.S[1].Pop()
	return res
}

// Peek get the front element.
// time complexity: O(1)
// space complexity: O(1)
func (q *MyQueue) Peek() int {
	if q.Empty() {
		return -1
	}

	if q.S[1].IsEmpty() {
		return q.front
	}

	return q.S[1].Top().(int)
}

// Empty returns whether the queue is empty.
// time complexity: O(1)
// space complexity: O(1)
func (q *MyQueue) Empty() bool {
	return q.S[0].IsEmpty() && q.S[1].IsEmpty()
}
