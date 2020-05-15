package tmpl1

import (
	"strconv"
	"testing"

	"github.com/sko00o/leetcode-adventure/queue-stack/stack/dfs"
	"github.com/sko00o/leetcode-adventure/queue-stack/stack/dfs/tmpl1"
	"github.com/stretchr/testify/require"
)

func TestBFS(t *testing.T) {
	E := &dfs.Node{}
	B := &dfs.Node{
		Neighbors: []*dfs.Node{E},
	}
	G := &dfs.Node{}
	F := &dfs.Node{
		Neighbors: []*dfs.Node{G},
	}
	D := &dfs.Node{
		Neighbors: []*dfs.Node{G},
	}
	C := &dfs.Node{
		Neighbors: []*dfs.Node{E, F},
	}
	A := &dfs.Node{
		Neighbors: []*dfs.Node{B, C, D},
	}

	tests := []struct {
		root, target *dfs.Node
		want         bool
	}{
		{A, B, true},
		{A, C, true},
		{A, D, true},
		{A, E, true},
		{A, F, true},
		{A, G, true},
		{C, G, true},
		{B, G, false},
		{D, E, false},
	}

	for idx, DFS := range []func(root, target *dfs.Node) bool{
		tmpl1.DFS,
	} {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			for idx, tst := range tests {
				t.Run(strconv.Itoa(idx), func(t *testing.T) {
					assert := require.New(t)
					assert.Equal(tst.want, DFS(tst.root, tst.target))
				})
			}
		})
	}
}
