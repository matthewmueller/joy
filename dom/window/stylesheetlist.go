package window

import "github.com/matthewmueller/golly/js"

// StyleSheetList struct
// js:"StyleSheetList,omit"
type StyleSheetList struct {
}

// Item fn
// js:"item"
func (*StyleSheetList) Item(index *uint) (s StyleSheet) {
	js.Rewrite("$_.item($1)", index)
	return s
}

// Length prop
// js:"length"
func (*StyleSheetList) Length() (length int) {
	js.Rewrite("$_.length")
	return length
}
