package htmlmodelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLModElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLModElement) GetCite() (cite string) {
	js.Rewrite("$<.cite")
	return cite
}

func (*HTMLModElement) SetCite(cite string) {
	js.Rewrite("$<.cite = $1", cite)
}

func (*HTMLModElement) GetDateTime() (dateTime string) {
	js.Rewrite("$<.dateTime")
	return dateTime
}

func (*HTMLModElement) SetDateTime(dateTime string) {
	js.Rewrite("$<.dateTime = $1", dateTime)
}
