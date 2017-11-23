package mediaquery

import "github.com/matthewmueller/golly/js"

// js:"MediaQueryList,omit"
type MediaQueryList struct {
}

// AddListener
func (*MediaQueryList) AddListener(listener func(mql *MediaQueryList)) {
	js.Rewrite("$<.AddListener($1)", listener)
}

// RemoveListener
func (*MediaQueryList) RemoveListener(listener func(mql *MediaQueryList)) {
	js.Rewrite("$<.RemoveListener($1)", listener)
}

// Matches
func (*MediaQueryList) Matches() (matches bool) {
	js.Rewrite("$<.Matches")
	return matches
}

// Media
func (*MediaQueryList) Media() (media string) {
	js.Rewrite("$<.Media")
	return media
}
