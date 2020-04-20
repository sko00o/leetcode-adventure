package practice

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"ex1",
			args{
				[]int{2, 7, 11, 15},
				9,
			},
			[]int{1, 2},
		},
		{
			"ex2",
			args{
				[]int{1, 2, 3, 4, 4, 9, 56, 90},
				8,
			},
			[]int{4, 5},
		},
	}
	for fIdx, f := range []func([]int, int) []int{twoSum, twoSum1} {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("func #%d, got: %v, want: %v", fIdx, got, tt.want)
				}
			})
		}
	}
}
