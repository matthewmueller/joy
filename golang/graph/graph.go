package graph

import (
	"sort"

	"github.com/apex/log"
	"github.com/matthewmueller/golly/golang/def"
	"github.com/pkg/errors"
	"github.com/stevenle/topsort"
)

// Graph struct
type Graph struct {
	defs  map[string]def.Definition
	nodes map[string]string
	edges map[string]map[string]bool
}

// New function
func New() *Graph {
	return &Graph{
		// id => definition
		defs: map[string]def.Definition{},
		// path => id
		nodes: map[string]string{},
		// parent path => map[child path]bool
		edges: map[string]map[string]bool{},
	}
}

// AddDependency fn
func (g *Graph) AddDependency(parent def.Definition, children ...def.Definition) {
	parentPath := parent.Path()
	parentID := parent.ID()

	// add the parent
	g.nodes[parentPath] = parentID
	g.defs[parentID] = parent

	// build the module dep graph
	for _, child := range children {
		childPath := child.Path()
		childID := child.ID()

		if parentPath == childPath {
			continue
		}

		g.nodes[childPath] = childID
		g.defs[childID] = child

		if g.edges[parentPath] == nil {
			g.edges[parentPath] = map[string]bool{}
		}
		g.edges[parentPath][childPath] = true
	}
}

// Sort declarations topologically based on their package path
// TODO: this could be a lot better
func (g *Graph) Sort(d def.Definition) (defs []def.Definition, err error) {
	graph := topsort.NewGraph()

	// add the nodes in a sorted order
	nodes := sortNodeKeys(g.nodes)
	for _, node := range nodes {
		graph.AddNode(node)
	}

	for parentPath, edge := range g.edges {
		children := sortEdgeKeys(edge)
		for _, childPath := range children {
			log.Debugf("%s -> %s", parentPath, childPath)
			graph.AddEdge(parentPath, childPath)
		}
	}

	order, e := graph.TopSort(d.Path())
	if e != nil {
		return defs, errors.Wrap(e, "graph/toposort")
	}

	log.Debugf("main %s", d.ID())
	for _, o := range order {
		log.Debugf("order: %s", o)
	}

	// group definitions into modules
	buckets := map[string][]def.Definition{}
	ids := sortDefKeys(g.defs)
	for _, id := range ids {
		node := g.defs[id]
		path := node.Path()
		if buckets[path] == nil {
			buckets[path] = []def.Definition{}
		}
		buckets[path] = append(buckets[path], node)
	}

	// order the modules
	for _, path := range order {
		for _, def := range buckets[path] {
			defs = append(defs, def)
		}
	}

	return defs, nil
}

func sortNodeKeys(m map[string]string) (sorted []string) {
	for k := range m {
		sorted = append(sorted, k)
	}
	sort.Strings(sorted)
	return sorted
}

func sortEdgeKeys(m map[string]bool) (sorted []string) {
	for k := range m {
		sorted = append(sorted, k)
	}
	sort.Strings(sorted)
	return sorted
}

func sortDefKeys(m map[string]def.Definition) (sorted []string) {
	for k := range m {
		sorted = append(sorted, k)
	}
	sort.Strings(sorted)
	return sorted
}
