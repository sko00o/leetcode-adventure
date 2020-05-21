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
