package mediaquerylist

import "github.com/matthewmueller/golly/js"

type MediaQueryList struct {
}

func (*MediaQueryList) AddListener(listener func(mql *MediaQueryList)) {
	js.Rewrite("$<.addListener($1)", listener)
}

func (*MediaQueryList) RemoveListener(listener func(mql *MediaQueryList)) {
	js.Rewrite("$<.removeListener($1)", listener)
}

func (*MediaQueryList) GetMatches() (matches bool) {
	js.Rewrite("$<.matches")
	return matches
}

func (*MediaQueryList) GetMedia() (media string) {
	js.Rewrite("$<.media")
	return media
}
