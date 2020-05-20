package problems

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_duplicateZeros(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{
				arr: []int{1, 0, 2, 3, 0, 4, 5, 0},
			},
			want: []int{1, 0, 0, 2, 3, 0, 0, 4},
		},
		{
			name: "example 2",
			args: args{
				arr: []int{1, 2, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "test 1",
			args: args{
				arr: []int{1, 1, 0},
			},
			want: []int{1, 1, 0},
		},
		{
			name: "test 2",
			args: args{
				arr: []int{0, 0, 0},
			},
			want: []int{0, 0, 0},
		},
		{
			name: "test 3",
			args: args{
				arr: []int{1, 2, 3, 0, 0, 0, 0, 4},
			},
			want: []int{1, 2, 3, 0, 0, 0, 0, 0},
		},
	}

	for idx, f := range []func([]int){
		duplicateZeros,
		duplicateZeros1,
	} {
		t.Run(fmt.Sprintf("func#%d", idx), func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					var arr = make([]int, len(tt.args.arr))
					copy(arr, tt.args.arr)
					assert := require.New(t)
					f(arr)
					assert.Equal(tt.want, arr)
				})
			}
		})
	}
}
