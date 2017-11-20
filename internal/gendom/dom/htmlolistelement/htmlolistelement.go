package htmlolistelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLOListElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLOListElement) GetCompact() (compact bool) {
	js.Rewrite("$<.compact")
	return compact
}

func (*HTMLOListElement) SetCompact(compact bool) {
	js.Rewrite("$<.compact = $1", compact)
}

func (*HTMLOListElement) GetStart() (start int) {
	js.Rewrite("$<.start")
	return start
}

func (*HTMLOListElement) SetStart(start int) {
	js.Rewrite("$<.start = $1", start)
}

func (*HTMLOListElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*HTMLOListElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}
