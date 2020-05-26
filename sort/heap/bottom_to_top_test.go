package heap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_buildMaxHeap(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{
			arr:  []int{1, 4, 3, 7, 8, 9, 10},
			want: []int{10, 8, 9, 7, 4, 1, 3},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case#%d", i), func(t *testing.T) {
			assert := require.New(t)
			var arr = make([]int, len(tt.arr))
			copy(arr, tt.arr)
			buildMaxHeap(arr)
			assert.Equal(tt.want, arr)
		})
	}
}

func Test_heapSort(t *testing.T) {
	tests := []struct {
		arr      []int
		wantSort []int
	}{
		{
			arr:      []int{7, 1, 10, 4, 3, 8, 9},
			wantSort: []int{1, 3, 4, 7, 8, 9, 10},
		},
	}

	for idx, f := range []func([]int){
		heapSort,
		heapSort1,
	} {
		t.Run(fmt.Sprintf("func#%d", idx), func(t *testing.T) {
			for i, tt := range tests {
				t.Run(fmt.Sprintf("case#%d", i), func(t *testing.T) {
					assert := require.New(t)
					var arr = make([]int, len(tt.arr))
					copy(arr, tt.arr)
					f(arr)
					assert.Equal(tt.wantSort, arr)
				})
			}
		})
	}
}

func Test_buildMinHeap(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{
			arr:  []int{10, 8, 9, 7, 4, 1, 3},
			want: []int{1, 4, 3, 7, 8, 9, 10},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case#%d", i), func(t *testing.T) {
			assert := require.New(t)
			var arr = make([]int, len(tt.arr))
			copy(arr, tt.arr)
			buildMinHeap(arr)
			assert.Equal(tt.want, arr)
		})
	}
}

func Test_heapSortRevert(t *testing.T) {
	tests := []struct {
		arr      []int
		wantSort []int
	}{
		{
			arr:      []int{7, 1, 10, 4, 3, 8, 9},
			wantSort: []int{10, 9, 8, 7, 4, 3, 1},
		},
	}

	for idx, f := range []func([]int){
		heapSortRevert,
		heapSortRevert1,
	} {
		t.Run(fmt.Sprintf("func#%d", idx), func(t *testing.T) {
			for i, tt := range tests {
				t.Run(fmt.Sprintf("case#%d", i), func(t *testing.T) {
					assert := require.New(t)
					var arr = make([]int, len(tt.arr))
					copy(arr, tt.arr)
					f(arr)
					assert.Equal(tt.wantSort, arr)
				})
			}
		})
	}
}
