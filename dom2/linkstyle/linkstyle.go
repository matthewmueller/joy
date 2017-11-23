package linkstyle

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// LinkStyle struct
// js:"LinkStyle,omit"
type LinkStyle struct {
}

// Sheet
func (*LinkStyle) Sheet() (sheet window.StyleSheet) {
	js.Rewrite("$<.Sheet")
	return sheet
}
