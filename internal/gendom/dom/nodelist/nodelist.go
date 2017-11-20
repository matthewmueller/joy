package nodelist

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/node"
	"github.com/matthewmueller/golly/js"
)

type NodeList struct {
}

func (*NodeList) Item(index uint) (n *node.Node) {
	js.Rewrite("$<.item($1)", index)
	return n
}

func (*NodeList) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}
