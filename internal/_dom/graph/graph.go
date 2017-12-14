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

func (g *Graph) Node(node Node) {
	g.nodes[node.ID()] = node
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

// ToposortBy sorts topologically
func (g *Graph) ToposortBy(node Node) (sorted []string) {
	sorted = g.dfs(node, nil)
	return sorted
}

// Toposort sorts topologically
func (g *Graph) Toposort() (sorted []Node, hasCycles bool) {
	reversed := g.reverse()
	var roots []string
	for _, node := range reversed.nodes {
		if len(reversed.edges[node.ID()]) == 0 {
			roots = append(roots, node.ID())
		}
	}

	clone := g.Clone()
	for len(roots) > 0 {
		root := roots[0]
		roots = roots[1:]
		sorted = append(sorted, g.nodes[root])
		for child := range clone.edges[root] {
			delete(clone.edges[root], child)
			if len(clone.edges[root]) == 0 {
				sorted = append(sorted, g.nodes[child])
			}
		}
	}

	for parent := range clone.edges {
		if len(clone.edges[parent]) > 0 {
			return sorted, true
		}
	}

	return sorted, false
}

// Clone the graph
func (g *Graph) Clone() *Graph {
	clone := New()
	for parent, children := range g.edges {
		for child := range children {
			clone.Edge(g.nodes[parent], g.nodes[child])
		}
	}
	return clone
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
	clone := New()

	for parent, children := range g.edges {
		for child := range children {
			clone.Edge(g.nodes[child], g.nodes[parent])
		}
	}

	return clone
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
