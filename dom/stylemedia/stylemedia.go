package stylemedia

import "github.com/matthewmueller/golly/js"

// StyleMedia struct
// js:"StyleMedia,omit"
type StyleMedia struct {
}

// MatchMedium fn
// js:"matchMedium"
func (*StyleMedia) MatchMedium(mediaquery string) (b bool) {
	js.Rewrite("$_.matchMedium($1)", mediaquery)
	return b
}

// Type prop
// js:"type"
func (*StyleMedia) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
