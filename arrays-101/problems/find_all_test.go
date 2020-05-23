package problems

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_findDisappearedNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{
				nums: []int{4, 3, 2, 7, 8, 2, 3, 1},
			},
			want: []int{5, 6},
		},
	}

	for idx, f := range []func([]int) []int{
		findDisappearedNumbers,
		findDisappearedNumbers1,
		findDisappearedNumbers2,
	} {
		t.Run(fmt.Sprintf("func#%d", idx), func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := f(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
						t.Errorf("findDisappearedNumbers() = %v, want %v", got, tt.want)
					}
				})
			}
		})
	}
}
