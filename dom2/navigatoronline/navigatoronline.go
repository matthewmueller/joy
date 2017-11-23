package navigatoronline

import "github.com/matthewmueller/golly/js"

// NavigatorOnLine struct
// js:"NavigatorOnLine,omit"
type NavigatorOnLine struct {
}

// OnLine
func (*NavigatorOnLine) OnLine() (onLine bool) {
	js.Rewrite("$<.OnLine")
	return onLine
}
