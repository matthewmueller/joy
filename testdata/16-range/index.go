package main

// Node struct
type Node struct {
	NodeName string
}

func main() {
	var nodes []Node

	nodes = append(nodes, Node{
		NodeName: "div",
	}, Node{
		NodeName: "strong",
	})

	for i, node := range nodes {
		println(i)
		println(node.NodeName)
	}

	for i := range nodes {
		println(i)
	}
}
