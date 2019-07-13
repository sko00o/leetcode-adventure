package problems

import "testing"

func Test_peakIndexInMountainArray(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"example 1",
			args{
				[]int{0, 1, 0},
			},
			1,
		},
		{
			"example 2",
			args{
				[]int{0, 2, 1, 0},
			},
			1,
		},
	}

	for fIdx, f := range []func([]int) int{peakIndexInMountainArray, peakIndexInMountainArray1} {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.A); got != tt.want {
					t.Errorf("func #%d, got: %v, expect: %v", fIdx, got, tt.want)
				}
			})
		}
	}
}
