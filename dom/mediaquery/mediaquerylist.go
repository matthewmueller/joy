package mediaquery

import "github.com/matthewmueller/joy/macro"

// MediaQueryList struct
// js:"MediaQueryList,omit"
type MediaQueryList struct {
}

// AddListener fn
// js:"addListener"
func (*MediaQueryList) AddListener(listener func(mql *MediaQueryList)) {
	macro.Rewrite("$_.addListener($1)", listener)
}

// RemoveListener fn
// js:"removeListener"
func (*MediaQueryList) RemoveListener(listener func(mql *MediaQueryList)) {
	macro.Rewrite("$_.removeListener($1)", listener)
}

// Matches prop
// js:"matches"
func (*MediaQueryList) Matches() (matches bool) {
	macro.Rewrite("$_.matches")
	return matches
}

// Media prop
// js:"media"
func (*MediaQueryList) Media() (media string) {
	macro.Rewrite("$_.media")
	return media
}
