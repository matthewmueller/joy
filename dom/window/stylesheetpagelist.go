package window

import "github.com/matthewmueller/joy/js"

// StyleSheetPageList struct
// js:"StyleSheetPageList,omit"
type StyleSheetPageList struct {
}

// Item fn
// js:"item"
func (*StyleSheetPageList) Item(index uint) (c *CSSPageRule) {
	js.Rewrite("$_.item($1)", index)
	return c
}

// Length prop
// js:"length"
func (*StyleSheetPageList) Length() (length int) {
	js.Rewrite("$_.length")
	return length
}
