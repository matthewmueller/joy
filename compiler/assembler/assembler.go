package assembler

import (
	"github.com/apex/log"
	"github.com/matthewmueller/golly/compiler/file"
	"github.com/matthewmueller/golly/compiler/graph"
)

// Assemble the graph into files
func Assemble(g *graph.Graph) (files []*file.File, err error) {
	// get the roots of the graph, these will be our files
	for _, node := range g.Roots() {
		sorted := node.Sort()
		log.Infof("sorted! %+v", sorted)
	}

	return files, nil
}
