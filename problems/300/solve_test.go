package problems

import (
	"fmt"
	"testing"
)

func Test_lengthOfLTS(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "test 1",
			nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			want: 4, // [2,3,7,101]
		},
		{
			name: "test 2",
			nums: []int{0, 1, 0, 3, 2, 3},
			want: 4, // [0,1,2,3]
		},
		{
			name: "test 3",
			nums: []int{7, 7, 7, 7, 7, 7, 7},
			want: 1, // [7]
		},
	}
	for _, tt := range tests {
		for idx, fn := range []func([]int) int{lengthOfLIS, lengthOfLIS1} {
			t.Run(tt.name+fmt.Sprintf("(%d)", idx), func(t *testing.T) {
				if got := fn(tt.nums); got != tt.want {
					t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}
