package problems

import (
	"fmt"
	"testing"
)

func Test_strStr(t *testing.T) {
	type args struct {
		S string
		W string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				S: "sadbutsad",
				W: "sad",
			},
			want: 0,
		},
		{
			name: "test 2",
			args: args{
				S: "leetcode",
				W: "leeto",
			},
			want: -1,
		},
		{
			name: "test 3",
			args: args{
				S: "abccba",
				W: "cb",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.S, tt.args.W); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_strStr(b *testing.B) {
	fs := []func(string, string) int{strStr, strStr1}

	for _, f := range fs {
		b.Run(fmt.Sprintf("func %T", f), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				f("sadbutsad", "sad")
			}
		})
	}
}
