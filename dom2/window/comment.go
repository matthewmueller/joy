package window

import "github.com/matthewmueller/golly/js"

// Comment struct
// js:"Comment,omit"
type Comment struct {
	CharacterData
}

// Text
func (*Comment) Text() (text string) {
	js.Rewrite("$<.Text")
	return text
}

// Text
func (*Comment) SetText(text string) {
	js.Rewrite("$<.Text = $1", text)
}
