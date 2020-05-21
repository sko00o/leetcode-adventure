package queue

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueue(t *testing.T) {
	for idx, q := range []Queue{
		&SliceQueue{},
	} {
		t.Run(fmt.Sprintf("Queue#%d", idx), func(t *testing.T) {
			assert := require.New(t)
			assert.Nil(q.Front())
			assert.True(q.IsEmpty())
			assert.True(q.EnQueue(1))
			assert.False(q.IsEmpty())
			assert.Equal(1, q.Front())
			assert.True(q.EnQueue("hello"))
			assert.False(q.IsEmpty())
			assert.Equal(1, q.Front())
			assert.True(q.DeQueue())
			assert.Equal("hello", q.Front())
			assert.True(q.DeQueue())
			assert.True(q.IsEmpty())
			assert.False(q.DeQueue())
			assert.Nil(q.Front())

			// currency test
			var popFailedCnt int
			var popResChan = make(chan bool)
			var wg sync.WaitGroup
			go func() {
				for r := range popResChan {
					if !r {
						popFailedCnt++
					}
				}
			}()
			for i := 1000; i > 0; i-- {
				wg.Add(2)
				go func() {
					q.EnQueue(1)
					wg.Done()
				}()
				go func() {
					popResChan <- q.DeQueue()
					wg.Done()
				}()
			}
			wg.Wait()
			assert.Equal(popFailedCnt, q.Size())
		})
	}
}
