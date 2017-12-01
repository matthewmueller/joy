package mediaquery

import "github.com/matthewmueller/golly/js"

// MediaQueryList struct
// js:"MediaQueryList,omit"
type MediaQueryList struct {
}

// AddListener fn
// js:"addListener"
func (*MediaQueryList) AddListener(listener func(mql *MediaQueryList)) {
	js.Rewrite("$_.addListener($1)", listener)
}

// RemoveListener fn
// js:"removeListener"
func (*MediaQueryList) RemoveListener(listener func(mql *MediaQueryList)) {
	js.Rewrite("$_.removeListener($1)", listener)
}

// Matches prop
// js:"matches"
func (*MediaQueryList) Matches() (matches bool) {
	js.Rewrite("$_.matches")
	return matches
}

// Media prop
// js:"media"
func (*MediaQueryList) Media() (media string) {
	js.Rewrite("$_.media")
	return media
}
