package window

import "github.com/matthewmueller/golly/js"

// js:"StyleSheetPageList,omit"
type StyleSheetPageList struct {
}

// Item
func (*StyleSheetPageList) Item(index uint) (c *CSSPageRule) {
	js.Rewrite("$<.Item($1)", index)
	return c
}

// Length
func (*StyleSheetPageList) Length() (length int) {
	js.Rewrite("$<.Length")
	return length
}
