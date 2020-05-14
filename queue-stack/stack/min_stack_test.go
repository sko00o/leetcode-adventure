package stack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinStack(t *testing.T) {
	assert := require.New(t)
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	assert.Equal(-3, minStack.GetMin()) // return -3
	minStack.Pop()
	assert.Equal(0, minStack.Top())     // return 0
	assert.Equal(-2, minStack.GetMin()) // return -2
}
