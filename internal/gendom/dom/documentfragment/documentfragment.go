package documentfragment

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/node"
	"github.com/matthewmueller/golly/internal/gendom/dom/nodeselector"
)

type DocumentFragment struct {
	*node.Node
	*nodeselector.NodeSelector
}
