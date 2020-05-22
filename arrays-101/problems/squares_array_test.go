package problems

import (
	"fmt"
	"reflect"
	"testing"
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
