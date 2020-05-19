package conclusion

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	// copy image
	var img = make([][]int, len(image))
	for i := range image {
		img[i] = make([]int, len(image[i]))
		for j := range image[i] {
			img[i][j] = image[i][j]
		}
	}

	fill(img, sr, sc, img[sr][sc], newColor)
	return img
}

func fill(img [][]int, i, j int, oc, nc int) {
	if oc == nc {
		return
	}
	if isInvalid(img, i, j, oc) {
		return
	}

	img[i][j] = nc
	for _, d := range []struct{ i, j int }{
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
		{-1, 0}, // top
	} {
		fill(img, i+d.i, j+d.j, oc, nc)
	}
}

func isInvalid(img [][]int, i, j int, c int) bool {
	return i < 0 ||
		i >= len(img) ||
		j < 0 ||
		j >= len(img[0]) ||
		img[i][j] != c
}

func floodFill1(img [][]int, sr int, sc int, newColor int) [][]int {
	oldColor := img[sr][sc]
	if oldColor == newColor {
		return img
	}
	R, C := len(img), len(img[0])
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if img[r][c] == oldColor {
			img[r][c] = newColor
			if r > 0 {
				dfs(r-1, c)
			}
			if r < R-1 {
				dfs(r+1, c)
			}
			if c > 0 {
				dfs(r, c-1)
			}
			if c < C-1 {
				dfs(r, c+1)
			}
		}
	}
	dfs(sr, sc)
	return img
}
