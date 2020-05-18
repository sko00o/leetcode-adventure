package conclusion

import "github.com/sko00o/leetcode-adventure/queue-stack/stack"

// Stack is a LIFO Data Structure.
type Stack struct {
	stack.Stack
}

// MyQueue is a queue using stack.
type MyQueue struct {
	S [2]Stack
}

// Constructor return MyQueue object.
func Constructor() MyQueue {
	return MyQueue{
		S: [2]Stack{},
	}
}

// Push element x to the back of queue.
func (q *MyQueue) Push(x int) {
	for !q.S[1].IsEmpty() {
		q.S[0].Push(q.S[1].Top())
		q.S[1].Pop()
	}

	q.S[0].Push(x)
}

// Pop removes the element from in front of queue and returns that element.
func (q *MyQueue) Pop() int {
	if q.Empty() {
		return -1
	}

	for !q.S[0].IsEmpty() {
		q.S[1].Push(q.S[0].Top())
		q.S[0].Pop()
	}

	res := q.S[1].Top().(int)
	q.S[1].Pop()
	return res
}

// Peek get the front element.
func (q *MyQueue) Peek() int {
	if q.Empty() {
		return -1
	}

	for !q.S[0].IsEmpty() {
		q.S[1].Push(q.S[0].Top())
		q.S[0].Pop()
	}

	return q.S[1].Top().(int)
}

// Empty returns whether the queue is empty.
func (q *MyQueue) Empty() bool {
	return q.S[0].IsEmpty() && q.S[1].IsEmpty()
}
