package window

import "github.com/matthewmueller/golly/js"

// StyleSheetPageList struct
// js:"StyleSheetPageList,omit"
type StyleSheetPageList struct {
}

// Item fn
func (*StyleSheetPageList) Item(index uint) (c *CSSPageRule) {
	js.Rewrite("$<.item($1)", index)
	return c
}

// Length prop
func (*StyleSheetPageList) Length() (length int) {
	js.Rewrite("$<.length")
	return length
}
