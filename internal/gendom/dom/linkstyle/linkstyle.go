package linkstyle

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/stylesheet"
	"github.com/matthewmueller/golly/js"
)

type LinkStyle struct {
}

func (*LinkStyle) GetSheet() (sheet *stylesheet.StyleSheet) {
	js.Rewrite("$<.sheet")
	return sheet
}
