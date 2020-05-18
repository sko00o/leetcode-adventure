package conclusion

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMyQueue(t *testing.T) {
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
	assert.Equal(2, q.Pop())
	assert.True(q.Empty())
	assert.Equal(-1, q.Pop())
	assert.Equal(-1, q.Peek())
}
