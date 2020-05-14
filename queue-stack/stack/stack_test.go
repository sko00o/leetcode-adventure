package stack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStack(t *testing.T) {
	assert := require.New(t)
	var s Stack
	assert.True(s.IsEmpty())
	assert.Nil(s.Top())
	s.Push(1)
	assert.Equal(1, s.Top())
	assert.False(s.IsEmpty())
	s.Push('a')
	assert.Equal('a', s.Top())
	assert.True(s.Pop())
	assert.Equal(1, s.Top())
	assert.True(s.Pop())
	assert.True(s.IsEmpty())
	assert.False(s.Pop())
	assert.Nil(s.Top())
}
