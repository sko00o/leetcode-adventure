package mem

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"ex1",
			args{2},
			2,
		},
		{
			"ex2",
			args{3},
			3,
		},
		{
			"ex3",
			args{4},
			5,
		},
	}

	for fIdx, f := range []func(int) int{
		climbStairs,
		climbStairs1,
		climbStairs2,
		climbStairs3,
		climbStairs4,
		climbStairs5,
	} {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.n); got != tt.want {
					t.Errorf("climbStairs #%d, got: %v, want: %v", fIdx, got, tt.want)
				}
			})
		}
	}
}
