package stylemedia

import "github.com/matthewmueller/golly/js"

// StyleMedia struct
// js:"StyleMedia,omit"
type StyleMedia struct {
}

// MatchMedium fn
func (*StyleMedia) MatchMedium(mediaquery string) (b bool) {
	js.Rewrite("$<.matchMedium($1)", mediaquery)
	return b
}

// Type prop
func (*StyleMedia) Type() (kind string) {
	js.Rewrite("$<.type")
	return kind
}
