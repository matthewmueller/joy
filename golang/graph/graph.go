package graph

import (
	"sort"

	"github.com/matthewmueller/golly/golang/def"
)

// Graph struct
type Graph struct {
	defs  map[string]def.Definition
	nodes map[string]bool
	edges map[string]map[string]string
}

// New function
func New() *Graph {
	return &Graph{
		defs:  map[string]def.Definition{},
		nodes: map[string]bool{},
		edges: map[string]map[string]string{},
	}
}

// Edge adds an edge
func (g *Graph) Edge(parent, child, kind string) {
	if parent == child {
		return
	}

	// TODO: is this legit?
	// maybe the package graph doesn't depend
	// on interface links at all
	if kind == "IMPLEMENTS" {
		return
	}

	g.nodes[parent] = true
	g.nodes[child] = true

	if g.edges[parent] == nil {
		g.edges[parent] = map[string]string{}
	}
	g.edges[parent][child] = kind
}

// Toposort sorts topologically
func (g *Graph) Toposort(path string) (sorted []string) {
	sorted = g.dfs(path, nil)
	// log.Infof("path %s => %+v", path, sorted)
	return sorted
}

func (g *Graph) dfs(path string, visited map[string]bool) (order []string) {
	if visited == nil {
		visited = map[string]bool{}
	}

	visited[path] = true

	children := g.edges[path]
	sortedChildren := sortKeys(children)
	for _, child := range sortedChildren {
		if visited[child] {
			continue
		}
		o := g.dfs(child, visited)
		order = append(order, o...)
	}

	order = append(order, path)
	return order
}

func sortKeys(m map[string]string) (sorted []string) {
	for k := range m {
		sorted = append(sorted, k)
	}
	sort.Strings(sorted)
	return sorted
}
