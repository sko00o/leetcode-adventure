package practice

import (
	"github.com/sko00o/leetcode-adventure/queue-stack/queue"
)

// Queue defines a queue for interface{} type.
type Queue struct {
	queue.Queue
}

func isInvalid(grid [][]byte, i, j int) bool {
	return i < 0 ||
		i >= len(grid) ||
		j < 0 ||
		j >= len(grid[0]) ||
		grid[i][j] == '0'
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

// DFS
func distroyIsland(grid [][]byte, i, j int) int {
	if isInvalid(grid, i, j) {
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

func numIslands2(grid [][]byte) int {
	visited := make([][]bool, len(grid))
	for i := range grid {
		visited[i] = make([]bool, len(grid[i]))
	}

	// store the num of islands
	var cnt int
	var queue [][2]int
	li := len(grid)
	for i := li - 1; i >= 0; i-- {
		lj := len(grid[i])
		for j := lj - 1; j >= 0; j-- {
			curr := [2]int{i, j}
			if visited[i][j] {
				continue
			}
			visited[i][j] = true

			// this is an island
			if grid[curr[0]][curr[1]] == '1' {
				queue = append(queue, curr)
				cnt++

				// visit an island in BFS
				for len(queue) != 0 {
					curr := queue[0]
					for _, d := range [][2]int{
						{0, 1},  // right
						{1, 0},  // down
						{0, -1}, // left
						{-1, 0}, // top
					} {
						ii, jj := curr[0]+d[0], curr[1]+d[1]
						if ii < li && jj < lj &&
							ii >= 0 && jj >= 0 {
							curr := [2]int{ii, jj}
							if visited[ii][jj] {
								continue
							}
							visited[ii][jj] = true

							if grid[ii][jj] == '1' {
								queue = append(queue, curr)
							}
						}
					}

					queue = queue[1:]
				}
			}
		}
	}

	return cnt
}

// destroy version
func numIslands3(grid [][]byte) int {
	var cnt int
	li := len(grid)
	for i := li - 1; i >= 0; i-- {
		lj := len(grid[i])
		for j := lj - 1; j >= 0; j-- {
			cnt += distroyIslandBFS(grid, i, j)
		}
	}

	return cnt
}

func distroyIslandBFS(grid [][]byte, i, j int) int {
	if isInvalid(grid, i, j) {
		return 0
	}
	if grid[i][j] == '0' {
		return 0
	}

	var queue [][2]int
	grid[i][j] = '0'
	queue = append(queue, [2]int{i, j})

	// visit an island in BFS
	for len(queue) != 0 {
		curr := queue[0]
		for _, d := range [][2]int{
			{0, 1},  // right
			{1, 0},  // down
			{0, -1}, // left
			{-1, 0}, // top
		} {
			ii, jj := curr[0]+d[0], curr[1]+d[1]
			if isInvalid(grid, ii, jj) {
				continue
			}
			grid[ii][jj] = '0'
			queue = append(queue, [2]int{ii, jj})
		}

		queue = queue[1:]
	}
	return 1
}
