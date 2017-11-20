package stylesheetlist

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/stylesheet"
	"github.com/matthewmueller/golly/js"
)

type StyleSheetList struct {
}

func (*StyleSheetList) Item(index uint) (s *stylesheet.StyleSheet) {
	js.Rewrite("$<.item($1)", index)
	return s
}

func (*StyleSheetList) GetLength() (length int) {
	js.Rewrite("$<.length")
	return length
}
