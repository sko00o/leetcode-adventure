package problems

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sortedSquares(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "worst 1",
			args: args{
				A: []int{-2, -2, -2, -2, 3},
			},
			want: []int{4, 4, 4, 4, 9},
		},
		{
			name: "worst 2",
			args: args{
				A: []int{-5, -3, -1, 2, 4, 6},
			},
			want: []int{1, 4, 9, 16, 25, 36},
		},
		{
			name: "example 1",
			args: args{
				A: []int{-4, -1, 0, 3, 10},
			},
			want: []int{0, 1, 9, 16, 100},
		},
		{
			name: "example 2",
			args: args{
				A: []int{-7, -3, 2, 3, 11},
			},
			want: []int{4, 9, 9, 49, 121},
		},
		{
			name: "test 1",
			args: args{
				A: []int{1, 2, 3},
			},
			want: []int{1, 4, 9},
		},
		{
			name: "test 2",
			args: args{
				A: []int{-3, -2, -1},
			},
			want: []int{1, 4, 9},
		},
		{
			name: "test 3",
			args: args{
				A: []int{3},
			},
			want: []int{9},
		},
		{
			name: "test 4",
			args: args{
				A: []int{-3},
			},
			want: []int{9},
		},
		{
			name: "test 5",
			args: args{
				A: []int{-1, 0},
			},
			want: []int{0, 1},
		},
	}

	for idx, f := range []func([]int) []int{
		sortedSquares,
		sortedSquares1,
		sortedSquares2,
		sortedSquares3,
		sortedSquares4,
		sortedSquares5,
	} {
		t.Run(fmt.Sprintf("func#%d", idx), func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					var A = make([]int, len(tt.args.A))
					copy(A, tt.args.A)
					if got := f(A); !reflect.DeepEqual(got, tt.want) {
						t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
					}
				})
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test 1",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
			},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "test 2",
			args: args{
				arr: []int{2, 3, 4, 5},
			},
			want: []int{5, 4, 3, 2},
		},
		{
			name: "empty",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			var arr []int
			if tt.args.arr != nil {
				arr = make([]int, len(tt.args.arr))
				copy(arr, tt.args.arr)
			}
			reverse(arr)
			assert.Equal(tt.want, arr)
		})
	}
}

func Test_rotate(t *testing.T) {
	type args struct {
		arr []int
		mid int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test 1",
			args: args{
				arr: []int{1, 2, 3, 4, 5, 6},
				mid: 3,
			},
			want: []int{4, 5, 6, 1, 2, 3},
		},
		{
			name: "test 2",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				mid: 4,
			},
			want: []int{5, 1, 2, 3, 4},
		},
		{
			name: "test 3",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				mid: 0,
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "test 4",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				mid: 5,
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "empty",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			var arr []int
			if tt.args.arr != nil {
				arr = make([]int, len(tt.args.arr))
				copy(arr, tt.args.arr)
			}
			rotate(arr, tt.args.mid)
			assert.Equal(tt.want, arr)
		})
	}
}
