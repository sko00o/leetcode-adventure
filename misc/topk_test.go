package misc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TopK(t *testing.T) {
	h := New(3)
	h.Insert(1)
	h.Insert(4)
	h.Insert(3)
	h.Insert(2)
	h.Insert(5)
	assert.Equal(t, 3+4+5, h.Sum())
}
