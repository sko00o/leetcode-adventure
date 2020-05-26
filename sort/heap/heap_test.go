package heap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxHeapSort(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{
			arr:  []int{4, 2, 5, 1, 3},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			arr:  []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3},
			want: []int{49, 25, 14, 9, 8, 6, 6, 6, 5, 4, 3, 1},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case#%d", i), func(t *testing.T) {
			assert := require.New(t)

			h := NewMaxHeap(len(tt.arr))
			for idx, n := range tt.arr {
				h.Push(n)
				assert.Equal(idx+1, h.Size())
			}

			var got []int
			var cnt = len(tt.arr)
			for !h.IsEmpty() {
				got = append(got, h.Top())
				h.Pop()
				cnt--
				assert.Equal(cnt, h.Size())
			}
			assert.Equal(tt.want, got)

			assert.False(h.Pop())
			assert.Equal(-1, h.Top())
		})
	}
}

func TestMinHeapSort(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{
			arr:  []int{4, 2, 5, 1, 3},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			arr:  []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3},
			want: []int{1, 3, 4, 5, 6, 6, 6, 8, 9, 14, 25, 49},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case#%d", i), func(t *testing.T) {
			assert := require.New(t)

			h := NewMinHeap(len(tt.arr))
			for _, n := range tt.arr {
				h.Push(n)
			}

			var got []int
			for !h.IsEmpty() {
				got = append(got, h.Top())
				h.Pop()
			}
			assert.Equal(tt.want, got)

			assert.False(h.Pop())
			assert.Equal(-1, h.Top())
		})
	}
}
