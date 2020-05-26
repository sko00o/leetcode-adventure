package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKthLargest(t *testing.T) {
	assert := require.New(t)

	kl := Constructor(3, []int{4, 5, 8, 2})
	assert.Equal(4, kl.Add(3))
	assert.Equal(5, kl.Add(5))
	assert.Equal(5, kl.Add(10))
	assert.Equal(8, kl.Add(9))
	assert.Equal(8, kl.Add(4))

	kl1 := Constructor(1, []int{-1})
	assert.Equal(-1, kl1.Add(-2))
	assert.Equal(2, kl1.Add(2))
	assert.Equal(2, kl1.Add(-1))

	kl2 := Constructor(1, nil)
	assert.Equal(-2, kl2.Add(-2))
	assert.Equal(-1, kl2.Add(-1))
	assert.Equal(-1, kl2.Add(-2))

	kl3 := Constructor(2, []int{0})
	assert.Equal(-1, kl3.Add(-1))
	assert.Equal(0, kl3.Add(1))
	assert.Equal(0, kl3.Add(-2))
}
