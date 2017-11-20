package htmltableheadercellelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmltablecellelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLTableHeaderCellElement struct {
	*htmltablecellelement.HTMLTableCellElement
}

func (*HTMLTableHeaderCellElement) GetScope() (scope string) {
	js.Rewrite("$<.scope")
	return scope
}

func (*HTMLTableHeaderCellElement) SetScope(scope string) {
	js.Rewrite("$<.scope = $1", scope)
}
