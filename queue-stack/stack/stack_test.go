package stack

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStack(t *testing.T) {
	for idx, s := range []Stack{
		&SliceStack{},
	} {
		t.Run(fmt.Sprintf("Stack#%d", idx), func(t *testing.T) {
			assert := require.New(t)
			assert.True(s.IsEmpty())
			assert.Nil(s.Top())
			s.Push(1)
			assert.Equal(1, s.Top())
			assert.False(s.IsEmpty())
			assert.Equal(1, s.Size())
			s.Push('a')
			assert.Equal('a', s.Top())
			assert.Equal(2, s.Size())
			assert.True(s.Pop())
			assert.Equal(1, s.Top())
			assert.True(s.Pop())
			assert.True(s.IsEmpty())
			assert.False(s.Pop())
			assert.Nil(s.Top())

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
					s.Push(1)
					wg.Done()
				}()
				go func() {
					popResChan <- s.Pop()
					wg.Done()
				}()
			}
			wg.Wait()
			assert.Equal(popFailedCnt, s.Size())
		})
	}
}
