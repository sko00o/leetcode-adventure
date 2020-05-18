package conclusion

import (
	"testing"

	"github.com/sko00o/leetcode-adventure/queue-stack/conclusion/queue-using-stack/impl1"
	"github.com/sko00o/leetcode-adventure/queue-stack/conclusion/queue-using-stack/impl2"

	"github.com/stretchr/testify/require"
)

type MyQueue interface {
	Push(int)
	Peek() int
	Pop() int
	Empty() bool
}

func TestMyQueue(t *testing.T) {
	for _, Constructor := range []func() MyQueue{
		func() MyQueue {
			q := impl1.Constructor()
			return &q
		},
		func() MyQueue {
			q := impl2.Constructor()
			return &q
		},
	} {
		q := Constructor()
		assert := require.New(t)
		assert.Equal(-1, q.Peek())
		assert.True(q.Empty())
		q.Push(1)
		assert.False(q.Empty())
		assert.Equal(1, q.Peek())
		q.Push(2)
		assert.False(q.Empty())
		assert.Equal(1, q.Peek())
		assert.Equal(1, q.Pop())
		assert.Equal(2, q.Peek())
		q.Push(3)
		q.Push(4)
		assert.Equal(2, q.Pop())
		assert.Equal(3, q.Peek())
		assert.False(q.Empty())
		assert.Equal(3, q.Pop())
		assert.Equal(4, q.Peek())
		assert.False(q.Empty())
		assert.Equal(4, q.Pop())
		assert.True(q.Empty())
		assert.Equal(-1, q.Peek())
	}
}
