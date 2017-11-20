package text

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/characterdata"
	"github.com/matthewmueller/golly/js"
)

type Text struct {
	*characterdata.CharacterData
}

func (*Text) SplitText(offset uint) (t *Text) {
	js.Rewrite("$<.splitText($1)", offset)
	return t
}

func (*Text) GetWholeText() (wholeText string) {
	js.Rewrite("$<.wholeText")
	return wholeText
}
