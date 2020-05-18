package conclusion

import (
	"fmt"
	"testing"

	"github.com/sko00o/leetcode-adventure/queue-stack/conclusion/stack-using-queue/impl1"
	"github.com/sko00o/leetcode-adventure/queue-stack/conclusion/stack-using-queue/impl2"
	"github.com/sko00o/leetcode-adventure/queue-stack/conclusion/stack-using-queue/impl3"

	"github.com/stretchr/testify/require"
)

type MyStack interface {
	Push(int)
	Pop() int
	Top() int
	Empty() bool
}

func TestMyStack(t *testing.T) {
	for idx, Constructor := range []func() MyStack{
		func() MyStack {
			q := impl1.Constructor()
			return &q
		},
		func() MyStack {
			q := impl2.Constructor()
			return &q
		},
		func() MyStack {
			q := impl3.Constructor()
			return &q
		},
	} {
		t.Run(fmt.Sprintf("impl%d", idx+1), func(t *testing.T) {
			s := Constructor()
			assert := require.New(t)
			assert.True(s.Empty())
			s.Push(1)
			assert.False(s.Empty())
			assert.Equal(1, s.Top())
			s.Push(2)
			assert.Equal(2, s.Top())
			assert.Equal(2, s.Pop())
			assert.Equal(1, s.Top())
			assert.Equal(1, s.Pop())
			assert.True(s.Empty())
			assert.Equal(-1, s.Top())
			assert.Equal(-1, s.Pop())
		})
	}
}
