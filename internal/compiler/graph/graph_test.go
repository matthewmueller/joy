package graph_test

import (
	"testing"

	"github.com/apex/log"
	"github.com/matthewmueller/golly/internal/compiler/graph"
)

var _ graph.Node = (*testNode)(nil)

type testNode struct {
	id   string
	path string
}

func (t *testNode) ID() string {
	return t.id
}

func (t *testNode) Path() string {
	return t.path
}

func node(id, path string) graph.Node {
	return &testNode{
		id:   id,
		path: path,
	}
}

func TestTopo(t *testing.T) {
	g := graph.New()
	am := node("am", "main")
	bm := node("bm", "main")
	cm := node("cm", "main")
	a1 := node("a1", "1")
	b1 := node("b1", "1")
	c1 := node("c1", "1")
	a2 := node("a2", "2")
	b2 := node("b2", "2")
	c2 := node("c2", "2")

	a3 := node("a3", "3")
	a4 := node("a4", "4")

	g.Edge(am, bm)
	g.Edge(am, cm)
	g.Edge(cm, a1)

	g.Edge(a1, a1)
	g.Edge(a1, b1)
	g.Edge(b1, c1)

	// disconnected
	g.Edge(c1, a2)
	g.Edge(a2, b2)
	g.Edge(b2, c2)

	g.Edge(c2, a4)
	g.Edge(a3, a4)

	g.Edge(am, a3)

	sorted := g.Toposort(am)
	log.Infof("sorted %+v", sorted)
}
