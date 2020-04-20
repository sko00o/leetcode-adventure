package prac

import (
	"reflect"
	"sort"
	"testing"
)

func Test_intersection(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"ex1",
			args{
				[]int{1, 2, 2, 1},
				[]int{2, 2},
			},
			[]int{2},
		},
		{
			"ex2",
			args{
				[]int{4, 9, 5},
				[]int{9, 4, 9, 8, 4},
			},
			[]int{9, 4},
		},
	}
	for fIdx, f := range []func([]int, []int) []int{intersection, intersection1} {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// result can be any order
				got := f(tt.args.nums1, tt.args.nums2)
				sort.Ints(got)
				sort.Ints(tt.want)
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("func #%d, got: %v, want: %v", fIdx, got, tt.want)
				}
			})
		}
	}
}
