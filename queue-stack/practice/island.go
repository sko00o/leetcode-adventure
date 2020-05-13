package practice

import (
	"github.com/sko00o/leetcode-adventure/queue-stack/queue"
)

// Queue defines a queue for interface{} type.
type Queue struct {
	queue.Queue
}

type pos [2]int

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	if len(grid[0]) == 0 {
		return 0
	}

	var queue Queue
	visited := make(map[pos]struct{})
	li, lj := len(grid), len(grid[0])

	// store the num of islands
	var cnt int

	for i := 0; i < li; i++ {
		for j := 0; j < lj; j++ {
			curr := pos{i, j}
			if _, ok := visited[curr]; ok {
				continue
			}
			visited[curr] = struct{}{}

			// this is an island
			if grid[curr[0]][curr[1]] == '1' {
				queue.EnQueue(curr)
				cnt++

				// visit an island in BFS
				for !queue.IsEmpty() {
					curr := queue.Front().(pos)
					for _, d := range []pos{
						{0, 1},  // right
						{1, 0},  // down
						{0, -1}, // left
						{-1, 0}, // top
					} {
						ii, jj := curr[0]+d[0], curr[1]+d[1]
						if ii < li && jj < lj &&
							ii >= 0 && jj >= 0 {
							curr := pos{ii, jj}
							if _, ok := visited[curr]; ok {
								continue
							}
							visited[curr] = struct{}{}

							if grid[ii][jj] == '1' {
								queue.EnQueue(curr)
							}
						}
					}

					queue.DeQueue()
				}
			}
		}
	}

	return cnt
}

func numIslands1(grid [][]byte) int {
	var cnt int
	for i := range grid {
		for j := range grid[0] {
			cnt += distroyIsland(grid, i, j)
		}
	}
	return cnt
}

func distroyIsland(grid [][]byte, i, j int) int {
	var li, lj int
	if li = len(grid); li == 0 {
		return 0
	}
	if lj = len(grid[0]); lj == 0 {
		return 0
	}
	if i < 0 || i >= li || j < 0 || j >= lj || grid[i][j] == '0' {
		return 0
	}
	grid[i][j] = '0'
	for _, d := range []struct{ i, j int }{
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
		{-1, 0}, // top
	} {
		distroyIsland(grid, i+d.i, j+d.j)
	}
	return 1
}
