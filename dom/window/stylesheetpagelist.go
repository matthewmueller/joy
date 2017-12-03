package window

import "github.com/matthewmueller/joy/macro"

// StyleSheetPageList struct
// js:"StyleSheetPageList,omit"
type StyleSheetPageList struct {
}

// Item fn
// js:"item"
func (*StyleSheetPageList) Item(index uint) (c *CSSPageRule) {
	macro.Rewrite("$_.item($1)", index)
	return c
}

// Length prop
// js:"length"
func (*StyleSheetPageList) Length() (length int) {
	macro.Rewrite("$_.length")
	return length
}
