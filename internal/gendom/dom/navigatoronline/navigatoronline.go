package navigatoronline

import "github.com/matthewmueller/golly/js"

type NavigatorOnLine struct {
}

func (*NavigatorOnLine) GetOnLine() (onLine bool) {
	js.Rewrite("$<.onLine")
	return onLine
}
