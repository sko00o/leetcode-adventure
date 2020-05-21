package problems

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func Test_sortArrayByParity(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{
				A: []int{3, 1, 2, 4},
			},
			want: []int{2, 4, 3, 1},
		},
	}

	for idx, f := range []func([]int) []int{
		sortArrayByParity,
		sortArrayByParity1,
	} {
		t.Run(fmt.Sprintf("func#%d", idx), func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					var A = make([]int,  len(tt.args.A)) 
					copy(A, tt.args.A)
					got := f(A)
					order(got)
					order(tt.want)
					if !reflect.DeepEqual(got, tt.want) {
						t.Errorf("sortArrayByParity() = %v, want %v", got, tt.want)
					}
				})
			}
		})
	}
}

func order(A []int) {
	var i = 0
	for i < len(A) && A[i]%2 == 0 {
		i++
	}

	sort.Sort(sort.IntSlice(A[:i]))
	sort.Sort(sort.IntSlice(A[i:]))
}
