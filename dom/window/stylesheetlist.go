package window

import "github.com/matthewmueller/golly/js"

// StyleSheetList struct
// js:"StyleSheetList,omit"
type StyleSheetList struct {
}

// Item fn
func (*StyleSheetList) Item(index *uint) (s StyleSheet) {
	js.Rewrite("$<.item($1)", index)
	return s
}

// Length prop
func (*StyleSheetList) Length() (length int) {
	js.Rewrite("$<.length")
	return length
}
