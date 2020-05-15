package practice

// Node defines a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	return cloneGraphDFS(node, make(map[*Node]*Node))
}

func cloneGraphDFS(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}

	if nd, ok := visited[node]; ok {
		return nd
	}
	n := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, 0, len(node.Neighbors)),
	}
	visited[node] = n
	for i := range node.Neighbors {
		n.Neighbors = append(n.Neighbors, cloneGraphDFS(node.Neighbors[i], visited))
	}

	return n
}

func newGraph(al [][]int) *Node {
	nodeSet := make(map[int]*Node)
	for i := range al {
		val := i + 1
		node, ok := nodeSet[val]
		if !ok {
			node = &Node{Val: val}
			nodeSet[val] = node
		}
		for j := range al[i] {
			val := al[i][j]
			nei, ok := nodeSet[val]
			if !ok {
				nei = &Node{Val: val}
				nodeSet[val] = nei
			}
			node.Neighbors = append(node.Neighbors, nei)
		}
	}

	val, ok := nodeSet[1]
	if !ok {
		return nil
	}
	return val
}

func isCopyed(ori, gen *Node) bool {
	return isCopyedDFS(ori, gen, make(map[*Node]bool))
}

func isCopyedDFS(ori, gen *Node, visited map[*Node]bool) bool {
	if ori == nil && gen == nil {
		return true
	}
	if ori == gen || ori == nil || gen == nil {
		return false
	}
	if visited[ori] != visited[gen] {
		return false
	}
	visited[ori] = true
	visited[gen] = true
	if ori.Val != gen.Val {
		return false
	}
	if len(ori.Neighbors) != len(gen.Neighbors) {
		return false
	}
	for i := len(ori.Neighbors) - 1; i >= 0; i-- {
		nOri, nGen := ori.Neighbors[i], gen.Neighbors[i]
		if !visited[nOri] && !visited[nGen] {
			if isCopyedDFS(nOri, nGen, visited) == false {
				return false
			}
		}
	}
	return true
}
