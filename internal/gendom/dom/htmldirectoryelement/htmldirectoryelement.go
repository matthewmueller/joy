package htmldirectoryelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLDirectoryElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLDirectoryElement) GetCompact() (compact bool) {
	js.Rewrite("$<.compact")
	return compact
}

func (*HTMLDirectoryElement) SetCompact(compact bool) {
	js.Rewrite("$<.compact = $1", compact)
}
