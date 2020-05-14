package queue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueue(t *testing.T) {
	var q Queue
	assert := require.New(t)
	assert.True(q.IsEmpty())
	assert.True(q.EnQueue(1))
	assert.False(q.IsEmpty())
	assert.True(q.EnQueue("hello"))
	assert.False(q.IsEmpty())
	assert.Equal(1, q.Front())
	assert.True(q.DeQueue())
	assert.Equal("hello", q.Front())
	assert.True(q.DeQueue())
	assert.True(q.IsEmpty())
	assert.False(q.DeQueue())
}
