package queue

import (
	"sync"
)

// Queue is a FIFO Data Structure.
type Queue interface {
	EnQueue(interface{}) bool
	DeQueue() bool
	Front() interface{}
	IsEmpty() bool
	Size() int
}

// SliceQueue is a Queue implement based on slice.
type SliceQueue struct {
	Data []interface{}
	lock sync.Mutex
}

// EnQueue insert an element into the queue.
// Return nil true if the operation is successful.
func (q *SliceQueue) EnQueue(val interface{}) bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.Data = append(q.Data, val)
	return true
}

// DeQueue delete an element from the queue.
// Return nil true if the operation is successful.
func (q *SliceQueue) DeQueue() bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.IsEmpty() {
		return false
	}

	q.Data = q.Data[1:]
	return true
}

// Front get the front item from the queue.
// If the queue is empty, return nil.
func (q *SliceQueue) Front() interface{} {
	if q.IsEmpty() {
		return nil
	}

	return q.Data[0]
}

// IsEmpty checks whether the queue is empty or not.
func (q *SliceQueue) IsEmpty() bool {
	return len(q.Data) == 0
}

// Size return length of the stack.
func (q *SliceQueue) Size() int {
	return len(q.Data)
}

// LinkQueue is a Queue implement based on linked list.
type LinkQueue struct {
	head, tail *LinkNode
	size       int
	lock       sync.Mutex
}

// LinkNode defines node structure in linked list.
type LinkNode struct {
	Next *LinkNode
	Data interface{}
}

// EnQueue insert an element into the queue.
// Return nil true if the operation is successful.
func (q *LinkQueue) EnQueue(val interface{}) bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.head == nil {
		q.head = &LinkNode{
			Data: val,
		}
		q.tail = q.head
	} else {
		q.tail.Next = &LinkNode{
			Data: val,
		}
		q.tail = q.tail.Next
	}

	q.size++
	return true
}

// DeQueue delete an element from the queue.
// Return nil true if the operation is successful.
func (q *LinkQueue) DeQueue() bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.IsEmpty() {
		return false
	}

	q.head = q.head.Next
	q.size--
	return true
}

// Front get the front item from the queue.
// If the queue is empty, return nil.
func (q *LinkQueue) Front() interface{} {
	if q.IsEmpty() {
		return nil
	}

	return q.head.Data
}

// IsEmpty checks whether the queue is empty or not.
func (q *LinkQueue) IsEmpty() bool {
	return q.size == 0
}

// Size return length of the stack.
func (q *LinkQueue) Size() int {
	return q.size
}
