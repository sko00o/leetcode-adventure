package practice

import (
	"fmt"
	"testing"
)

func Test_evalRPN(t *testing.T) {
	type args struct {
		tokens []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				tokens: []string{
					"2", "1", "+", "3", "*",
				},
			},
			want: 9,
		},
		{
			name: "example 2",
			args: args{
				tokens: []string{
					"4", "13", "5", "/", "+",
				},
			},
			want: 6,
		},
		{
			name: "example 3",
			args: args{
				tokens: []string{
					"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+",
				},
			},
			want: 22,
		},
	}

	for idx, f := range []func([]string) int{
		evalRPN,
		evalRPN1,
	} {
		t.Run(fmt.Sprintf("func#%d", idx), func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := f(tt.args.tokens); got != tt.want {
						t.Errorf("evalRPN() = %v, want %v", got, tt.want)
					}
				})
			}
		})
	}
}
