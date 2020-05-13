package queue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCircularQueue(t *testing.T) {
	for _, q := range []CircularQueue{
		MyCircularQueue(3), // set the size to be 3
	} {
		t.Run("common test", func(t *testing.T) {
			assert := require.New(t)
			assert.True(q.enQueue(1))  // return true
			assert.True(q.enQueue(2))  // return true
			assert.True(q.enQueue(3))  // return true
			assert.False(q.enQueue(4)) // return false, the queue is full
			assert.Equal(3, q.Rear())  // return 3
			assert.True(q.isFull())    // return true
			assert.True(q.deQueue())   // return true
			assert.True(q.enQueue(4))  // return true
			assert.Equal(4, q.Rear())  // return 4
		})
	}
}
