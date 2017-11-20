package stylemedia

import "github.com/matthewmueller/golly/js"

type StyleMedia struct {
}

func (*StyleMedia) MatchMedium(mediaquery string) (b bool) {
	js.Rewrite("$<.matchMedium($1)", mediaquery)
	return b
}

func (*StyleMedia) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}
