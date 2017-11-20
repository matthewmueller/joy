package comment

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/characterdata"
	"github.com/matthewmueller/golly/js"
)

type Comment struct {
	*characterdata.CharacterData
}

func (*Comment) GetText() (text string) {
	js.Rewrite("$<.text")
	return text
}

func (*Comment) SetText(text string) {
	js.Rewrite("$<.text = $1", text)
}
