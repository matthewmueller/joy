package linkstyle

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// LinkStyle struct
// js:"LinkStyle,omit"
type LinkStyle struct {
}

// Sheet prop
func (*LinkStyle) Sheet() (sheet window.StyleSheet) {
	js.Rewrite("$<.sheet")
	return sheet
}
