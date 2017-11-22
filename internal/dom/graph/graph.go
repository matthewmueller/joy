package graph

import (
	"sort"
)

// Graph struct
type Graph struct {
	nodes map[string]Node
	edges map[string]map[string]bool
}

// Node interface
type Node interface {
	ID() string
}

// New function
func New() *Graph {
	return &Graph{
		nodes: map[string]Node{},
		edges: map[string]map[string]bool{},
	}
}

// Edge adds an edge
func (g *Graph) Edge(parent Node, child Node) {
	g.nodes[parent.ID()] = parent

	if parent.ID() == child.ID() {
		return
	}
	g.nodes[child.ID()] = child

	if g.edges[parent.ID()] == nil {
		g.edges[parent.ID()] = map[string]bool{}
	}
	g.edges[parent.ID()][child.ID()] = true
}

// Toposort sorts topologically
func (g *Graph) Toposort(node Node) (sorted []string) {
	sorted = g.dfs(node, nil)
	return sorted
}

// Cliques gets the circular subgraphs
// We're going to group these together in a single
// package, so Go doesn't barf.
func (g *Graph) Cliques() (cliques [][]Node) {
	// step 1: visit each node and built a list
	visited := map[string]bool{}
	list := []string{}

	var visit func(node string)
	visit = func(node string) {
		if visited[node] {
			return
		}
		visited[node] = true

		for child := range g.edges[node] {
			visit(child)
		}
		list = append([]string{node}, list...)
	}

	for node := range g.nodes {
		visit(node)
	}

	visited = map[string]bool{}
	cliquemap := map[string][]string{}
	reversed := g.reverse()

	// step 2: assign each node to a clique
	var assign func(node string, root string)
	assign = func(node string, root string) {
		if visited[node] {
			return
		}
		visited[node] = true

		if root == "" {
			cliquemap[node] = []string{node}
		} else {
			cliquemap[root] = append(cliquemap[root], node)
		}

		for child := range reversed.edges[node] {
			if root == "" {
				assign(child, node)
			} else {
				assign(child, root)
			}
		}
	}

	// go through the stack reversed
	for _, node := range list {
		assign(node, "")
	}

	// turn cliquemap into cliques
	for _, clique := range cliquemap {
		nodes := []Node{}
		for _, node := range clique {
			nodes = append(nodes, g.nodes[node])
		}
		cliques = append(cliques, nodes)
	}

	return cliques
}

// reverse the edge direction of our graph
func (g *Graph) reverse() *Graph {
	edges := map[string]map[string]bool{}
	for parent, children := range g.edges {
		for child := range children {
			if edges[child] == nil {
				edges[child] = map[string]bool{}
			}
			edges[child][parent] = true
		}
	}

	// for parent, children := range g.edges {
	// 	for child := range children {
	// 		log.Infof("edges %s -> %s", parent, child)
	// 	}
	// }
	// log.Infof("reversed")
	// for parent, children := range edges {
	// 	for child := range children {
	// 		log.Infof("edges %s -> %s", parent, child)
	// 	}
	// }

	return &Graph{
		nodes: g.nodes,
		edges: edges,
	}
}

func (g *Graph) dfs(node Node, visited map[string]bool) (order []string) {
	if visited == nil {
		visited = map[string]bool{}
	}

	visited[node.ID()] = true

	children := g.edges[node.ID()]
	sortedChildren := sortKeys(children)
	for _, child := range sortedChildren {
		if visited[child] {
			continue
		}
		o := g.dfs(g.nodes[child], visited)
		order = append(order, o...)
	}

	order = append(order, node.ID())
	return order
}

func (g *Graph) unique(ids []string) []string {
	l := len(ids)

	order := make([]string, 0, l)
	seen := make(map[string]bool)

	// get the unique ids in reverse order
	// to maintain the topology
	for i := 0; i < l; i++ {
		id := ids[i]

		if _, ok := seen[id]; !ok {
			seen[id] = true
			order = append(order, id)
		}
	}

	return order
}

func (g *Graph) reverseUnique(ids []string) []string {
	l := len(ids)

	order := make([]string, 0, l)
	seen := make(map[string]bool)

	// get the unique ids in reverse order
	// to maintain the topology
	for i := l - 1; i >= 0; i-- {
		id := ids[i]

		if _, ok := seen[id]; !ok {
			seen[id] = true
			order = append(order, id)
		}
	}

	return order
}

func sortKeys(m map[string]bool) (sorted []string) {
	for k := range m {
		sorted = append(sorted, k)
	}
	sort.Strings(sorted)
	return sorted
}
