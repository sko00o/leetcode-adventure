package search

type SearchAble interface {
	Len() int
	Val(int) int
}

type SearchCol struct {
	Matrix [][]int
	Col    int
}

func (s SearchCol) Len() int {
	return len(s.Matrix)
}
func (s SearchCol) Val(i int) int {
	return s.Matrix[i][s.Col]
}

type SearchRow struct {
	Matrix [][]int
	Row    int
}

func (s SearchRow) Len() int {
	return len(s.Matrix[s.Row])
}
func (s SearchRow) Val(i int) int {
	return s.Matrix[s.Row][i]
}

func search(s SearchAble, target int) (idx int, exist bool) {
	l, h := 0, s.Len()-1
	var m int
	for l <= h {
		m = l + ((h - l) >> 1)
		mid := s.Val(m)
		if mid == target {
			return m, true
		} else if mid > target {
			h = m - 1
		} else {
			l = m + 1
		}
	}
	if s.Val(m) < target {
		return m + 1 , false
	}
	return m, false
}
