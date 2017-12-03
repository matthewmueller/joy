package linkstyle

import "github.com/matthewmueller/joy/dom/window"

// LinkStyle interface
// js:"LinkStyle"
type LinkStyle interface {

	// Sheet prop
	// js:"sheet"
	// jsrewrite:"$_.sheet"
	Sheet() (sheet window.StyleSheet)
}
