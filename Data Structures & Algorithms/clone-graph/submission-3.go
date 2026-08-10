func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	oldToNew := make(map[*Node]*Node)

	oldToNew[node] = &Node{
		Val: node.Val,
	}

	queue := []*Node{node}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, nei := range curr.Neighbors {
			if _, ok := oldToNew[nei]; !ok {
				oldToNew[nei] = &Node{
					Val: nei.Val,
				}
				queue = append(queue, nei)
			}

			oldToNew[curr].Neighbors = append(
				oldToNew[curr].Neighbors,
				oldToNew[nei],
			)
		}
	}

	return oldToNew[node]
}