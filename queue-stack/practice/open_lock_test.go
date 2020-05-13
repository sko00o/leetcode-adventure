package practice

import (
	"fmt"
	"testing"
)

func Test_nextStats(t *testing.T) {
	type args struct {
		s string
		d []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test 1",
			args: args{
				s: "0000",
				d: []int{0, 0, 0, 1},
			},
			want: "0001",
		},
		{
			name: "test 2",
			args: args{
				s: "0000",
				d: []int{0, 0, 1, 1},
			},
			want: "0011",
		},
		{
			name: "test 3",
			args: args{
				s: "0000",
				d: []int{0, 0, -1, 1},
			},
			want: "0091",
		},
		{
			name: "test 4",
			args: args{
				s: "0000",
				d: []int{0, -1, -1, 1},
			},
			want: "0991",
		},
		{
			name: "test 5",
			args: args{
				s: "1234",
				d: []int{99, -23, 23, -99},
			},
			want: "0965",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextStats(tt.args.s, tt.args.d); got != tt.want {
				t.Errorf("nextStats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_openLock(t *testing.T) {
	type args struct {
		deadends []string
		target   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				deadends: []string{"0201", "0101", "0102", "1212", "2002"},
				target:   "0202",
			},
			want: 6,
		},
		{
			name: "example 2",
			args: args{
				deadends: []string{"8888"},
				target:   "0009",
			},
			want: 1,
		},
		{
			name: "example 3",
			args: args{
				deadends: []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"},
				target:   "8888",
			},
			want: -1,
		},
		{
			name: "example 4",
			args: args{
				deadends: []string{"0000"},
				target:   "8888",
			},
			want: -1,
		},
	}

	for idx, f := range []func(deadends []string, target string) int{
		openLock,
		openLock1,
	} {
		t.Run(fmt.Sprintf("func#%d", idx), func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := f(tt.args.deadends, tt.args.target); got != tt.want {
						t.Errorf("openLock() = %v, want %v", got, tt.want)
					}
				})
			}
		})
	}
}
