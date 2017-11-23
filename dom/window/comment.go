package window

import "github.com/matthewmueller/golly/js"

// Comment struct
// js:"Comment,omit"
type Comment struct {
	CharacterData
}

// Text prop
func (*Comment) Text() (text string) {
	js.Rewrite("$<.text")
	return text
}

// Text prop
func (*Comment) SetText(text string) {
	js.Rewrite("$<.text = $1", text)
}
