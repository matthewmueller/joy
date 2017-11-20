package htmltimeelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLTimeElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLTimeElement) GetDateTime() (dateTime string) {
	js.Rewrite("$<.dateTime")
	return dateTime
}

func (*HTMLTimeElement) SetDateTime(dateTime string) {
	js.Rewrite("$<.dateTime = $1", dateTime)
}
