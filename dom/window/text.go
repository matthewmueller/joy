package window

import "github.com/matthewmueller/golly/js"

// js:"Text,omit"
type Text interface {
	CharacterData

	// SplitText
	// js:"splitText,rewrite=splittext"
	SplitText(offset uint) (t Text)

	// WholeText prop
	// js:"wholeText,rewrite=wholetext"
	WholeText() (wholeText string)
}

// splittext fn
func splittext(offset uint) (t Text) {
	js.Rewrite("$<.splitText($1)", offset)
	return t
}

// wholetext prop
func wholetext() (wholeText string) {
	js.Rewrite("$<.wholeText")
	return wholeText
}
