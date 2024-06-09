package main

type Node struct {
	value    string
	children []*Node
}

type Item struct {
	node  *Node
	depth int
}

func CheckDuplicateIDs(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []Item{{node: root, depth: 0}}
	nodeMap := make(map[string]int)

	for len(queue) > 0 {
		currentItem := queue[0]
		queue = queue[1:]

		node, depth := currentItem.node, currentItem.depth

		if _, exists := nodeMap[node.value]; exists {
			return node
		} else {
			nodeMap[node.value] = depth
		}

		for _, childNode := range node.children {
			queue = append(queue, Item{node: childNode, depth: depth + 1})
		}
	}

	return nil
}