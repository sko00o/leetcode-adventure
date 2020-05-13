package queue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueue(t *testing.T) {
	for _, q := range []Queue{
		new(myQueue),
		MyCircularQueue(2),
	} {
		t.Run("common test", func(t *testing.T) {
			assert := require.New(t)
			assert.True(q.isEmpty())
			assert.True(q.enQueue(1))
			assert.False(q.isEmpty())
			assert.True(q.enQueue(2))
			assert.False(q.isEmpty())
			assert.Equal(1, q.Front())
			assert.True(q.deQueue())
			assert.Equal(2, q.Front())
			assert.True(q.deQueue())
			assert.True(q.isEmpty())
			assert.False(q.deQueue())
		})
	}
}
