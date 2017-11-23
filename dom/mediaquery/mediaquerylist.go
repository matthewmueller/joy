package mediaquery

import "github.com/matthewmueller/golly/js"

// MediaQueryList struct
// js:"MediaQueryList,omit"
type MediaQueryList struct {
}

// AddListener fn
func (*MediaQueryList) AddListener(listener func(mql *MediaQueryList)) {
	js.Rewrite("$<.addListener($1)", listener)
}

// RemoveListener fn
func (*MediaQueryList) RemoveListener(listener func(mql *MediaQueryList)) {
	js.Rewrite("$<.removeListener($1)", listener)
}

// Matches prop
func (*MediaQueryList) Matches() (matches bool) {
	js.Rewrite("$<.matches")
	return matches
}

// Media prop
func (*MediaQueryList) Media() (media string) {
	js.Rewrite("$<.media")
	return media
}
