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
	Path() string
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
	sorted = g.reverseUniqueByPath(sorted)
	// log.Infof("path %s => %+v", path, sorted)
	return sorted
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

func (g *Graph) reverseUniqueByPath(ids []string) []string {
	l := len(ids)

	buckets := map[string][]string{}
	order := make([]string, 0, l)
	seen := make(map[string]bool)

	// get the unique paths in reverse order
	// to maintain the topology
	for i := l - 1; i >= 0; i-- {
		id := ids[i]
		node := g.nodes[id]
		path := node.Path()

		if buckets[path] == nil {
			buckets[path] = []string{}
		}
		buckets[path] = append(buckets[path], id)

		if _, ok := seen[path]; !ok {
			seen[path] = true
			order = append(order, path)
		}
	}

	ids = []string{}
	// order of paths now needs to be reversed
	// from the previous operation
	for i := len(order) - 1; i >= 0; i-- {
		path := order[i]
		bucket := buckets[path]
		// the bucket order also needs to be reversed
		// to maintain the ids within the path's topology
		for j := len(bucket) - 1; j >= 0; j-- {
			ids = append(ids, bucket[j])
		}
	}

	return ids
}

func sortKeys(m map[string]bool) (sorted []string) {
	for k := range m {
		sorted = append(sorted, k)
	}
	sort.Strings(sorted)
	return sorted
}
