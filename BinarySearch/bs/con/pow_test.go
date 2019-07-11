package con

import (
	"fmt"
	"testing"
)

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"example 1",
			args{
				2.00000, 10,
			},
			1024.00000,
		},
		{
			"example 2",
			args{
				2.10000, 3,
			},
			9.26100,
		},
		{
			"example 3",
			args{
				2.00000, -2,
			},
			0.25000,
		},
		{
			"example 4",
			args{
				8.88023,
				3,
			},
			700.28148,
		},
		{
			"example 5",
			args{
				-3.00000,
				3,
			},
			-27.00000,
		},
	}

	for idx, f := range []func(float64, int) float64{myPow, myPow1} {
		t.Run(fmt.Sprintf("func #%d", idx), func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := f(tt.args.x, tt.args.n); got-tt.want >= 0.00001 { // precision fix
						t.Errorf("got: %v, want: %v", got, tt.want)
					}
				})
			}
		})
	}
}
