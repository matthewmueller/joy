package stylemedia

import "github.com/matthewmueller/joy/macro"

// StyleMedia struct
// js:"StyleMedia,omit"
type StyleMedia struct {
}

// MatchMedium fn
// js:"matchMedium"
func (*StyleMedia) MatchMedium(mediaquery string) (b bool) {
	macro.Rewrite("$_.matchMedium($1)", mediaquery)
	return b
}

// Type prop
// js:"type"
func (*StyleMedia) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
