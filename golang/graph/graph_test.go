package graph_test

import (
	"testing"

	"github.com/apex/log"
	"github.com/matthewmueller/golly/golang/graph"
)

func TestTopo(t *testing.T) {
	g := graph.New()
	// g.Edge("B", "D", "DEPENDS")
	// g.Edge("B", "C", "DEPENDS")
	// g.Edge("C", "G", "DEPENDS")
	// g.Edge("B", "E", "DEPENDS")
	// g.Edge("E", "F", "DEPENDS")
	// g.Edge("F", "G", "DEPENDS")
	// g.Edge("F", "B", "DEPENDS")
	g.Edge("B", "A", "DEPENDS")
	g.Edge("A", "B", "DEPENDS")
	sorted := g.Toposort("A")
	log.Infof("sorted %+v", sorted)
}
