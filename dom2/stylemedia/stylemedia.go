package stylemedia

import "github.com/matthewmueller/golly/js"

// js:"StyleMedia,omit"
type StyleMedia struct {
}

// MatchMedium
func (*StyleMedia) MatchMedium(mediaquery string) (b bool) {
	js.Rewrite("$<.MatchMedium($1)", mediaquery)
	return b
}

// Type
func (*StyleMedia) Type() (kind string) {
	js.Rewrite("$<.Type")
	return kind
}
