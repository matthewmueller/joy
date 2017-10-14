package parser

import (
	"github.com/matthewmueller/golly/compiler/decl"
	"github.com/matthewmueller/golly/compiler/graph"
	"github.com/matthewmueller/golly/compiler/indexer"
	"github.com/pkg/errors"
)

// Parse the packages into a dependency graph (Phase III)
func Parse(idx *indexer.Index) (g *graph.Graph, err error) {
	queue := []decl.Declaration{}
	visited := map[string]bool{}
	g = graph.New()

	mains, err := idx.Mains()
	if err != nil {
		return g, err
	}
	queue = append(queue, mains...)

	for len(queue) > 0 {
		// dequeue
		decl := queue[0]
		queue = queue[1:]

		// get the dependencies
		deps, err := decl.Dependencies()
		if err != nil {
			return g, errors.Wrap(err, "error getting dependencies")
		}

		// add the dependencies to the graph
		g.AddDependency(decl, deps...)

		// add any dependency that hasn't
		// already been visited
		for _, dep := range deps {
			if !visited[dep.ID()] {
				queue = append(queue, dep)
			}
		}
	}

	return g, nil
}
