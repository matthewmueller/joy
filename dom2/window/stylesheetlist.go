package window

import "github.com/matthewmueller/golly/js"

// js:"StyleSheetList,omit"
type StyleSheetList struct {
}

// Item
func (*StyleSheetList) Item(index *uint) (s StyleSheet) {
	js.Rewrite("$<.Item($1)", index)
	return s
}

// Length
func (*StyleSheetList) Length() (length int) {
	js.Rewrite("$<.Length")
	return length
}
