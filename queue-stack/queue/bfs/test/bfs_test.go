package tmpl1

import (
	"strconv"
	"testing"

	"github.com/sko00o/leetcode-adventure/queue-stack/queue/bfs"
	"github.com/sko00o/leetcode-adventure/queue-stack/queue/bfs/tmpl1"
	"github.com/sko00o/leetcode-adventure/queue-stack/queue/bfs/tmpl2"
	"github.com/stretchr/testify/require"
)

func TestBFS(t *testing.T) {
	E := &bfs.Node{}
	B := &bfs.Node{
		Neighbors: []*bfs.Node{E},
	}
	G := &bfs.Node{}
	F := &bfs.Node{
		Neighbors: []*bfs.Node{G},
	}
	D := &bfs.Node{
		Neighbors: []*bfs.Node{G},
	}
	C := &bfs.Node{
		Neighbors: []*bfs.Node{E, F},
	}
	A := &bfs.Node{
		Neighbors: []*bfs.Node{B, C, D},
	}

	tests := []struct {
		root, target *bfs.Node
		step         int
	}{
		{A, B, 1},
		{A, C, 1},
		{A, D, 1},
		{A, E, 2},
		{A, F, 2},
		{A, G, 2},
		{C, G, 2},
		{B, G, -1},
		{D, E, -1},
	}

	for idx, BFS := range []func(root, target *bfs.Node) int{
		tmpl1.BFS,
		tmpl2.BFS,
	} {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			for idx, tst := range tests {
				t.Run(strconv.Itoa(idx), func(t *testing.T) {
					assert := require.New(t)
					assert.Equal(tst.step, BFS(tst.root, tst.target))
				})
			}
		})
	}
}
