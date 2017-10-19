package graph

import (
	"errors"
	"path"
	"sort"

	"github.com/apex/log"
	"github.com/matthewmueller/golly/golang/def"
	toposort "github.com/philopon/go-toposort"
)

// Graph struct
type Graph struct {
	defs  map[string]def.Definition
	nodes map[string]string
	edges map[string][]string
}

// New function
func New() *Graph {
	return &Graph{
		defs:  map[string]def.Definition{},
		nodes: map[string]string{},
		edges: map[string][]string{},
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
			g.edges[parentPath] = []string{}
		} else if !contains(g.edges[parentPath], childPath) {
			g.edges[parentPath] = append(g.edges[parentPath], childPath)
		}
	}
}

// Sort declarations topologically based on their package path
// TODO: this is very inefficient
func (g *Graph) Sort(d def.Definition) (defs []def.Definition, err error) {
	topo := toposort.NewGraph(len(g.nodes))
	for path := range g.nodes {
		if ok := topo.AddNode(path); !ok {
			return defs, errors.New("Sort: unable to add node")
		}
	}

	for parentPath, edge := range g.edges {
		sort.Strings(edge)
		for _, childPath := range edge {
			log.Debugf("%s -> %s", path.Base(parentPath), path.Base(childPath))
			if ok := topo.AddEdge(parentPath, childPath); !ok {
				return defs, errors.New("Sort: unable to add edge")
			}
		}
	}

	order, ok := topo.Toposort()
	if !ok {
		return defs, errors.New("Sort: cyclical dependency")
	}
	order = reverse(order)
	for _, o := range order {
		log.Debugf("order: %s", o)
	}

	// sort the definitions within the packages
	// TODO: right now this will sort alphabetically
	// but it's probably better to sort based on order
	// in the original source file
	nodes := []string{}
	for _, n := range g.defs {
		nodes = append(nodes, n.ID())
	}
	sort.Strings(nodes)

	// group definitions into modules
	defmap := map[string][]def.Definition{}
	for _, n := range nodes {
		node := g.defs[n]
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

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func contains(haystack []string, needle string) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}
