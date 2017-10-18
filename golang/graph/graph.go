package graph

import (
	"sort"

	"github.com/matthewmueller/golly/golang/def"
	"github.com/stevenle/topsort"
)

// Graph struct
type Graph struct {
	nodes       map[string]def.Definition
	edges       map[string]map[string]bool
	modules     map[string]bool
	moduleEdges map[string]map[string]bool
	pathgraph   *topsort.Graph
	defgraphs   map[string]*topsort.Graph
}

// New function
func New() *Graph {
	return &Graph{
		nodes: map[string]def.Definition{},
		// edges:       map[string]map[string]bool{},
		// modules:     map[string]bool{},
		// moduleEdges: map[string]map[string]bool{},
		pathgraph: topsort.NewGraph(),
		// defgraphs:   map[string]*topsort.Graph{},
	}
}

// AddDependency fn
func (g *Graph) AddDependency(parent def.Definition, children ...def.Definition) {
	parentPath := parent.Path()
	g.pathgraph.AddNode(parentPath)
	g.nodes[parent.ID()] = parent

	// build the module dep graph
	for _, child := range children {
		childPath := child.Path()
		if parentPath == childPath {
			continue
		}
		g.pathgraph.AddNode(childPath)
		g.pathgraph.AddEdge(parentPath, childPath)
		g.nodes[child.ID()] = child
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

// Sort declarations topologically based on their package path
func (g *Graph) Sort(d def.Definition) (defs []def.Definition, err error) {
	// topologically sort the module paths
	order, e := g.pathgraph.TopSort(d.Path())
	if e != nil {
		return defs, e
	}

	// sort the definitions within the packages
	// TODO: right now this will sort alphabetically
	// but it's probably better to sort based on order
	// in the original source file
	nodes := []string{}
	for _, n := range g.nodes {
		nodes = append(nodes, n.ID())
	}
	sort.Strings(nodes)

	// group definitions into modules
	defmap := map[string][]def.Definition{}
	for _, n := range nodes {
		node := g.nodes[n]
		path := node.Path()
		if defmap[path] == nil {
			defmap[path] = []def.Definition{}
		}
		defmap[path] = append(defmap[path], node)
	}

	// order the modules
	for _, path := range order {
		defs = append(defs, defmap[path]...)
	}

	return defs, nil
}
