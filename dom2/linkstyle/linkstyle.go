package linkstyle

import "github.com/matthewmueller/golly/dom2/window"

// js:"LinkStyle,omit"
type LinkStyle interface {

	// Sheet
	Sheet() (sheet window.StyleSheet)
}
