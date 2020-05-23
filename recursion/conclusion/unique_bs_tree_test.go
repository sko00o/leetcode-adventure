package conclusion

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTreeNode(t *testing.T) {
	n := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
			},
		},
	}

	assert := require.New(t)
	assert.Equal("[1,null,3,2]", n.String())
}

func Test_generateTrees(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		{
			name: "example 1",
			args: args{n:3},
			want: []*TreeNode{
				{
					Val: 1, 
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 2,
						},
					},
				},
				{
					Val: 1, 
					Right: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val: 3,
						},
					},
				},
				{
					Val: 2, 
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				{
					Val: 3, 
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 1,
						},
					},
				},
				{
					Val: 3, 
					Left: &TreeNode{
						Val: 1,
						Right: &TreeNode{
							Val: 2,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			var wantStr = make([]string, len(tt.want)) 
			for i := range tt.want {
				wantStr[i] = tt.want[i].String()
			}
			sort.Strings(wantStr)

			got := generateTrees(tt.args.n)
			var gotStr = make([]string, len(got))
			for i := range got {
				gotStr[i] = got[i].String()
			}
			sort.Strings(gotStr)

			assert.Equal(gotStr, wantStr)
		})
	}
}