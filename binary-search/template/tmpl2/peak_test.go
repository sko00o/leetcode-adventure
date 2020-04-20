package tmpl2

import (
	"testing"
)

func Test_findPeakElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		ok   map[int]bool
	}{
		{
			name: "one peak1",
			args: args{
				nums: []int{1, 2, 3, 1},
			},
			ok: map[int]bool{
				2: true,
			},
		},
		{
			name: "one peak2",
			args: args{
				nums: []int{1, 2},
			},
			ok: map[int]bool{
				1: true,
			},
		},
		{
			name: "one peak3",
			args: args{
				nums: []int{3, 2},
			},
			ok: map[int]bool{
				0: true,
			},
		},
		{
			name: "one peak4",
			args: args{
				nums: []int{6},
			},
			ok: map[int]bool{
				0: true,
			},
		},
		{
			name: "two peaks",
			args: args{
				nums: []int{1, 2, 1, 3, 5, 6, 4},
			},
			ok: map[int]bool{
				1: true,
				5: true,
			},
		},
		{
			name: "nil input",
			args: args{
				nums: []int{},
			},
			ok: map[int]bool{
				-1: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPeakElement(tt.args.nums); tt.ok[got] != true {
				t.Errorf("findPeakElement() = %v, expect in %v", got, tt.ok)
			}
		})
	}
}
