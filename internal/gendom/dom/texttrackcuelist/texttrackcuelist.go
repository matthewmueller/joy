package texttrackcuelist

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/texttrackcue"
	"github.com/matthewmueller/golly/js"
)

type TextTrackCueList struct {
}

func (*TextTrackCueList) GetCueByID(id string) (t *texttrackcue.TextTrackCue) {
	js.Rewrite("$<.getCueById($1)", id)
	return t
}

func (*TextTrackCueList) Item(index uint) (t *texttrackcue.TextTrackCue) {
	js.Rewrite("$<.item($1)", index)
	return t
}

func (*TextTrackCueList) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}
