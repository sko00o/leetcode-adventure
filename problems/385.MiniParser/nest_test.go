package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNestedInteger(t *testing.T) {
	assert := require.New(t)

	var ni, ni1 NestedInteger

	assert.False(ni.IsInteger())
	ni.SetInteger(1)
	assert.True(ni.IsInteger())
	ni1.SetInteger(2)
	assert.True(ni1.IsInteger())
	assert.Len(ni1.GetList(), 0)

	ni.Add(ni1)
	assert.False(ni.IsInteger())
	assert.Equal([]*NestedInteger{
		{data: 1},
		{data: 2},
	}, ni.GetList())

	ni.Add(NestedInteger{
		data: []*NestedInteger{
			{data: 1},
		},
	})
	assert.Equal([]*NestedInteger{
		{data: 1},
		{data: 2},
		{
			data: []*NestedInteger{
				{data: 1},
			},
		},
	}, ni.GetList())
}
