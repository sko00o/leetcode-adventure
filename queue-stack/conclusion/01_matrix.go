package conclusion

// BFS
// Runtime: 236 ms
// Memory Usage: 8.1 MB
func updateMatrix(mat [][]int) [][]int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return mat
	}

	var got = make(map[[2]int]bool)
	for i := range mat {
		for j := range mat[i] {
			mat[i][j] = getClosest(mat, i, j, got)
		}
	}

	return mat
}

func getClosest(mat [][]int, i, j int, got map[[2]int]bool) int {
	if mat[i][j] == 0 {
		got[[2]int{i, j}] = true
		return 0
	}
	var queue [][2]int
	var visited = make(map[[2]int]bool)
	var dis int

	queue = append(queue, [2]int{i, j})
	visited[[2]int{i, j}] = true
	R, C := len(mat), len(mat[0])

	for len(queue) != 0 {
		for k := len(queue); k > 0; k-- {
			curr := queue[0]
			ii, jj := curr[0], curr[1]
			if mat[ii][jj] == 0 {
				got[curr] = true
				return dis
			}
			if got[curr] {
				return dis + mat[ii][jj]
			}

			for _, d := range [][2]int{
				{0, 1},  // right
				{1, 0},  // down
				{0, -1}, // left
				{-1, 0}, // top
			} {
				ii, jj := curr[0]+d[0], curr[1]+d[1]
				if ii < 0 || ii > R-1 ||
					jj < 0 || jj > C-1 {
					continue
				}

				nPos := [2]int{ii, jj}
				if visited[nPos] {
					continue
				}
				visited[nPos] = true
				queue = append(queue, nPos)
			}

			queue = queue[1:]
		}
		dis++
	}

	return -1
}

// BFS
// Runtime: 108 ms
// Memory Usage: 6.9 MB
func updateMatrix1(mat [][]int) [][]int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return mat
	}

	var rows, cols = len(mat), len(mat[0])
	var max = rows + cols
	var queue = make([][2]int, 0, rows*cols)
	var dist = make([][]int, rows)

	// put all 0 in queue
	for i := range mat {
		dist[i] = make([]int, cols)
		for j := range mat[i] {
			if mat[i][j] == 0 {
				dist[i][j] = 0
				queue = append(queue, [2]int{i, j})
			} else {
				dist[i][j] = max
			}
		}
	}

	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]
		for _, d := range [][2]int{
			{0, 1},  // right
			{1, 0},  // down
			{0, -1}, // left
			{-1, 0}, // top
		} {
			nr, nc := curr[0]+d[0], curr[1]+d[1]
			if nr < 0 || nr > rows-1 ||
				nc < 0 || nc > cols-1 {
				continue
			}

			if ndist := dist[curr[0]][curr[1]] + 1; dist[nr][nc] > ndist {
				dist[nr][nc] = ndist
				queue = append(queue, [2]int{nr, nc})
			}
		}
	}

	return dist
}

// DP
// Runtime: 84 ms
// Memory Usage: 6.8 MB
func updateMatrix2(mat [][]int) [][]int {
	rows := len(mat)
	if rows == 0 {
		return mat
	}
	cols := len(mat[0])
	if cols == 0 {
		return mat
	}
	var dist = make([][]int, rows)

	// check for left and top
	for i := range mat {
		dist[i] = make([]int, cols)
		for j := range mat[i] {
			if mat[i][j] == 0 {
				dist[i][j] = 0
			} else {
				dist[i][j] = rows + cols
				if i > 0 && dist[i][j] > dist[i-1][j]+1 {
					dist[i][j] = dist[i-1][j] + 1
				}
				if j > 0 && dist[i][j] > dist[i][j-1]+1 {
					dist[i][j] = dist[i][j-1] + 1
				}
			}
		}
	}

	// check for right and bottom
	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j >= 0; j-- {
			if i < rows-1 && dist[i][j] > dist[i+1][j]+1 {
				dist[i][j] = dist[i+1][j] + 1
			}
			if j < cols-1 && dist[i][j] > dist[i][j+1]+1 {
				dist[i][j] = dist[i][j+1] + 1
			}
		}
	}

	return dist
}
