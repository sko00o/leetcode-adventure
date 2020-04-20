package relation

func generate(numRows int) [][]int {
	tri := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		ll := i + 1
		tri[i] = make([]int, ll)
		tri[i][0] = 1
		tri[i][ll-1] = 1
		for j := 1; j < ll-1; j++ {
			tri[i][j] = tri[i-1][j-1] + tri[i-1][j]
		}
	}
	return tri
}

func getRow0(rowIndex int) []int {
	out := generate(rowIndex + 1)
	return out[rowIndex]
}

func getRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		res[i] = 1
		for j := i - 1; j >= 1; j-- {
			res[j] += res[j-1]
		}
	}
	return res
}
