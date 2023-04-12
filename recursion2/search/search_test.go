package search

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_search_col(t *testing.T) {
	tests := []struct {
		matric    [][]int
		target    int
		wantIdx   int
		wantExist bool
	}{
		{
			matric: [][]int{
				{-9, -8, -7, -5, -3, -2, -1, 0, 1, 2, 4, 5, 6, 9},
			},
			target:    -6,
			wantIdx:   3,
			wantExist: false,
		},
		{
			matric: [][]int{
				{-9, -8, -7, -5, -3, -2, -1, 0, 1, 2, 4, 5, 6, 9},
			},
			target:    -5,
			wantIdx:   3,
			wantExist: true,
		},
		{
			matric: [][]int{
				{-9, -8, -7, -5, -3, -2, -1, 0, 1, 2, 4, 5, 6, 9},
			},
			target:    -4,
			wantIdx:   4,
			wantExist: false,
		},
		{
			matric: [][]int{
				{-9, -8, -7, -5, -3, -2, -1, 0, 1},
			},
			target:    -6,
			wantIdx:   3,
			wantExist: false,
		},
		{
			matric: [][]int{
				{-9, -8, -7, -5, -3, -2, -1, 0, 1},
			},
			target:    -5,
			wantIdx:   3,
			wantExist: true,
		},
		{
			matric: [][]int{
				{-9, -8, -7, -5, -3, -2, -1, 0, 1},
			},
			target:    -4,
			wantIdx:   4,
			wantExist: false,
		},
		{
			matric: [][]int{
				{-9, -8, -7, -5, -3, -2, -1, 0},
			},
			target:    -6,
			wantIdx:   3,
			wantExist: false,
		},
		{
			matric: [][]int{
				{-9, -8, -7, -5, -3, -2, -1, 0},
			},
			target:    -5,
			wantIdx:   3,
			wantExist: true,
		},
		{
			matric: [][]int{
				{-9, -8, -7, -5, -3, -2, -1, 0},
			},
			target:    -4,
			wantIdx:   4,
			wantExist: false,
		},
		{
			matric: [][]int{
				{-9, -8, -7, -5},
			},
			target:    -6,
			wantIdx:   3,
			wantExist: false,
		},
		{
			matric: [][]int{
				{-9, -8, -7, -5},
			},
			target:    -5,
			wantIdx:   3,
			wantExist: true,
		},
		{
			matric: [][]int{
				{-9, -8, -7, -5},
			},
			target:    -4,
			wantIdx:   4,
			wantExist: false,
		},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("task#%d", i), func(t *testing.T) {
			s := SearchRow{
				Matrix: tc.matric,
				Row:    0,
			}
			idx, exist := search(s, tc.target)
			require.Equal(t, tc.wantExist, exist)
			require.Equal(t, tc.wantIdx, idx)
		})
	}
}
