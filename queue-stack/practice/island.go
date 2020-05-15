package practice

func numIslands(grid [][]byte) int {
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

func isInvalid(grid [][]byte, i, j int) bool {
	return i < 0 ||
		i >= len(grid) ||
		j < 0 ||
		j >= len(grid[0]) ||
		grid[i][j] == '0'
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

// BFS
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

func numIslands2(grid [][]byte) int {
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
