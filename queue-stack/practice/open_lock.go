package practice

func openLock(deadends []string, target string) int {
	if len(target) != 4 {
		return -1
	}

	var queue Queue
	var step int
	visited := make(map[string]struct{})
	for _, d := range deadends {
		if len(d) == 4 {
			visited[d] = struct{}{}
		}
	}

	// init
	start := "0000"
	if _, ok := visited[start]; ok {
		return -1
	}
	visited[start] = struct{}{}
	queue.EnQueue(start)

	d := []int{1, 0, 0, 0, -1, 0, 0, 0, 1, 0, 0}
	for !queue.IsEmpty() {
		for k := len(queue.Data); k > 0; k-- {
			curr := queue.Front().(string)
			if curr == target {
				return step
			}

			for i := 0; i < 8; i++ {
				next := nextStats(curr, d[i:i+4])

				if _, ok := visited[next]; ok {
					continue
				}
				visited[next] = struct{}{}
				queue.EnQueue(next)
			}

			queue.DeQueue()
		}

		step++
	}

	return -1
}

func nextStats(s string, d []int) string {
	curr := []byte(s)
	if len(curr) != 4 || len(d) != 4 {
		return ""
	}

	for i := 0; i < 4; i++ {
		rd := int(curr[i] - '0')
		rd = (rd + d[i]) % 10
		if rd < 0 {
			rd += 10
		}
		curr[i] = '0' + byte(rd)
	}

	return string(curr)
}
