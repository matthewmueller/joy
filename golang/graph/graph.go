package graph

import (
	"sort"

	"github.com/matthewmueller/golly/golang/def"
)

// Graph struct
type Graph struct {
	nodes map[string]def.Definition
	edges map[string]map[string]bool
}

// New function
func New() *Graph {
	return &Graph{
		nodes: map[string]def.Definition{},
		edges: map[string]map[string]bool{},
	}
}

// AddDependency fn
func (g *Graph) AddDependency(parent def.Definition, children ...def.Definition) {
	pid := parent.ID()

	// add the parent if not already present
	g.nodes[pid] = parent

	for _, child := range children {
		cid := child.ID()

		g.nodes[cid] = child

		if g.edges[pid] == nil {
			g.edges[pid] = map[string]bool{}
		}
		g.edges[pid][cid] = true
	}
}

// Roots gets a list of root declarations (those without any dependants)
func (g *Graph) Roots() (decls []def.Definition) {
	for id, decl := range g.nodes {
		if _, isset := g.edges[id]; !isset {
			decls = append(decls, decl)
		}
	}
	return decls
}

// Sort declarations topologically
func (g *Graph) Sort(d def.Definition) (defs []def.Definition) {
	var visitAll func(ids map[string]bool)
	seen := map[string]bool{}
	order := []string{}

	visitAll = func(ids map[string]bool) {
		var sorted []string
		for id := range ids {
			sorted = append(sorted, id)
		}
		sort.Strings(sorted)

		for _, id := range sorted {
			if !seen[id] {
				seen[id] = true
				visitAll(g.edges[id])
				order = append(order, id)
			}
		}
	}

	visitAll(g.edges[d.ID()])

	for _, id := range order {
		defs = append(defs, g.nodes[id])
	}

	return defs
}
