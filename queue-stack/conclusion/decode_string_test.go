package conclusion

import (
	"fmt"
	"testing"
)

func Test_decodeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example 1",
			args: args{
				s: "3[a]2[bc]",
			},
			want: "aaabcbc",
		},
		{
			name: "example 2",
			args: args{
				s: "3[a2[c]]",
			},
			want: "accaccacc",
		},
		{
			name: "example 3",
			args: args{
				s: "2[abc]3[cd]ef",
			},
			want: "abcabccdcdcdef",
		},
		{
			name: "test 1",
			args: args{
				s: "10[a]",
			},
			want: "aaaaaaaaaa",
		},
	}
	for idx, f := range []func(string) string{
		decodeString,
		decodeString1,
	} {
		t.Run(fmt.Sprintf("func#%d", idx), func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := f(tt.args.s); got != tt.want {
						t.Errorf("decodeString() = %v, want %v", got, tt.want)
					}
				})
			}
		})
	}
}
