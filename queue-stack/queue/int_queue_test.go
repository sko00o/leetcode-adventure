package queue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntQueue(t *testing.T) {
	for _, q := range []IntQueue{
		new(myIntQueue),
		NewMyCircularQueue(2),
	} {
		t.Run("common test", func(t *testing.T) {
			assert := require.New(t)
			assert.Equal(-1, q.Front())
			assert.True(q.IsEmpty())
			assert.True(q.EnQueue(1))
			assert.False(q.IsEmpty())
			assert.Equal(1, q.Front())
			assert.True(q.EnQueue(2))
			assert.False(q.IsEmpty())
			assert.Equal(1, q.Front())
			assert.True(q.DeQueue())
			assert.Equal(2, q.Front())
			assert.True(q.DeQueue())
			assert.True(q.IsEmpty())
			assert.False(q.DeQueue())
			assert.Equal(-1, q.Front())
		})
	}
}
