package defs

import (
	"github.com/matthewmueller/golly/golang/def"
)

type edge struct {
	kind string
	def  def.Definition
}

func (e *edge) Type() string {
	return e.kind
}

func (e *edge) Definition() def.Definition {
	return e.def
}

type edges struct {
	edges []*edge
}

// Add if we don't already have the edge
func (e *edges) Add(kind string, d def.Definition) {
	for _, edge := range e.edges {
		if d.ID() == edge.def.ID() {
			return
		}
	}

	e.edges = append(e.edges, &edge{
		kind: kind,
		def:  d,
	})
}

func (e *edges) Edges() (edges []def.Edge) {
	for _, edge := range e.edges {
		edges = append(edges, edge)
	}
	return edges
}
