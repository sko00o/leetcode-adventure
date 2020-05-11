package problems

// NestedInteger is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
	data interface{}
}

// IsInteger return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool {
	_, ok := n.data.(int)
	return ok
}

// GetInteger return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() int {
	return n.data.(int)
}

// SetInteger set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {
	n.data = value
}

// Add set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {
	switch data := n.data.(type) {
	case []*NestedInteger:
		n.data = append(data, &elem)
	case int:
		n.data = []*NestedInteger{
			{data: data}, 
			&elem,
		}
	}
}

// GetList return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []*NestedInteger {
	data, ok := n.data.([]*NestedInteger)
	if !ok {
		return nil
	}

	return data
}
