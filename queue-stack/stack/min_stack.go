package stack

// MinStack is a Stack that supports push, pop, top,
// and retrieving the minimum element in constant time.
type MinStack struct {
	n []node
}

type node struct {
	data int
	min  int
}

// Constructor initialize a min stack,
func Constructor() MinStack {
	return MinStack{}
}

// Push insert an element into the stack.
func (s *MinStack) Push(x int) {
	min := x
	if !s.IsEmpty() {
		if m := s.n[len(s.n)-1].min; min > m {
			min = m
		}
	}

	s.n = append(s.n, node{
		data: x,
		min:  min,
	})
}

// Pop delete an element from the stack.
func (s *MinStack) Pop() {
	if s.IsEmpty() {
		return
	}
	s.n = s.n[:len(s.n)-1]
}

// Top get the top item from the stack.
func (s *MinStack) Top() int {
	if s.IsEmpty() {
		return -1
	}
	return s.n[len(s.n)-1].data
}

// GetMin get the minimum item from the stack.
func (s *MinStack) GetMin() int {
	if s.IsEmpty() {
		return -1
	}
	return s.n[len(s.n)-1].min
}

// IsEmpty checks whether the stack is empty or not.
func (s *MinStack) IsEmpty() bool {
	return len(s.n) == 0
}
