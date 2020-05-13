package queue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueue(t *testing.T) {
	for _, q := range []Queue{
		new(myQueue),
		NewMyCircularQueue(2),
	} {
		t.Run("common test", func(t *testing.T) {
			assert := require.New(t)
			assert.True(q.IsEmpty())
			assert.True(q.EnQueue(1))
			assert.False(q.IsEmpty())
			assert.True(q.EnQueue(2))
			assert.False(q.IsEmpty())
			assert.Equal(1, q.Front())
			assert.True(q.DeQueue())
			assert.Equal(2, q.Front())
			assert.True(q.DeQueue())
			assert.True(q.IsEmpty())
			assert.False(q.DeQueue())
		})
	}
}
