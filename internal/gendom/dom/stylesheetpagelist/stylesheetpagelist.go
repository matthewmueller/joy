package stylesheetpagelist

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/csspagerule"
	"github.com/matthewmueller/golly/js"
)

type StyleSheetPageList struct {
}

func (*StyleSheetPageList) Item(index uint) (c *csspagerule.CSSPageRule) {
	js.Rewrite("$<.item($1)", index)
	return c
}

func (*StyleSheetPageList) GetLength() (length int) {
	js.Rewrite("$<.length")
	return length
}
