package search

import "testing"

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				matrix: [][]int{
					{1, 4, 7, 11, 15},
					{2, 5, 8, 12, 19},
					{3, 6, 9, 16, 22},
					{10, 13, 14, 17, 24},
					{18, 21, 23, 26, 30},
				},
				target: 5,
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				matrix: [][]int{
					{1, 4, 7, 11, 15},
					{2, 5, 8, 12, 19},
					{3, 6, 9, 16, 22},
					{10, 13, 14, 17, 24},
					{18, 21, 23, 26, 30},
				},
				target: 20,
			},
			want: false,
		},
		{
			name: "Test 1",
			args: args{
				matrix: [][]int{
					{1, 3, 5},
				},
				target: 3,
			},
			want: true,
		},
		{
			name: "Test 2",
			args: args{
				matrix: [][]int{
					{1, 1},
				},
				target: 2,
			},
			want: false,
		},
		{
			name: "Empty",
			args: args{
				matrix: [][]int{},
				target: 3,
			},
			want: false,
		},
	}

	for _, f := range []func(matrix [][]int, target int) bool{
		searchMatrix,
		searchMatrix0,
		searchMatrix1,
	} {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.matrix, tt.args.target); got != tt.want {
					t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}
