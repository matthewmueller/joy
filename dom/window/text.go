package window

// Text interface
// js:"Text"
type Text interface {
	CharacterData

	// SplitText
	// js:"splitText"
	// jsrewrite:"$_.splitText($1)"
	SplitText(offset uint) (t Text)

	// WholeText prop
	// js:"wholeText"
	// jsrewrite:"$_.wholeText"
	WholeText() (wholeText string)
}
