package queue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCircularQueue(t *testing.T) {
	for _, q := range []CircularQueue{
		NewMyCircularQueue(3), // set the size to be 3
	} {
		t.Run("common test", func(t *testing.T) {
			assert := require.New(t)
			assert.True(q.EnQueue(1))  // return true
			assert.True(q.EnQueue(2))  // return true
			assert.True(q.EnQueue(3))  // return true
			assert.False(q.EnQueue(4)) // return false, the queue is full
			assert.Equal(3, q.Rear())  // return 3
			assert.True(q.IsFull())    // return true
			assert.True(q.DeQueue())   // return true
			assert.True(q.EnQueue(4))  // return true
			assert.Equal(4, q.Rear())  // return 4
		})
	}
}
