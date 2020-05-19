package problems

// BFS
// Runtime: 56 ms
// Memory Usage: 6.5 MB
func canReach(arr []int, start int) bool {
	var visited = make(map[int]bool)
	var queue = make([]int, 0, len(arr))

	queue = append(queue, start)
	visited[start] = true

	for len(queue) != 0 {
		curr := queue[0]
		if arr[curr] == 0 {
			return true
		}

		queue = queue[1:]
		for _, next := range []int{
			curr + arr[curr],
			curr - arr[curr],
		} {
			if next < 0 || next > len(arr)-1 {
				continue
			}
			if visited[next] {
				continue
			}
			visited[next] = true
			queue = append(queue, next)
		}
	}

	return false
}

// DFS
// Runtime: 56 ms
// Memory Usage: 6.7 MB
func canReach1(arr []int, start int) bool {
	return dfs(arr, start, make(map[int]bool))
}

func dfs(arr []int, curr int, visited map[int]bool) bool {
	if curr < 0 || curr > len(arr)-1 {
		return false
	}
	if visited[curr] {
		return false
	}
	visited[curr] = true
	return arr[curr] == 0 ||
		dfs(arr, curr+arr[curr], visited) ||
		dfs(arr, curr-arr[curr], visited)
}

// DFS
// Runtime: 60 ms
// Memory Usage: 6.7 MB
func canReach2(arr []int, curr int) bool {
	return dfs1(arr, curr)
}

func dfs1(arr []int, curr int) bool {
	if curr < 0 || curr > len(arr)-1 || arr[curr] < 0 {
		return false
	}

	arr[curr] = -arr[curr]
	return arr[curr] == 0 ||
		dfs1(arr, curr+arr[curr]) ||
		dfs1(arr, curr-arr[curr])
}
