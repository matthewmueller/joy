package graph

import (
	"github.com/matthewmueller/golly/compiler/types"
)

// Graph struct
type Graph struct {
	declarations map[string]types.Declaration
	edges        map[string][]string
}

// New function
func New() *Graph {
	return &Graph{
		declarations: map[string]types.Declaration{},
		edges:        map[string][]string{},
	}
}

// AddDependency fn
func (g *Graph) AddDependency(parent types.Declaration, children ...types.Declaration) {
	pid := parent.ID()

	// add the parent if not already present
	g.declarations[pid] = parent

	for _, child := range children {
		cid := child.ID()

		g.declarations[cid] = child

		if g.edges[pid] == nil {
			g.edges[pid] = []string{}
		}
		g.edges[pid] = append(g.edges[pid], cid)
	}
}

// Roots gets a list of root declarations (those without any dependants)
func (g *Graph) Roots() (decls []types.Declaration) {
	for id, decl := range g.declarations {
		if _, isset := g.edges[id]; !isset {
			decls = append(decls, decl)
		}
	}

	return decls
}

// Sort declarations topologically
func (g *Graph) Sort(d types.Declaration) (decls []types.Declaration) {
	var visitAll func(ids []string)
	seen := map[string]bool{}
	order := []string{}

	visitAll = func(ids []string) {
		for _, id := range ids {
			if !seen[id] {
				seen[id] = true
				visitAll(g.edges[id])
				order = append(order, id)
			}
		}
	}

	visitAll(g.edges[d.ID()])

	for _, id := range order {
		decls = append(decls, g.declarations[id])
	}
	decls = append(decls, d)

	return decls
}

// // Node struct
// type Node struct {
// 	declaration types.Declaration
// }

// // ID fn
// func (n *Node) ID() string {
// 	return n.declaration.ID()
// }

// // Declaration fn
// func (n *Node) Declaration() types.Declaration {
// 	return n.declaration
// }

// // Dependencies fn
// func (n *Node) Dependencies() (deps []Node) {
// 	return deps
// }

// // Sort topologically
// func (n *Node) Sort() (list []Node) {
// 	return list
// }
