package xmlserializer

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/node"
	"github.com/matthewmueller/golly/js"
)

type XMLSerializer struct {
}

func (*XMLSerializer) SerializeToString(target *node.Node) (s string) {
	js.Rewrite("$<.serializeToString($1)", target)
	return s
}
