package practice

import (
	"github.com/sko00o/leetcode-adventure/queue-stack/queue"
)

// Queue defines a queue for interface{} type.
type Queue struct {
	queue.Queue
}

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
		rd = (rd + (d[i] % 10) + 10) % 10
		curr[i] = '0' + byte(rd)
	}

	return string(curr)
}

func openLock1(deadends []string, target string) int {
	start := "0000"
	visited := make(map[string]bool)
	for _, d := range deadends {
		visited[d] = true
	}
	if visited[start] {
		return -1
	}
	visited[start] = true

	d := []int{1, 0, 0, 0, -1, 0, 0, 0, 1, 0, 0}
	for q, cnt := []string{start}, 0; len(q) != 0; cnt++ {
		nq := []string{}
		for _, curr := range q {
			if curr == target {
				return cnt
			}
			for i := 0; i < 8; i++ {
				next := nextStats(curr, d[i:i+4])
				if visited[next] {
					continue
				}
				visited[next] = true
				nq = append(nq, next)
			}
		}
		q = nq
	}

	return -1
}
