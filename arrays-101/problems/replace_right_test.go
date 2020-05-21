package problems

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_replaceElements(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{
				arr: []int{17, 18, 5, 4, 6, 1},
			},
			want: []int{18, 6, 6, 6, 1, -1},
		},
	}
	for idx, f := range []func([]int) []int{
		replaceElements,
	} {
		t.Run(fmt.Sprintf("func#%d", idx), func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					var arr = make([]int, len(tt.args.arr))
					copy(arr, tt.args.arr)
					if got := f(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
						t.Errorf("replaceElements() = %v, want %v", got, tt.want)
					}
				})
			}
		})
	}
}
