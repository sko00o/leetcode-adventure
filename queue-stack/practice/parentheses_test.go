package practice

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{
				s: "()",
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				s: "()[]{}",
			},
			want: true,
		},
		{
			name: "example 3",
			args: args{
				s: "(]",
			},
			want: false,
		},
		{
			name: "example 4",
			args: args{
				s: "([)]",
			},
			want: false,
		},
		{
			name: "example 5",
			args: args{
				s: "{[]}",
			},
			want: true,
		},
	}

	for _, f := range []func(string) bool{
		isValid,
		isValid1,
	} {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.s); got != tt.want {
					t.Errorf("isValid() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}
