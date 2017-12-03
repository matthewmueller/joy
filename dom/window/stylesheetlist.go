package window

import "github.com/matthewmueller/joy/macro"

// StyleSheetList struct
// js:"StyleSheetList,omit"
type StyleSheetList struct {
}

// Item fn
// js:"item"
func (*StyleSheetList) Item(index *uint) (s StyleSheet) {
	macro.Rewrite("$_.item($1)", index)
	return s
}

// Length prop
// js:"length"
func (*StyleSheetList) Length() (length int) {
	macro.Rewrite("$_.length")
	return length
}
