package window

// js:"Text,omit"
type Text interface {
	CharacterData

	// SplitText
	// js:"splitText"
	SplitText(offset uint) (t Text)

	// WholeText prop
	// js:"wholeText"
	WholeText() (wholeText string)
}
