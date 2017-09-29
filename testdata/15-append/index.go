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

	println(nodes[0].NodeName)
	println(nodes[1].NodeName)
}
