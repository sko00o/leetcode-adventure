package conclusion

// Runtime: 4 ms
// Memory Usage: 4.5 MB
func canVisitAllRooms(rooms [][]int) bool {
	var visited = make(map[int]bool)
	var queue = make([]int, 0, len(rooms))

	queue = append(queue, 0)
	visited[0] = true
	for len(queue) != 0 {
		for k := len(queue); k > 0; k-- {
			curr := queue[0]
			queue = queue[1:]
			for _, key := range rooms[curr] {
				if visited[key] {
					continue
				}
				visited[key] = true
				queue = append(queue, key)
			}
		}
	}
	return len(visited) == len(rooms)
}
