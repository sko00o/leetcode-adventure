package stack

import (
	"sync"
)

// Stack is a LIFO Data Structure.
type Stack interface {
	Push(interface{})
	Pop() bool
	Top() interface{}
	IsEmpty() bool
	Size() int
}

// SliceStack is a Stack implement based on slice.
type SliceStack struct {
	Data []interface{}
	lock sync.Mutex
}

// Push insert an element into the stack.
func (s *SliceStack) Push(val interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.Data = append(s.Data, val)
}

// Pop delete an element from the stack.
// Return true if the operation is successful.
func (s *SliceStack) Pop() bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.IsEmpty() {
		return false
	}

	s.Data = s.Data[:len(s.Data)-1]
	return true
}

// Top get the top item from the stack.
func (s *SliceStack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.Data[len(s.Data)-1]
}

// IsEmpty checks whether the stack is empty or not.
func (s *SliceStack) IsEmpty() bool {
	return len(s.Data) == 0
}

// Size return length of the stack.
func (s *SliceStack) Size() int {
	return len(s.Data)
}

// LinkStack is a Stack implement based on linked list.
type LinkStack struct {
	root *LinkNode
	size int
	lock sync.Mutex
}

// LinkNode defines node structure in linked list.
type LinkNode struct {
	Next *LinkNode
	Data interface{}
}

// Push insert an element into the stack.
func (s *LinkStack) Push(val interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.root = &LinkNode{
		Data: val,
		Next: s.root,
	}

	s.size++
}

// Pop delete an element from the stack.
// Return true if the operation is successful.
func (s *LinkStack) Pop() bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.IsEmpty() {
		return false
	}

	s.root = s.root.Next
	s.size--

	return true
}

// Top get the top item from the stack.
func (s *LinkStack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}

	return s.root.Data
}

// IsEmpty checks whether the stack is empty or not.
func (s *LinkStack) IsEmpty() bool {
	return s.size == 0
}

// Size return length of the stack.
func (s *LinkStack) Size() int {
	return s.size
}
