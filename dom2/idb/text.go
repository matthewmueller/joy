package idb

// js:"Text,omit"
type Text interface {
	CharacterData

	// SplitText
	SplitText(offset uint) (t Text)

	// WholeText
	WholeText() (wholeText string)
}
