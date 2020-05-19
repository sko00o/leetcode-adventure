package problems

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_deserialize(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *NestedInteger
	}{
		{
			name: "empty list 1",
			args: args{
				s: "[[]]",
			},
			want: &NestedInteger{
				[]*NestedInteger{
					{},
				},
			},
		},
		{
			name: "simple integer",
			args: args{
				s: "123",
			},
			want: &NestedInteger{
				int(123),
			},
		},
		{
			name: "simple neg integer",
			args: args{
				s: "-123",
			},
			want: &NestedInteger{
				int(-123),
			},
		},
		{
			name: "nested",
			args: args{
				s: "[123,[456,[789]]]",
			},
			want: &NestedInteger{
				[]*NestedInteger{
					{
						int(123),
					},
					{
						[]*NestedInteger{
							{
								int(456),
							},
							{
								[]*NestedInteger{
									{
										int(789),
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "simple list",
			args: args{
				s: "[0,[1,2,3],[]]",
			},
			want: &NestedInteger{
				[]*NestedInteger{
					{int(0)},
					{
						[]*NestedInteger{
							{int(1)},
							{int(2)},
							{int(3)},
						},
					},
					{},
				},
			},
		},
		{
			name: "empty list",
			args: args{
				s: "[]",
			},
			want: &NestedInteger{},
		},
	}

	for idx, f := range []func(string) *NestedInteger{
		deserialize,
		deserialize1,
	} {
		t.Run(fmt.Sprintf("func#%d", idx), func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					assert := require.New(t)
					got := f(tt.args.s)
					assert.Equal(tt.want, got)
				})
			}
		})
	}
}
