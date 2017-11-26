package graph_test

import (
	"testing"

	"github.com/matthewmueller/golly/internal/dom/graph"
	"github.com/stretchr/testify/assert"
)

var _ graph.Node = (*testNode)(nil)

type testNode struct {
	id string
}

func (t *testNode) ID() string {
	return t.id
}

func node(id string) *testNode {
	return &testNode{id}
}

func TestCliqueOne(t *testing.T) {
	a := node("A")
	b := node("B")
	c := node("C")
	d := node("D")
	e := node("E")
	f := node("F")
	g := node("G")
	h := node("H")
	i := node("I")
	j := node("J")
	k := node("K")

	gr := graph.New()
	gr.Edge(a, b)
	gr.Edge(b, c)
	gr.Edge(c, a)
	gr.Edge(b, d)
	gr.Edge(d, e)
	gr.Edge(e, f)
	gr.Edge(f, d)
	gr.Edge(g, f)
	gr.Edge(g, h)
	gr.Edge(h, i)
	gr.Edge(i, j)
	gr.Edge(j, g)
	gr.Edge(j, k)

	// for i, clique := range gr.Cliques() {
	// 	log.Infof("%dth clique", i)
	// 	for _, node := range clique {
	// 		log.Infof("node %s", node.ID())
	// 	}
	// }
	assert.Equal(t, 4, len(gr.Cliques()))
	// assert.Equal(t, 3, len(gr.Cliques()[0]))
}

// func TestCliqueTwo(t *testing.T) {
// 	a := node("A")
// 	b := node("B")
// 	c := node("C")
// 	d := node("D")

// 	g := graph.New()
// 	g.Edge(a, b)
// 	g.Edge(b, c)
// 	g.Edge(c, a)
// 	g.Edge(c, d)

// 	assert.Equal(t, 2, len(g.Cliques()))
// 	assert.Equal(t, 3, len(g.Cliques()[0]))
// 	assert.Equal(t, 1, len(g.Cliques()[1]))
// }
