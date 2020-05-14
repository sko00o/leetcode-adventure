package practice

func numSquares(n int) int {
	var queue []int
	var step int
	visited := make(map[int]bool)

	for i := 1; i*i <= n; i++ {
		visited[i*i] = true
		queue = append(queue, i*i)
	}

	for len(queue) != 0 {
		step++

		for k := len(queue); k > 0; k-- {
			if queue[0] == n {
				return step
			}

			for i := 1; queue[0]+i*i <= n; i++ {
				ii := queue[0] + i*i
				if visited[ii] {
					continue
				}
				visited[ii] = true
				queue = append(queue, ii)
			}

			queue = queue[1:]
		}
	}

	return -1
}
